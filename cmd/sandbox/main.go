package main

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	a2 "google.golang.org/adk/agent"
	"google.golang.org/adk/artifact/gcsartifact"
	"google.golang.org/adk/cmd/launcher"
	"google.golang.org/adk/cmd/launcher/console"
	"google.golang.org/adk/cmd/launcher/universal"
	"google.golang.org/adk/cmd/launcher/web"
	"google.golang.org/adk/cmd/launcher/web/api"
	"google.golang.org/adk/cmd/launcher/web/webui"
	"google.golang.org/adk/model/gemini"
	"ril.api-ia/internal/agent"
	"ril.api-ia/internal/agent/subagents"
	"ril.api-ia/internal/agent/subagents/securityagent"
)

func main() {
	ctx := context.Background()
	_ = godotenv.Overload()
	coordinatorAgent, err := agent.NewRilAgent(ctx)
	model2_5, err := gemini.NewModel(ctx, os.Getenv("AGENT_RAG_MODEL"), nil)
	model3, err := gemini.NewModel(ctx, os.Getenv("AGENT_MODEL"), nil)
	agentKnowledge, err := subagents.NewRagAgent(model2_5)
	securityAgent, err := securityagent.NewSecurityAgent(model3)
	if err != nil {
		panic(err)
	}
	loader, _ := a2.NewMultiLoader(coordinatorAgent, agentKnowledge, securityAgent)
	artifactService, _ := gcsartifact.NewService(ctx, os.Getenv("ARTIFACT_BUCKET_NAME"))
	config := &launcher.Config{
		ArtifactService: artifactService,
		AgentLoader:     loader,
	}
	l := universal.NewLauncher(
		console.NewLauncher(),
		web.NewLauncher(api.NewLauncher(), webui.NewLauncher()),
	)
	if err = l.Execute(ctx, config, os.Args[1:]); err != nil {
		log.Fatalf("Run failed: %v\n\n%s", err, l.CommandLineSyntax())
	}
}
