package ask_context_plugin

import (
	"log/slog"
	"strings"

	"google.golang.org/adk/agent"
	"google.golang.org/adk/model"
	"google.golang.org/adk/plugin"
	"google.golang.org/genai"
)

const toolName = "ask_context_agent"

const introOnlyInstruction = `El usuario ya está viendo las preguntas de ask_context_agent en pantalla.
Escribí solo una introducción breve (1-2 líneas) de por qué necesitás más contexto.
No repitas, no enumeres ni resumas las preguntas ni sus opciones.`

// New controls the model call that follows an ask_context_agent response, so
// the user sees a single introduction and the questions only once, in the form.
func New(name string) (*plugin.Plugin, error) {
	if name == "" {
		name = "ask_context_plugin"
	}
	return plugin.New(plugin.Config{
		Name:                name,
		BeforeModelCallback: beforeModel,
	})
}

func beforeModel(ctx agent.CallbackContext, req *model.LLMRequest) (*model.LLMResponse, error) {
	n := len(req.Contents)
	if n < 2 || !answersAskContext(req.Contents[n-1]) {
		return nil, nil
	}
	// The introduction was already written next to the tool call: end the turn
	// without calling the model. Nil content makes ADK end the turn silently.
	if hasText(req.Contents[n-2]) {
		slog.InfoContext(ctx, "ask_context: summarization skipped, intro already sent",
			"agent", ctx.AgentName(), "session_id", ctx.SessionID())
		return &model.LLMResponse{}, nil
	}
	if req.Config == nil {
		req.Config = &genai.GenerateContentConfig{}
	}
	if req.Config.SystemInstruction == nil {
		req.Config.SystemInstruction = &genai.Content{}
	}
	req.Config.SystemInstruction.Parts = append(req.Config.SystemInstruction.Parts, genai.NewPartFromText(introOnlyInstruction))
	return nil, nil
}

// answersAskContext reports whether content carries the questions returned by
// ask_context_agent. A skipped call returns no questions and is left alone.
func answersAskContext(content *genai.Content) bool {
	if content == nil {
		return false
	}
	for _, p := range content.Parts {
		if p == nil || p.FunctionResponse == nil || p.FunctionResponse.Name != toolName {
			continue
		}
		if _, ok := p.FunctionResponse.Response["question_blocks"]; ok {
			return true
		}
	}
	return false
}

func hasText(content *genai.Content) bool {
	if content == nil {
		return false
	}
	for _, p := range content.Parts {
		if p != nil && !p.Thought && strings.TrimSpace(p.Text) != "" {
			return true
		}
	}
	return false
}
