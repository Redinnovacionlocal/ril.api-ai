package securityagent

import (
	"context"
	"embed"
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

func NewSecurityAgent(m model.LLM, treeManager *tree_agent.TreeCacheManager) (agent.Agent, error) {
	ctx := context.Background()

	saveUserMemory, _ := tools.NewSaveMemoryTool()
	getUserMemory, _ := tools.NewGetMemoryTool()
	useRagDocument, _ := tools.NewRagTool("ril-security-knowledge_1775562649372_gcs_store")

	cfg := shared.AgentConfig{
		Name:             os.Getenv("AGENT_SECURITY_NAME"),
		DomainPrefix:     "security",
		Description:      "Agente especializado en acompañar a municipios en la mejora de su gestión de seguridad ciudadana. Su función es empujar a los municipios a avanzar: completar datos, mejorar lo que ya tienen, priorizar lo que importa, y ejecutar cambios concretos. Para eso, utiliza el conocimiento experto del árbol de criterios de calidad construido por los facilitadores de RIL, y lo aplica al contexto específico de cada municipio para ofrecer recomendaciones personalizadas y accionables.",
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
