package main

import (
	"context"
	cryptoRand "crypto/rand"
	"log"
	"math/rand"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"google.golang.org/adk/artifact"
	"google.golang.org/adk/runner"
	"google.golang.org/adk/session"
	"ril.api-ia/internal/agent"
	"ril.api-ia/internal/application/usecase"
	"ril.api-ia/internal/domain/entity"
	"ril.api-ia/internal/infrastructure/http/handler"
	"ril.api-ia/internal/infrastructure/http/middleware"
	m "ril.api-ia/internal/infrastructure/repository/memory"
)

func main() {
	ctx := context.Background()
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}
	//AGENT
	sessionService := session.InMemoryService()
	artifactService := artifact.InMemoryService()
	runn, err := runner.New(runner.Config{
		AppName:         os.Getenv("APP_NAME"),
		Agent:           agent.GetRilAgent(ctx),
		SessionService:  sessionService,
		ArtifactService: artifactService,
	})
	if err != nil {
		log.Fatal(err)
	}
	//REPOSITORIES
	userRepository := m.NewUserRepository()

	//Fixers and seeds (ONLY FOR LOCAL Development)
	if os.Getenv("APP_ENV") == "local" {
		FillMockUsers(userRepository)
	}
	//USE-CASES
	sessionUseCase := usecase.NewSessionUseCase(ctx, sessionService, userRepository)

	//HTTP
	r := gin.Default()

	//MIDDLEWARES
	r.Use(middleware.AuthMiddleware(userRepository))

	//HANDLERS
	sessionHandler := handler.NewSessionHandler(sessionUseCase)
	runHandler := handler.NewRunHandler(ctx, *runn, *sessionUseCase)

	//ROUTES
	sessions := r.Group("/sessions")
	{
		sessions.POST("", sessionHandler.CreateSession)
		sessions.GET("", sessionHandler.ListSessions)
		sessions.GET("/:session_id", sessionHandler.GetSession)
		sessions.DELETE("/:session_id", sessionHandler.DeleteSession)
	}
	r.POST("/run-sse", runHandler.RunSSE)

	//RUNNER
	r.Run(":8080")
}

func FillMockUsers(userRepository *m.UserRepository) {
	apiKey := cryptoRand.Text()
	apiKey2 := "HRBRN2GP7BS65B7JQAPON4L4UJ"
	apiKey3 := cryptoRand.Text()
	users := []*entity.User{
		{
			Id:         rand.Int63(),
			FirstName:  "John",
			LastName:   "Doe",
			IdTeam:     rand.Int63(),
			ApiAiToken: &apiKey,
		},
		{
			Id:         rand.Int63(),
			FirstName:  "Jane",
			LastName:   "Doe",
			IdTeam:     rand.Int63(),
			ApiAiToken: &apiKey2,
		},
		{
			Id:         rand.Int63(),
			FirstName:  "Janet",
			LastName:   "Doe",
			IdTeam:     rand.Int63(),
			ApiAiToken: &apiKey3,
		},
	}

	for _, user := range users {
		log.Printf("token: %s\n", *user.ApiAiToken)
		err := userRepository.Save(user)
		if err != nil {
			log.Fatal(err)
		}
	}
}
