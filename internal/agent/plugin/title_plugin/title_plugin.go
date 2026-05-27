package title_plugin

import (
	"context"
	"errors"
	"fmt"

	"google.golang.org/adk/agent"
	"google.golang.org/adk/model"
	"google.golang.org/adk/plugin"
	"google.golang.org/adk/session"
	"google.golang.org/genai"
)

const prompt = `Genera un título conciso y descriptivo (máximo 5 palabras) que capture el tema principal o la pregunta.
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
		Título:`

const LLMModel = "gemini-2.5-flash-lite"

type titlePlugin struct {
	name   string
	client *genai.Client
}

// New creates a new instance of the setTitlePlugin.
// It initializes the GenAI client and sets up the plugin configuration with the AfterModel callback.
func New(ctx context.Context, name string) (*plugin.Plugin, error) {
	if name == "" {
		name = "set_title_plugin"
	}
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		Backend:  genai.BackendVertexAI,
		Location: "global",
	})
	if err != nil {
		return nil, fmt.Errorf("error creating GenAI client: %w", err)
	}
	titlePlugin := &titlePlugin{
		name:   name,
		client: client,
	}
	return plugin.New(
		plugin.Config{
			Name:               name,
			AfterModelCallback: titlePlugin.afterModel,
		})
}

func (s *titlePlugin) afterModel(ctx agent.CallbackContext, llmResponse *model.LLMResponse, llmResponseError error) (*model.LLMResponse, error) {
	_, err := ctx.State().Get("title")
	if errors.Is(err, session.ErrStateKeyNotExist) && llmResponse.Partial == false {
		fmt.Printf("Generating title for session: %v\n", ctx.SessionID())
		if llmResponse.Content.Role == genai.RoleModel &&
			llmResponse.Content.Parts[0].Text != "" {
			modelContent := llmResponse.Content.Parts[0].Text
			title, err := s.generateTitle(ctx, ctx.UserContent().Parts[0].Text, modelContent)
			if err != nil {
				fmt.Printf("Error generating title: %v\n", err)
				return llmResponse, nil
			}
			err = ctx.State().Set("title", title)
			if err != nil {
				fmt.Printf("Error setting title: %v\n", err)
			}
			fmt.Printf("Change session Title to: %v\n", title)
		}

	}
	return nil, nil

}

func (s *titlePlugin) generateTitle(ctx agent.CallbackContext, userContent string, modelContent string) (string, error) {
	temperature := float32(0.5)
	result, err := s.client.Models.GenerateContent(ctx, LLMModel,
		genai.Text(fmt.Sprintf(prompt, userContent, modelContent)),
		&genai.GenerateContentConfig{
			Temperature:     &temperature,
			MaxOutputTokens: 20,
		},
	)
	if err != nil {
		fmt.Printf("Error generating model: %v\n", err)
		return "", err
	}
	title := ""
	if len(result.Candidates) > 0 && len(result.Candidates[0].Content.Parts) > 0 {
		title = result.Candidates[0].Content.Parts[0].Text
	}
	return title, err
}
