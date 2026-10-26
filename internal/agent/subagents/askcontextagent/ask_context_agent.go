package askcontextagent

import (
	"context"
	"log"
	"os"

	"google.golang.org/adk/agent"
	"google.golang.org/adk/agent/llmagent"
	"google.golang.org/adk/model/gemini"
	"google.golang.org/genai"
)

const SystemInstruction = `
Sos un asistente que le hace al usuario una serie de preguntas para reunir la información necesaria para una tarea.
Tu objetivo es hacer preguntas claras y concisas que ayuden a entender lo que necesita el usuario y a darle la mejor asistencia posible.

# Idioma — CRÍTICO:
- Escribí todas las preguntas y todas las opciones en el idioma del usuario final, aunque el pedido que recibas esté escrito en otro idioma.
- Si el pedido indica el idioma del usuario, usá ese. Si no, usá el idioma en el que el usuario le escribe al asistente según el contexto del pedido.
- Nunca escribas en inglés salvo que el usuario escriba en inglés.

# Terminología:
- Escribí "gobierno local" en lugar de "municipio" o "municipalidad", y "local" o "del gobierno local" en lugar de "municipal", aunque el pedido use esas palabras. En portugués, "governo local". Excepción: nombres propios de organismos, programas o normas.

# Reglas:
- Hacé siempre preguntas claras y concisas.
- Ofrecé siempre un conjunto de respuestas posibles para cada pregunta, para guiar al usuario.
- Poné "multi_select" en true cuando el pedido marca la pregunta con "(admite varias)" o cuando las opciones no son excluyentes y el usuario puede tener más de una (por ejemplo, "¿Qué funcionalidades necesitás?"), y en false cuando se espera una sola (por ejemplo, "¿Cuál es tu idioma preferido?").
- Si ya tenés información suficiente para completar la tarea, no hagas más preguntas y seguí con la asistencia.

# Opciones de respuesta cerradas — CRÍTICO:
- Cada respuesta en "answers" TIENE que ser una opción concreta, seleccionable y cerrada. Debajo de las opciones, la interfaz ya le muestra al usuario un campo aparte para escribir su propia respuesta.
- Por eso NUNCA incluyas opciones abiertas, genéricas o del tipo "completá vos", como "Otro", "Otro (especificá)", "Ninguna de las anteriores (explicá)", "Respuesta personalizada", o cualquier variante que implique que el usuario tiene que escribir algo: ese campo ya existe.
- Si no estás seguro de que tu lista de respuestas sea completa, pensá más sobre el tema y enumerá vos mismo las opciones concretas realistas: no le delegues ese trabajo al usuario con una opción abierta.
- No te cubras agregando una opción vaga que abarque todo. Cada respuesta tiene que ser algo que el usuario pueda elegir sin escribir nada más.
`

func NewAskContextAgent(ctx context.Context) (agent.Agent, error) {
	m, err := gemini.NewModel(ctx, os.Getenv("ASK_CONTEXT_AGENT_MODEL"), nil)
	if err != nil {
		log.Fatal(err)
	}

	AskComponentOut := &genai.Schema{
		Type:        genai.TypeObject,
		Description: "Salida del AskContextAgent: una lista de bloques de preguntas, cada uno con la pregunta, sus respuestas posibles y el orden en que se hace.",
		Properties: map[string]*genai.Schema{
			"question_blocks": {
				Type:        genai.TypeArray,
				Description: "Los bloques de preguntas para hacerle al usuario y reunir la información necesaria para la tarea.",
				Items: &genai.Schema{
					Type:        genai.TypeObject,
					Description: "Un bloque con una pregunta, sus respuestas posibles y su orden.",
					Properties: map[string]*genai.Schema{
						"question": {
							Type:        genai.TypeString,
							Description: "La pregunta para el usuario, en su idioma.",
						},
						"answers": {
							Type:        genai.TypeArray,
							Description: "Las respuestas posibles a la pregunta, en el idioma del usuario.",
							Items: &genai.Schema{
								Type: genai.TypeString,
							},
						},
						"sort": {
							Type:        genai.TypeInteger,
							Description: "El orden en que se hace la pregunta.",
						},
						"multi_select": {
							Type:        genai.TypeBoolean,
							Description: "Si el usuario puede elegir más de una respuesta para esta pregunta.",
						},
					},
					Required: []string{"question", "answers", "sort", "multi_select"},
				},
			},
		},
		Required: []string{"question_blocks"},
	}
	temperature := float32(0.3)
	return llmagent.New(llmagent.Config{
		Name:        "ask_context_agent",
		Description: "Agente que le hace al usuario una serie de preguntas para reunir la información necesaria para una tarea. Indicá en el pedido el idioma del usuario.",
		Instruction: SystemInstruction,
		Model:       formModel{LLM: m},
		GenerateContentConfig: &genai.GenerateContentConfig{
			Temperature:     &temperature,
			MaxOutputTokens: 8000,
		},
		OutputSchema: AskComponentOut,
	})
}
