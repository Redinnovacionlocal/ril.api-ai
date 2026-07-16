package agent

import (
	"context"
	"log"
	"os"
	"strconv"

	"github.com/googleapis/mcp-toolbox-sdk-go/tbadk"
	"github.com/jmoiron/sqlx"
	"google.golang.org/adk/agent"
	"google.golang.org/adk/agent/llmagent"
	"google.golang.org/adk/model/gemini"
	"google.golang.org/adk/tool"
	"google.golang.org/adk/tool/functiontool"
	"google.golang.org/genai"
	"ril.api-ia/internal/agent/subagents"
	"ril.api-ia/internal/agent/tools"
	"ril.api-ia/internal/domain/entity"
)

func NewRilAgent(ctx context.Context, db *sqlx.DB, toolboxClient tbadk.ToolboxClient, genaiClient *genai.Client) (agent.Agent, error) {
	// Overall configuration
	m, err := gemini.NewModel(ctx, os.Getenv("AGENT_MODEL"), nil)
	if err != nil {
		log.Fatal(err)
	}
	temperature64, _ := strconv.ParseFloat(os.Getenv("TEMPERATURE"), 32)
	temperature32 := float32(temperature64)
	maxOutputTokens, _ := strconv.Atoi(os.Getenv("MAX_OUTPUT_TOKENS"))
	contentConfiguration := &genai.GenerateContentConfig{
		Temperature:     &temperature32,
		MaxOutputTokens: int32(maxOutputTokens),
		SafetySettings: []*genai.SafetySetting{
			{
				Category:  genai.HarmCategoryDangerousContent,
				Threshold: genai.HarmBlockThresholdBlockMediumAndAbove,
			},
		},
	}

	toolboxTool, err := toolboxClient.LoadTool("get_user_data_by_id", ctx)
	getCertificateToolboxTool, _ := toolboxClient.LoadTool("get_certificate_by_id_team", ctx)
	getAllCertificateToolboxTool, _ := toolboxClient.LoadTool("get_all_certificates_active", ctx)
	getAllQuestionnaireActive, _ := toolboxClient.LoadTool("get_all_questionnaire_active", ctx)
	getQuestionnarieQuestionsByIdOrName, _ := toolboxClient.LoadTool("get_questionnarie_questions_by_id_or_name", ctx)
	getRilAliances, _ := toolboxClient.LoadTool("get_ril_aliances", ctx)
	getRilAliancesByAccountName, _ := toolboxClient.LoadTool("get_ril_aliances_by_account_name", ctx)
	getRilAliancesByYear, _ := toolboxClient.LoadTool("get_ril_aliances_by_year", ctx)
	getRilStaff, _ := toolboxClient.LoadTool("get_ril_staff", ctx)
	getEvaluationByName, _ := toolboxClient.LoadTool("get_evaluation_by_name", ctx)

	// Custom tools
	toolGenerateDocument, _ := functiontool.New(functiontool.Config{
		Name:        "generate_document",
		Description: "Genera un documento a partir de un prompt específico. El prompt debe incluir instrucciones claras sobre el formato, la estructura y el contenido esperado del documento. Esta herramienta es ideal para crear informes, resúmenes ejecutivos, propuestas o cualquier otro tipo de documento que requiera una presentación profesional y coherente.",
	}, tools.GenerateDocumentsToolFunc)
	if err != nil {
		log.Fatalf("Failed to load tool: %v", err)
	}
	toolGetUserMemory, err := functiontool.New(functiontool.Config{
		Name:        "get_user_memory",
		Description: "Recupera la memoria acumulada del usuario sobre su municipio. Devuelve datos concretos aportados por el usuario, oportunidades de mejora identificadas y contexto relevante que se ha registrado en conversaciones anteriores. Esta herramienta es esencial para mantener la continuidad y personalización del acompañamiento, permitiendo al agente recordar lo que ya se sabe sobre el municipio y evitar pedir información redundante.",
	}, tools.GetUserMemoryToolFunc)

	saveMetadataFunc := func(toolCtx tool.Context, metadata *genai.GroundingMetadata) {
		sessionID := toolCtx.SessionID()
		if sessionID != "" {
			entity.GroundingCache.Store(sessionID, metadata)
		} else {
			log.Println(">> ERROR: La herramienta no recibió un SessionID")
		}
	}

	ragModelName := os.Getenv("AGENT_RAG_MODEL")
	ragProxyTool, err := subagents.NewRagProxyTool(ctx, genaiClient, ragModelName, saveMetadataFunc)
	if err != nil {
		log.Fatalf("Failed to create RAG proxy tool: %v", err)
	}
	return llmagent.New(llmagent.Config{
		Name:                  "rilia_agent",
		Description:           "Eres un asistente especialista en todo lo relacionado al ambito público. Ayudas a los usuarios a encontrar información relevante y precisa sobre estos temas, utilizando un lenguaje claro y accesible.",
		Instruction:           SystemInstruction,
		GenerateContentConfig: contentConfiguration,
		GlobalInstruction:     GlobalInstruction,
		Model:                 m,
		Tools: []tool.Tool{
			toolGenerateDocument,
			toolGetUserMemory,
			&toolboxTool,
			&getCertificateToolboxTool,
			&getAllCertificateToolboxTool,
			&getAllQuestionnaireActive,
			&getQuestionnarieQuestionsByIdOrName,
			&getRilAliances,
			&getRilAliancesByAccountName,
			&getRilAliancesByYear,
			&getRilStaff,
			&getEvaluationByName,
			ragProxyTool,
		},
	})
}
