package girsuagent

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

func NewGirsuAgent(m model.LLM, treeManager *tree_agent.TreeCacheManager) (agent.Agent, error) {
	ctx := context.Background()

	saveUserMemory, err := tools.NewSaveMemoryTool()
	if err != nil {
		return nil, fmt.Errorf("creating save_user_memory tool: %w", err)
	}

	getUserMemory, err := tools.NewGetMemoryTool()
	if err != nil {
		return nil, fmt.Errorf("creating get_user_memory tool: %w", err)
	}

	useRagDocument, err := tools.NewRagTool("rilia-girsu-rag-agent_1783019204455_gcs_store")
	if err != nil {
		return nil, fmt.Errorf("creating rag tool: %w", err)
	}

	cfg := shared.AgentConfig{
		Name:             os.Getenv("AGENT_GIRSU_NAME"),
		DomainPrefix:     "girsu",
		SearchCategory:   os.Getenv("AGENT_GIRSU_SEARCH_CATEGORY"),
		Description:      "Agente especializado en acompañar a municipios en la mejora de su gestión integral de residuos sólidos urbanos, fortaleciendo la planificación, la eficiencia operativa y la sostenibilidad ambiental. Su función es empujar a los municipios a avanzar: completar datos, mejorar lo que ya tienen, priorizar lo que importa, y ejecutar cambios concretos. Para eso, utiliza el conocimiento experto del árbol de criterios de calidad construido por los facilitadores de RIL, y lo aplica al contexto específico de cada municipio para ofrecer recomendaciones personalizadas y accionables.",
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
