package securityagent

import (
	"bytes"
	"context"
	"embed"
	"fmt"
	"os"
	"strings"
	"text/template"

	"github.com/googleapis/mcp-toolbox-sdk-go/tbadk"
	"google.golang.org/adk/agent"
	"google.golang.org/adk/agent/llmagent"
	"google.golang.org/adk/model"
	"google.golang.org/adk/tool"
	"google.golang.org/adk/tool/agenttool"
	"google.golang.org/adk/tool/functiontool"
	"google.golang.org/adk/tool/geminitool"
	"google.golang.org/genai"
	agent2 "ril.api-ia/internal/agent"
	"ril.api-ia/internal/agent/subagents/askcontextagent"
	tools2 "ril.api-ia/internal/agent/subagents/securityagent/tools"
	"ril.api-ia/internal/agent/tools"
	"ril.api-ia/internal/infrastructure/repository/tree_agent"
)

//go:embed instructions/*.tmpl
var instructionFiles embed.FS

type PromptData struct {
	Tags       []string
	Dimensions []string
}

func buildSystemInstruction(data PromptData) (string, error) {
	funcMap := template.FuncMap{
		"join": strings.Join,
	}

	tmpl, err := template.New("").Funcs(funcMap).ParseFS(instructionFiles, "instructions/*.tmpl")
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	if err := tmpl.ExecuteTemplate(&buf, "main_instruction.tmpl", data); err != nil {
		return "", err
	}

	return buf.String(), nil
}

func NewSecurityAgent(m model.LLM, toolboxClient tbadk.ToolboxClient, treeRepo *tree_agent.TreeQuestionRepository) (agent.Agent, error) {
	securityAgentName := os.Getenv("AGENT_SECURITY_NAME")
	ctx := context.Background()

	dimensions, err := treeRepo.GetDimensions()
	if err != nil {
		return nil, err
	}

	tags, err := treeRepo.GetTags()
	if err != nil {
		return nil, err
	}

	SystemInstruction, err := buildSystemInstruction(PromptData{
		Dimensions: dimensions,
		Tags:       tags,
	})
	if err != nil {
		return nil, err
	}

	lookupTreeQuestions, err := toolboxClient.LoadTool("lookup_tree_questions", ctx)
	if err != nil {
		return nil, fmt.Errorf("Failed to load lookup_tree_questions: %w", err)
	}

	toolGenerateDocument, _ := functiontool.New(functiontool.Config{
		Name:        "generate_document",
		Description: "Genera un documento a partir de un prompt específico. El prompt debe incluir instrucciones claras sobre el formato, la estructura y el contenido esperado del documento. Esta herramienta es ideal para crear informes, resúmenes ejecutivos, propuestas o cualquier otro tipo de documento que requiera una presentación profesional y coherente.",
	}, tools.GenerateDocumentsToolFunc)
	saveUserMemory, err := functiontool.New(functiontool.Config{
		Name:        "save_user_memory",
		Description: "Guarda en la memoria del usuario todo lo que el municipio aporta durante la conversación. Permite registrar datos concretos sobre el municipio, oportunidades de mejora identificadas y contexto relevante para personalizar recomendaciones futuras. Es fundamental para construir una memoria acumulada que permita un acompañamiento cada vez más adaptado y efectivo.",
	}, tools2.SaveUserMemoryToolFunc)
	if err != nil {
		return nil, err
	}

	getUserMemory, err := functiontool.New(functiontool.Config{
		Name:        "get_user_memory",
		Description: "Recupera la memoria acumulada del usuario sobre su municipio. Devuelve datos concretos aportados por el usuario, oportunidades de mejora identificadas y contexto relevante que se ha registrado en conversaciones anteriores. Esta herramienta es esencial para mantener la continuidad y personalización del acompañamiento, permitiendo al agente recordar lo que ya se sabe sobre el municipio y evitar pedir información redundante.",
	}, tools2.GetUserMemoryToolFunc)

	UseRagDocument, err := functiontool.New(functiontool.Config{
		Name:        "rilia_security_rag_agent",
		Description: "Permite al agente utilizar un documento recuperado del RAG como parte de su respuesta al usuario. El agente puede extraer información relevante del documento para enriquecer sus recomendaciones y respuestas, asegurando que el conocimiento específico de las bases de RIL se integre de manera efectiva en la conversación con el municipio.",
	}, tools2.UseRagVertexAISearchToolFunc)

	AskContext, err := askcontextagent.NewAskContextAgent(ctx)
	if err != nil {
		return nil, err
	}
	return llmagent.New(llmagent.Config{
		Name:              securityAgentName,
		Instruction:       SystemInstruction,
		GlobalInstruction: agent2.GlobalInstruction,
		Description:       "Agente especializado en acompañar a municipios en la mejora de su gestión de seguridad ciudadana. Su función es empujar a los municipios a avanzar: completar datos, mejorar lo que ya tienen, priorizar lo que importa, y ejecutar cambios concretos. Para eso, utiliza el conocimiento experto del árbol de criterios de calidad construido por los facilitadores de RIL, y lo aplica al contexto específico de cada municipio para ofrecer recomendaciones personalizadas y accionables.",
		Model:             m,
		Tools: []tool.Tool{
			UseRagDocument,
			toolGenerateDocument,
			saveUserMemory,
			getUserMemory,
			agenttool.New(AskContext, &agenttool.Config{
				SkipSummarization: true,
			}),
			&lookupTreeQuestions,
		},
	})
}

func NewSecurityRagAgent(m model.LLM) (agent.Agent, error) {
	maxRagResults := int32(10)
	return llmagent.New(llmagent.Config{
		Name:        "rilia_security_rag_agent",
		Description: "Agente especializado en acompañar a municipios en la mejora de su gestión de seguridad ciudadana, con acceso a bases de conocimiento específicas de RIL. Utiliza herramientas de búsqueda semántica para obtener información relevante de guías, normativas, casos de ciudades, modelos y templates relacionados con seguridad ciudadana. Su función es empujar a los municipios a avanzar: completar datos, mejorar lo que ya tienen, priorizar lo que importa, y ejecutar cambios concretos, apoyándose en el conocimiento específico disponible en las bases de RIL.",
		Instruction: "Agente especializado en acompañar a municipios en la mejora de su gestión de seguridad ciudadana, con acceso a bases de conocimiento específicas de RIL. Utiliza herramientas de búsqueda semántica para obtener información relevante de guías, normativas, casos de ciudades, modelos y templates relacionados con seguridad ciudadana. Su función es empujar a los municipios a avanzar: completar datos, mejorar lo que ya tienen, priorizar lo que importa, y ejecutar cambios concretos, apoyándose en el conocimiento específico disponible en las bases de RIL.",
		Model:       m,
		Tools: []tool.Tool{
			geminitool.New("rag_security_knowledge",
				"Get security knowledge from RIL's knowledge bases. Use this tool to search for relevant information on guides, regulations, city cases, models, and templates related to citizen security. This tool is essential for providing informed recommendations and actionable insights to municipalities based on the specific knowledge available in RIL's databases.",
				&genai.Tool{
					Retrieval: &genai.Retrieval{
						VertexAISearch: &genai.VertexAISearch{
							MaxResults: &maxRagResults,
							Datastore:  "projects/ril-admin/locations/global/collections/default_collection/dataStores/ril-security-knowledge_1775562649372_gcs_store",
						},
					},
				}),
		},
	})
}
