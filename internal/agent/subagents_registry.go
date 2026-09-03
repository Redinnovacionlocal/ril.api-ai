package agent

import (
	"log"
	"os"
	"strings"

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

type DomainAgentConstructor func(m model.LLM, treeManager *tree_agent.TreeCacheManager) (agent.Agent, error)

type DomainAgentSpec struct {
	Name        string
	DomainLabel string
	UseCase     string
	Constructor DomainAgentConstructor
}

var domainAgentSpecs = []DomainAgentSpec{
	{"girsu_agent", "Residuos", "gestión de residuos", girsuagent.NewGirsuAgent},
	{"security_agent", "Seguridad", "seguridad pública", securityagent.NewSecurityAgent},
	{"professionalization_agent", "Profesionalización", "profesionalización del municipio", professionalizationagent.NewProfessionalizationAgent},
	{"education_agent", "Educación", "educación", educationagent.NewEducationAgent},
}

func EnabledDomainAgentSpecs() []DomainAgentSpec {
	raw := os.Getenv("ENABLED_DOMAIN_AGENTS")
	if raw == "" {
		return domainAgentSpecs
	}

	enabled := make(map[string]bool, len(domainAgentSpecs))
	for _, name := range strings.Split(raw, ";") {
		enabled[strings.TrimSpace(name)] = true
	}

	out := make([]DomainAgentSpec, 0, len(domainAgentSpecs))
	for _, spec := range domainAgentSpecs {
		if enabled[spec.Name] {
			out = append(out, spec)
		}
	}
	log.Printf("Enabled domain agents: %v", out)
	return out
}

func buildAgentTools(m model.LLM, treeManager *tree_agent.TreeCacheManager) []tool.Tool {
	specs := EnabledDomainAgentSpecs()
	agentTools := make([]tool.Tool, 0, len(specs))
	for _, spec := range specs {
		a, err := spec.Constructor(m, treeManager)
		if err != nil {
			log.Fatalf("Failed to create %s agent: %v", spec.Name, err)
		}
		agentTools = append(agentTools, agenttool.New(a, &agenttool.Config{}))
	}
	return agentTools
}
