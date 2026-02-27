package agent

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/adk/agent"
	memory2 "google.golang.org/adk/memory"
	"google.golang.org/adk/model"
	"google.golang.org/adk/session"
	"google.golang.org/adk/tool"
	"google.golang.org/adk/tool/functiontool"
	"google.golang.org/genai"
)

// MemorySearchTool is a custom tool for searching the agent's memory.
func memorySearchToolFunc(tctx tool.Context, args Args) (Result, error) {
	fmt.Printf("Tool: Searching memory for query: '%s'\n", args.Query)

	searchResults, err := tctx.SearchMemory(context.Background(), args.Query)
	if err != nil {
		log.Printf("Error searching memory: %v", err)
		return Result{}, fmt.Errorf("failed memory search")
	}

	// FIX: Initialize with an empty slice instead of leaving it nil
	results := []string{}

	for _, res := range searchResults.Memories {
		if res.Content != nil {
			for _, part := range res.Content.Parts {
				if part.Text != "" {
					results = append(results, part.Text)
				}
			}
		}
	}

	// Now returns [] instead of null when no results are found
	return Result{Results: results}, nil
}
func addSessionToMemory(sessionService session.Service, memoryService memory2.Service) agent.AfterAgentCallback {
	return func(ctx agent.CallbackContext) (*genai.Content, error) {
		fmt.Println("Add session to memory callback executed")
		sessionID, _ := sessionService.Get(ctx,
			&session.GetRequest{SessionID: ctx.SessionID(), UserID: ctx.UserID(), AppName: ctx.AppName()},
		)
		sessionInstance := sessionID.Session
		err := memoryService.AddSession(ctx, sessionInstance)
		if err != nil {
			fmt.Printf("failed to save to memory: %v\n", err)
		}
		return nil, nil
	}
}
func setTitleOfSession(ctx agent.CallbackContext, llmResponse *model.LLMResponse, llmResponseError error) (*model.LLMResponse, error) {
	hasTitle, _ := ctx.State().Get("title")
	if hasTitle != nil {
		return llmResponse, nil
	}
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		Backend: genai.BackendVertexAI,
	})
	if err != nil {
		log.Fatal()
	}

	temperature := float32(0.5)
	var modelResponse, userContent string
	if llmResponse.Content.Role == genai.RoleModel {
		if len(llmResponse.Content.Parts) > 0 {
			for _, part := range llmResponse.Content.Parts {
				modelResponse += part.Text
			}
		}
	}

	userContent += ctx.UserContent().Parts[0].Text
	m := "gemini-2.5-flash-lite"
	prompt := fmt.Sprintf(`Genera un título conciso y descriptivo (máximo 5 palabras) que capture el tema principal o la pregunta.

		Reglas:
		- Sin signos de puntuación
		- Sin prefijos como "Título:", "Title:", o similares
		- Usa mayúsculas iniciales en palabras principales
		- Sé específico y descriptivo
		- Evita palabras genéricas como "Chat", "Conversación", "Discusión"
		- Enfócate en el tema o acción principal
		- Titulo humano y atractivo

		Ejemplos:
		- Usuario: "¿Cuáles son las mejores prácticas para la gestión de residuos en ciudades pequeñas?"
		  Título: Gestión de Residuos en Ciudades Pequeñas
		- Usuario: "Necesito ideas sobre cómo mejorar la participación ciudadana en proyectos locales."
		  Título: Mejora de la Participación Ciudadana Local	
		Mensaje del usuario: %s
		Respuesta del asistente: %s
		
		Título:`, userContent, modelResponse)

	result, err := client.Models.GenerateContent(ctx, m,
		genai.Text(prompt),
		&genai.GenerateContentConfig{
			Temperature:     &temperature,
			MaxOutputTokens: 20,
		},
	)
	if err != nil {
		log.Fatal("Error generating session title", err)
	}
	if len(result.Candidates) > 0 && len(result.Candidates[0].Content.Parts) > 0 {
		text := result.Candidates[0].Content.Parts[0].Text
		err = ctx.State().Set("title", text)
		if err != nil {
			log.Fatal(err)
		}
	}
	return llmResponse, nil
}

var memorySearchTool, _ = functiontool.New(
	functiontool.Config{
		IsLongRunning: false,
		Name:          "search_past_conversations",
		Description:   "Busca en el historial de conversaciones pasadas del usuario para encontrar información relevante que pueda ayudar a responder su consulta actual."},
	memorySearchToolFunc,
)

type Args struct {
	Query string `json:"query" jsonschema:"The query to search for in the memory."`
}

// Result defines the output structure for the memory search tool.
type Result struct {
	Results []string `json:"results"`
}
