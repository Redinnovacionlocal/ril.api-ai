package callbacks

import (
	"fmt"
	"log"

	"google.golang.org/adk/agent"
	"google.golang.org/adk/model"
	"google.golang.org/genai"
)

func SetTitleSessionBeforeCallback(ctx agent.CallbackContext, llmResponse *model.LLMResponse, llmResponseError error) (*model.LLMResponse, error) {
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
	m := "gemini-2.5-flash"
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
