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
You are an assistant designed to ask users a series of questions in order to gather the necessary information for a task. 
Your goal is to ask clear and concise questions that will help you understand the user's needs and provide the best possible assistance.
# Rules:
- Always ask questions in a clear and concise manner.
- Always provide a set of possible answers for each question to guide the user.
- Set "multi_select" to true when the question allows selecting more than one answer (e.g. "Which features do you need?"), and false when only one answer is expected (e.g. "What is your preferred language?").
- If you have enough information to complete the task, do not ask more questions and proceed to provide assistance.
`

func NewAskContextAgent(ctx context.Context) (agent.Agent, error) {
	m, err := gemini.NewModel(ctx, os.Getenv("ASK_CONTEXT_AGENT_MODEL"), nil)
	if err != nil {
		log.Fatal(err)
	}

	AskComponentOut := &genai.Schema{
		Type:        genai.TypeObject,
		Description: "Schema for the output of the AskContextAgent. It contains a list of question blocks, where each block includes a question, a set of possible answers, and the order in which the question should be asked.",
		Properties: map[string]*genai.Schema{
			"question_blocks": {
				Type:        genai.TypeArray,
				Description: "The list of question blocks to ask the user in order to gather the necessary information for the task.",
				Items: &genai.Schema{
					Type:        genai.TypeObject,
					Description: "A single question block containing a question, possible answers, and sort order.",
					Properties: map[string]*genai.Schema{
						"question": {
							Type:        genai.TypeString,
							Description: "The question to ask the user.",
						},
						"answers": {
							Type:        genai.TypeArray,
							Description: "The possible answers to the question.",
							Items: &genai.Schema{
								Type: genai.TypeString,
							},
						},
						"sort": {
							Type:        genai.TypeInteger,
							Description: "The order in which the question should be asked.",
						},
						"multi_select": {
							Type:        genai.TypeBoolean,
							Description: "Whether the user can select more than one answer for this question.",
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
		Description: "An agent designed to ask users a series of questions in order to gather the necessary information for a task.",
		Instruction: SystemInstruction,
		Model:       m,
		GenerateContentConfig: &genai.GenerateContentConfig{
			Temperature:     &temperature,
			MaxOutputTokens: 8000,
		},
		OutputSchema: AskComponentOut,
	})
}
