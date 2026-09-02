package professionalizationagent

import (
	"context"
	"embed"
	"fmt"
	"os"

	"google.golang.org/adk/agent"
	"google.golang.org/adk/model"
	"google.golang.org/adk/tool"

	"ril.api-ia/internal/agent/subagents/shared"
	"ril.api-ia/internal/agent/subagents/shared/tools"
	"ril.api-ia/internal/infrastructure/repository/tree_agent"
)

//go:embed instructions/*.tmpl
var instructionFiles embed.FS

//go:embed all:skills
var skillsFiles embed.FS

func NewProfessionalizationAgent(m model.LLM, treeManager *tree_agent.TreeCacheManager) (agent.Agent, error) {
	ctx := context.Background()

	saveUserMemory, err := tools.NewSaveMemoryTool()
	if err != nil {
		return nil, fmt.Errorf("creating save_user_memory tool: %w", err)
	}

	getUserMemory, err := tools.NewGetMemoryTool()
	if err != nil {
		return nil, fmt.Errorf("creating get_user_memory tool: %w", err)
	}

	useRagDocument, err := tools.NewRagTool("rilia-professionalization-rag-agent_1784216614773_gcs_store")
	if err != nil {
		return nil, fmt.Errorf("creating rag tool: %w", err)
	}

	cfg := shared.AgentConfig{
		Name:             os.Getenv("AGENT_PROFESSIONALIZATION_NAME"),
		DomainPrefix:     "professionalization",
		SearchCategory:   os.Getenv("AGENT_PROFESSIONALIZATION_SEARCH_CATEGORY"),
		Description:      "Agente especializado en acompañar a gobiernos locales en la profesionalización de su gestión, fortaleciendo la planificación estratégica, la toma de decisiones basada en información y la capacidad de implementación. Su función es empujar a los municipios a avanzar: ordenar la planificación, generar información, priorizar lo que importa y ejecutar cambios concretos. Para eso, utiliza el conocimiento experto y las metodologías de planificación desarrolladas por RIL, adaptándolas al contexto específico de cada municipio para ofrecer recomendaciones personalizadas y accionables.",
		Model:            m,
		TreeManager:      treeManager,
		InstructionFiles: instructionFiles,
		SkillsFiles:      skillsFiles,
		DomainTools: []tool.Tool{
			saveUserMemory,
			getUserMemory,
			useRagDocument,
		},
	}

	return shared.NewDomainAgent(ctx, cfg)
}
