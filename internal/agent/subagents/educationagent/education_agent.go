package educationagent

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

func NewEducationAgent(m model.LLM, treeManager *tree_agent.TreeCacheManager) (agent.Agent, error) {
	ctx := context.Background()
	saveUserMemory, err := tools.NewSaveMemoryTool()
	if err != nil {
		return nil, fmt.Errorf("creating save_user_memory tool: %w", err)
	}

	getUserMemory, err := tools.NewGetMemoryTool()
	if err != nil {
		return nil, fmt.Errorf("creating get_user_memory tool: %w", err)
	}

	//hardcoded
	useRagDocument, err := tools.NewRagTool("rilia-education-rag-agent_1786727607100_gcs_store")
	if err != nil {
		return nil, fmt.Errorf("creating rag tool: %w", err)
	}

	cfg := shared.AgentConfig{
		Name:             os.Getenv("AGENT_EDUCATION_NAME"),
		DomainPrefix:     "education",
		SearchCategory:   os.Getenv("AGENT_EDUCATION_SEARCH_CATEGORY"),
		Description:      "Agente especializado en acompañar a municipios en la mejora de su gestión educativa. Su función es empujar a los municipios a avanzar: ordenar el área, generar información, mejorar lo que ya tienen, priorizar lo que importa y ejecutar acciones concretas. Para eso, utiliza el conocimiento experto sobre gestión educativa y lo aplica al contexto específico de cada municipio para ofrecer recomendaciones personalizadas, accionables y orientadas a la implementación.",
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
