package main

import (
	"context"
	cryptoRand "crypto/rand"
	"log"
	"math/rand"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
	"google.golang.org/adk/artifact"
	"google.golang.org/adk/runner"
	"google.golang.org/adk/session"
	"google.golang.org/adk/session/database"
	"gorm.io/driver/postgres"
	"ril.api-ia/internal/agent"
	"ril.api-ia/internal/application/usecase"
	"ril.api-ia/internal/domain/entity"
	"ril.api-ia/internal/domain/repository"
	"ril.api-ia/internal/infrastructure/http/handler"
	"ril.api-ia/internal/infrastructure/http/middleware"
	m "ril.api-ia/internal/infrastructure/repository/memory"
	"ril.api-ia/internal/infrastructure/repository/sql"
)

func main() {
	ctx := context.Background()

	_ = godotenv.Overload()
	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDR"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	})

	// Agents service and runners
	sessionService := initializeSessionService()
	artifactService := artifact.InMemoryService()
	userRepository, eventFeedbackRepository := initializeRepositories(ctx)
	runn := initializeRunner(ctx, sessionService, artifactService)

	// Use cases
	sessionUseCase := usecase.NewSessionUseCase(ctx, sessionService, userRepository)
	userUseCase := usecase.NewUserUseCase(ctx, userRepository, rdb)
	eventFeedbackUseCase := usecase.NewEventFeedbackUseCase(ctx, eventFeedbackRepository)

	// HTTP Server and routes
	router := setupRouter(ctx, sessionUseCase, userUseCase, eventFeedbackUseCase, runn)
	startServer(router)
}

func initializeSessionService() session.Service {
	if os.Getenv("APP_ENV") == "local" {
		return session.InMemoryService()
	}

	sessionService, err := database.NewSessionService(postgres.Open(os.Getenv("DATABASE_AGENT_DSN")))
	if err != nil {
		log.Fatal("Error initializing session service:", err)
	}
	return sessionService
}

func initializeRepositories(ctx context.Context) (repository.UserRepository, repository.EventFeedbackRepository) {
	if os.Getenv("APP_ENV") != "local" {
		log.Println("Running id dis production modes with SQL user repository")
		db, err := sqlx.ConnectContext(ctx, "mysql", os.Getenv("DATABASE_CORE_DSN"))
		if err != nil {
			log.Fatal("Error connecting to the database:", err)
		}
		dbAgen, err := sqlx.Open("pgx", os.Getenv("DATABASE_AGENT_DSN"))
		if err != nil {
			log.Fatal("Error connecting to the agent database:", err)
		}
		eventFeedbackRepository := sql.NewEventFeedbackRepository(dbAgen)
		userRepository := sql.NewUserRepository(db)
		return userRepository, eventFeedbackRepository
	}
	userRepository := m.NewUserRepository()
	eventFeedbackRepository := m.NewEventFeedbackRepository()
	seedMockUsers(userRepository)
	return userRepository, eventFeedbackRepository
}

func initializeRunner(ctx context.Context, sessionService session.Service, artifactService artifact.Service) *runner.Runner {
	runn, err := runner.New(runner.Config{
		AppName:         os.Getenv("APP_NAME"),
		Agent:           agent.GetRilAgent(ctx),
		SessionService:  sessionService,
		ArtifactService: artifactService,
	})
	if err != nil {
		log.Fatal("Error initializing runner:", err)
	}
	return runn
}

func setupRouter(ctx context.Context, sessionUseCase *usecase.SessionUseCase, userUseCase *usecase.UserUseCase, feedbackUseCase *usecase.EventFeedbackUseCase, runn *runner.Runner) *gin.Engine {
	r := gin.Default()
	r.Use(cors.Default())
	r.Use(middleware.AuthMiddleware(*userUseCase))

	sessionHandler := handler.NewSessionHandler(sessionUseCase)
	feedbackHandler := handler.NewFeedbackHandler(ctx, *feedbackUseCase)
	runHandler := handler.NewRunHandler(ctx, *runn, *sessionUseCase)

	registerRoutes(r, sessionHandler, runHandler, feedbackHandler)

	return r
}

func registerRoutes(r *gin.Engine, sessionHandler *handler.SessionHandler, runHandler *handler.RunHandler, feedbackHandler *handler.FeedbackHandler) {
	sessions := r.Group("/sessions")
	{
		sessions.POST("", sessionHandler.CreateSession)
		sessions.GET("", sessionHandler.ListSessions)
		sessions.GET("/:session_id", sessionHandler.GetSession)
		sessions.DELETE("/:session_id", sessionHandler.DeleteSession)
	}
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
