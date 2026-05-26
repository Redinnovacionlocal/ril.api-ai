package main

import (
	"context"
	"log"
	"os"
	"time"

	"cloud.google.com/go/storage"
	"github.com/googleapis/mcp-toolbox-sdk-go/tbadk"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
	a2 "google.golang.org/adk/agent"
	"google.golang.org/adk/artifact/gcsartifact"
	"google.golang.org/adk/cmd/launcher"
	"google.golang.org/adk/cmd/launcher/console"
	"google.golang.org/adk/cmd/launcher/universal"
	"google.golang.org/adk/cmd/launcher/web"
	"google.golang.org/adk/cmd/launcher/web/api"
	"google.golang.org/adk/cmd/launcher/web/webui"
	"google.golang.org/adk/model/gemini"
	"google.golang.org/adk/plugin"
	"google.golang.org/adk/runner"
	"ril.api-ia/internal/agent"
	"ril.api-ia/internal/agent/plugin/agent_active_plugin"
	"ril.api-ia/internal/agent/plugin/title_plugin"
	"ril.api-ia/internal/agent/subagents"
	"ril.api-ia/internal/agent/subagents/securityagent"
	"ril.api-ia/internal/infrastructure/repository/tree_agent"
)

func main() {
	ctx := context.Background()
	_ = godotenv.Overload()
	toolboxClient, err := tbadk.NewToolboxClient(os.Getenv("TOOLBOX_CLIENT_URL"))
	if err != nil {
		log.Fatal("Error initializing Toolbox client:", err)
	}
	coordinatorAgent, err := agent.NewRilAgent(ctx, toolboxClient)
	model2_5, err := gemini.NewModel(ctx, os.Getenv("AGENT_RAG_MODEL"), nil)
	model3, err := gemini.NewModel(ctx, os.Getenv("AGENT_MODEL"), nil)
	agentKnowledge, err := subagents.NewRagAgent(model2_5)
	gcsClient, _ := storage.NewClient(ctx)
	dbAgent, err := sqlx.Open("pgx", os.Getenv("DATABASE_AGENT_DSN"))
	if err != nil {
		log.Fatal("Error connecting to agent DB:", err)
	}
	defer dbAgent.Close()

	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDR"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	})
	treeRepo := tree_agent.NewQuestionTreeRepository(dbAgent)

	treeManager := tree_agent.NewTreeCacheManager(
		gcsClient,
		os.Getenv("AGENT_TREE_BUCKET"),
		treeRepo,
		rdb,
		1*time.Hour,
	)
	securityAgent, err := securityagent.NewSecurityAgent(model3, treeManager)
	if err != nil {
		panic(err)
	}
	loader, _ := a2.NewMultiLoader(coordinatorAgent, agentKnowledge, securityAgent)
	artifactService, _ := gcsartifact.NewService(ctx, os.Getenv("ARTIFACT_BUCKET_NAME"))
	titlePlugin, _ := title_plugin.New(ctx, "title_plugin")
	agentActivePlugin, _ := agent_active_plugin.New(ctx, "agent_active_plugin")
	config := &launcher.Config{
		ArtifactService: artifactService,
		AgentLoader:     loader,
		PluginConfig: runner.PluginConfig{
			Plugins: []*plugin.Plugin{
				titlePlugin,
				agentActivePlugin,
			},
		},
	}
	l := universal.NewLauncher(
		console.NewLauncher(),
		web.NewLauncher(api.NewLauncher(), webui.NewLauncher()),
	)
	if err = l.Execute(ctx, config, os.Args[1:]); err != nil {
		log.Fatalf("Run failed: %v\n\n%s", err, l.CommandLineSyntax())
	}
}
