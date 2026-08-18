package agent

import (
	"log"

	"google.golang.org/adk/agent"
	"google.golang.org/adk/model"
	"google.golang.org/adk/tool"
	"google.golang.org/adk/tool/agenttool"

	"ril.api-ia/internal/agent/subagents/educationagent"
	"ril.api-ia/internal/agent/subagents/girsuagent"
	"ril.api-ia/internal/agent/subagents/professionalizationagent"
	"ril.api-ia/internal/agent/subagents/securityagent"
	"ril.api-ia/internal/infrastructure/repository/tree_agent"
)

type domainAgentConstructor func(m model.LLM, treeManager *tree_agent.TreeCacheManager) (agent.Agent, error)

type domainAgentSpec struct {
	name        string
	constructor domainAgentConstructor
}

var domainAgentSpecs = []domainAgentSpec{
	{"girsu", girsuagent.NewGirsuAgent},
	{"security", securityagent.NewSecurityAgent},
	{"profesionalizacion", professionalizationagent.NewProfessionalizationAgent},
	{"education", educationagent.NewEducationAgent},
}

func buildAgentTools(m model.LLM, treeManager *tree_agent.TreeCacheManager) []tool.Tool {
	agentTools := make([]tool.Tool, 0, len(domainAgentSpecs))
	for _, spec := range domainAgentSpecs {
		a, err := spec.constructor(m, treeManager)
		if err != nil {
			log.Fatalf("Failed to create %s agent: %v", spec.name, err)
		}
		agentTools = append(agentTools, agenttool.New(a, &agenttool.Config{}))
	}
	return agentTools
}
