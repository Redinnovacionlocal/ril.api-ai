package agent_active_plugin

import (
	"context"
	"log"

	"google.golang.org/adk/agent"
	"google.golang.org/adk/plugin"
	"google.golang.org/genai"
)

type AgentActivePlugin struct {
	Name string
}

func New(ctx context.Context, name string) (*plugin.Plugin, error) {
	if name == "" {
		name = "agent_active_plugin"
	}

	a := &AgentActivePlugin{
		Name: name,
	}

	return plugin.New(plugin.Config{
		Name:                a.Name,
		BeforeAgentCallback: a.BeforeAgentCallback,
	})

}

func (a *AgentActivePlugin) BeforeAgentCallback(ctx agent.CallbackContext) (*genai.Content, error) {
	// Add last agent active
	agentAuthor := ctx.AgentName()
	err := ctx.State().Set("last_agent_active", agentAuthor)
	if err != nil {
		// Log the error but don't fail the callback
		log.Printf("failed to set last_agent_active: %v", err)
	}
	log.Printf("Last active agent set to: %s", agentAuthor)
	return nil, nil
}
