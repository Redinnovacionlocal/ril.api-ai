package main

import (
	"context"
	cryptoRand "crypto/rand"
	"log"
	"math/rand"
	"os"
	"time"

	"cloud.google.com/go/storage"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/googleapis/mcp-toolbox-sdk-go/tbadk"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
	internalagent "google.golang.org/adk/agent"
	"google.golang.org/adk/artifact"
	"google.golang.org/adk/artifact/gcsartifact"
	"google.golang.org/adk/memory"
	"google.golang.org/adk/model/gemini"
	"google.golang.org/adk/plugin"
	"google.golang.org/adk/runner"
	"google.golang.org/adk/session"
	"google.golang.org/adk/session/database"
	"gorm.io/driver/postgres"
	"ril.api-ia/internal/agent"
	"ril.api-ia/internal/agent/plugin/agent_active_plugin"
	"ril.api-ia/internal/agent/plugin/title_plugin"
	"ril.api-ia/internal/agent/subagents/securityagent"
	session2 "ril.api-ia/internal/application/service/session"
	"ril.api-ia/internal/application/usecase"
	"ril.api-ia/internal/domain/entity"
	"ril.api-ia/internal/domain/repository"
	"ril.api-ia/internal/infrastructure/http/handler"
	"ril.api-ia/internal/infrastructure/http/middleware"
	m "ril.api-ia/internal/infrastructure/repository/memory"
	"ril.api-ia/internal/infrastructure/repository/sql"
	"ril.api-ia/internal/infrastructure/repository/tree_agent"
)

func main() {
	ctx := context.Background()
	_ = godotenv.Overload()

	runMigrations(os.Getenv("DATABASE_AGENT_DSN"))

	//Init redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDR"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	})

	dbAgent, err := sqlx.Open("pgx", os.Getenv("DATABASE_AGENT_DSN"))
	if err != nil {
		log.Fatal("Error connecting to agent DB:", err)
	}
	defer dbAgent.Close()

	var dbCore *sqlx.DB
	if os.Getenv("APP_ENV") != "local" {
		dbCore, err = sqlx.ConnectContext(ctx, "mysql", os.Getenv("DATABASE_CORE_DSN"))
		if err != nil {
			log.Fatal("Error connecting to core DB:", err)
		}
		defer dbCore.Close()
	}

	// Agents service and runners
	sessionService := initializeSessionService()
	artifactService, _ := gcsartifact.NewService(ctx, os.Getenv("ARTIFACT_BUCKET_NAME"))
	userRepository, eventFeedbackRepository := InitializeRepositories(ctx, dbCore, dbAgent)

	runners := initializeRunner(ctx, sessionService, artifactService, dbAgent, rdb)

	// Use cases
	sessionUseCase := usecase.NewSessionUseCase(ctx, sessionService, userRepository)
	userUseCase := usecase.NewUserUseCase(ctx, userRepository, rdb)
	eventFeedbackUseCase := usecase.NewEventFeedbackUseCase(ctx, eventFeedbackRepository)
	transcribeUseCase := usecase.NewTranscribeUseCase(ctx)

	// HTTP Server and routes
	router := setupRouter(ctx, sessionUseCase, userUseCase, eventFeedbackUseCase, transcribeUseCase, runners)
	startServer(router)
}

func initializeSessionService() session2.Service {
	sessionService, err := database.NewSessionService(postgres.Open(os.Getenv("DATABASE_AGENT_DSN")))
	mySessionService := session2.NewMyDatabaseService(sessionService, postgres.Open(os.Getenv("DATABASE_AGENT_DSN")))
	if err != nil {
		log.Fatal("Error initializing session service:", err)
	}
	return mySessionService
}

func InitializeRepositories(ctx context.Context, dbCore *sqlx.DB, dbAgent *sqlx.DB) (repository.UserRepository, repository.EventFeedbackRepository) {
	if os.Getenv("APP_ENV") != "local" {
		log.Println("Running id dis production modes with SQL user repository")

		eventFeedbackRepository := sql.NewEventFeedbackRepository(dbAgent)
		userRepository := sql.NewUserRepository(dbCore)
		return userRepository, eventFeedbackRepository
	}
	userRepository := m.NewUserRepository()
	eventFeedbackRepository := m.NewEventFeedbackRepository()
	seedMockUsers(userRepository)
	return userRepository, eventFeedbackRepository
}

func initializeRunner(ctx context.Context, sessionService session.Service, artifactService artifact.Service, dbAgent *sqlx.DB, rdb *redis.Client) map[string]*runner.Runner {
	securityAgentName := os.Getenv("AGENT_SECURITY_NAME")
	toolboxClient, err := tbadk.NewToolboxClient(os.Getenv("TOOLBOX_CLIENT_URL"))
	if err != nil {
		log.Fatal("Error initializing Toolbox client:", err)
	}

	rilAgent, err := agent.NewRilAgent(ctx, toolboxClient)
	if err != nil {
		log.Fatal("Error initializing RilAgent:", err)
	}

	model, err := gemini.NewModel(ctx, os.Getenv("AGENT_MODEL"), nil)
	if err != nil {
		log.Fatal("Error initializing Gemini model:", err)
	}

	gcsClient, _ := storage.NewClient(ctx)
	if err != nil {
		log.Fatal("Error initializing GCS client:", err)
	}

	treeRepo := tree_agent.NewQuestionTreeRepository(dbAgent)

	treeManager := tree_agent.NewTreeCacheManager(
		gcsClient,
		os.Getenv("AGENT_TREE_BUCKET"),
		treeRepo,
		rdb,
		1*time.Hour,
	)

	securityAgent, err := securityagent.NewSecurityAgent(model, treeManager)
	if err != nil {
		log.Fatal("Error initializing SecurityAgent:", err)
	}

	return map[string]*runner.Runner{
		"orchestrator":    buildRunner(ctx, rilAgent, sessionService, artifactService),
		securityAgentName: buildRunner(ctx, securityAgent, sessionService, artifactService),
	}
}

func buildRunner(ctx context.Context, ag internalagent.Agent, sessionService session.Service, artifactService artifact.Service) *runner.Runner {
	memoryService := memory.InMemoryService()
	titlePlugin, _ := title_plugin.New(ctx, "title_plugin")
	agentActivePlugin, _ := agent_active_plugin.New(ctx, "agent_active_plugin")

	r, err := runner.New(runner.Config{
		AppName:         os.Getenv("APP_NAME"),
		Agent:           ag,
		SessionService:  sessionService,
		ArtifactService: artifactService,
		MemoryService:   memoryService,
		PluginConfig: runner.PluginConfig{
			Plugins: []*plugin.Plugin{
				titlePlugin,
				agentActivePlugin,
			},
		},
	})
	if err != nil {
		log.Fatal("Error initializing runner:", err)
	}
	return r
}

func setupRouter(ctx context.Context, sessionUseCase *usecase.SessionUseCase, userUseCase *usecase.UserUseCase, feedbackUseCase *usecase.EventFeedbackUseCase, transcribeUseCase *usecase.TranscribeUseCase, runners map[string]*runner.Runner) *gin.Engine {
	r := gin.Default()
	configCors := cors.DefaultConfig()
	configCors.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Accept", "Authorization", "x-agent"}
	configCors.AllowAllOrigins = true
	r.Use(cors.New(configCors))
	r.Use(middleware.AuthMiddleware(*userUseCase))

	sessionHandler := handler.NewSessionHandler(sessionUseCase)
	feedbackHandler := handler.NewFeedbackHandler(ctx, *feedbackUseCase)
	runHandler := handler.NewRunHandler(ctx, runners, *sessionUseCase)
	speechToTextHandler := handler.NewSpeechToTextHandler(ctx, transcribeUseCase)

	registerRoutes(r, sessionHandler, runHandler, feedbackHandler, speechToTextHandler)

	return r
}

func registerRoutes(r *gin.Engine, sessionHandler *handler.SessionHandler, runHandler *handler.RunHandler, feedbackHandler *handler.FeedbackHandler, speechToTextHandler *handler.SpeechToTextHandler) {
	sessions := r.Group("/sessions")
	{
		sessions.POST("", sessionHandler.CreateSession)
		sessions.GET("", sessionHandler.ListSessions)
		sessions.GET("/:session_id", sessionHandler.GetSession)
		sessions.DELETE("/:session_id", sessionHandler.DeleteSession)
		sessions.PUT("/:session_id/title", sessionHandler.UpdateSessionTitle)
	}
	r.POST("/speech-to-text", speechToTextHandler.GenerateTranscription)
	r.POST("/events/:invocation_id/feedback", feedbackHandler.SaveFeedback)
	r.POST("/run-sse", runHandler.RunSSE)
}

func startServer(router *gin.Engine) {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Starting server on port %s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatal("Error starting server:", err)
	}
}

func seedMockUsers(userRepository *m.UserRepository) {
	log.Println("Seeding mock users for local development")

	users := []*entity.User{
		createMockUser("John", "Doe"),
		createMockUser("Jane", "Doe"),
		createMockUser("Janet", "Doe"),
	}

	for _, user := range users {
		if err := userRepository.Save(user); err != nil {
			log.Fatal("Error saving mock user:", err)
		}
		log.Printf("Created mock user: %s %s (token: %s)", user.FirstName, user.LastName, *user.ApiAiToken)
	}
}

func createMockUser(firstName, lastName string) *entity.User {
	apiKey := generateAPIKey()
	return &entity.User{
		Id:         generateRandomID(),
		FirstName:  firstName,
		LastName:   lastName,
		IdTeam:     generateRandomID(),
		ApiAiToken: &apiKey,
	}
}

func generateAPIKey() string {
	return cryptoRand.Text()
}

func generateRandomID() int64 {
	return rand.Int63()
}
