package main

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	a2 "google.golang.org/adk/agent"
	"google.golang.org/adk/artifact/gcsartifact"
	"google.golang.org/adk/cmd/launcher"
	"google.golang.org/adk/cmd/launcher/full"
	"ril.api-ia/internal/agent"
)

func main() {
	ctx := context.Background()
	_ = godotenv.Overload()
	coordinatorAgent, err := agent.NewRilAgent(ctx)
	if err != nil {
		panic(err)
	}
	artifactService, _ := gcsartifact.NewService(ctx, os.Getenv("ARTIFACT_BUCKET_NAME"))
	config := &launcher.Config{
		ArtifactService: artifactService,
		AgentLoader:     a2.NewSingleLoader(coordinatorAgent),
	}
	l := full.NewLauncher()
	if err = l.Execute(ctx, config, os.Args[1:]); err != nil {
		log.Fatalf("Run failed: %v\n\n%s", err, l.CommandLineSyntax())
	}
}
