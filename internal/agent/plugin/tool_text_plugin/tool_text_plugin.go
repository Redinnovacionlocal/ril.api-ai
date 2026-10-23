package tool_text_plugin

import (
	"fmt"
	"log/slog"
	"strings"

	"google.golang.org/adk/agent"
	"google.golang.org/adk/model"
	"google.golang.org/adk/plugin"
	"google.golang.org/genai"
)

const askContextTool = "ask_context_agent"

const introOnlyInstruction = `El usuario ya está viendo las preguntas de ask_context_agent en pantalla.
Escribí solo una introducción breve (1-2 líneas) de por qué necesitás más contexto.
No repitas, no enumeres ni resumas las preguntas ni sus opciones.`

const noRepeatInstruction = `En este turno ya le escribiste al usuario este texto, que ya está en pantalla:
<ya_escrito>
%s
</ya_escrito>
No lo repitas ni lo reformules, y no vuelvas a anunciar lo que ibas a hacer. Seguí desde ahí con lo nuevo.`

// New controls the model calls that follow a tool response, so the user does
// not see the same text twice in one turn.
func New(name string) (*plugin.Plugin, error) {
	if name == "" {
		name = "tool_text_plugin"
	}
	return plugin.New(plugin.Config{
		Name:                name,
		BeforeModelCallback: beforeModel,
	})
}

func beforeModel(ctx agent.CallbackContext, req *model.LLMRequest) (*model.LLMResponse, error) {
	n := len(req.Contents)
	if n < 2 || !hasFunctionResponse(req.Contents[n-1]) {
		return nil, nil
	}
	if answersAskContext(req.Contents[n-1]) {
		// The introduction was already written earlier in the turn: end it
		// without calling the model. Nil content makes ADK end it silently.
		if writtenThisTurn(req.Contents) != "" {
			slog.InfoContext(ctx, "tool_text: summarization skipped, intro already sent",
				"agent", ctx.AgentName(), "session_id", ctx.SessionID())
			return &model.LLMResponse{}, nil
		}
		addInstruction(req, introOnlyInstruction)
		return nil, nil
	}
	if written := writtenThisTurn(req.Contents); written != "" {
		addInstruction(req, fmt.Sprintf(noRepeatInstruction, written))
	}
	return nil, nil
}

func addInstruction(req *model.LLMRequest, text string) {
	if req.Config == nil {
		req.Config = &genai.GenerateContentConfig{}
	}
	if req.Config.SystemInstruction == nil {
		req.Config.SystemInstruction = &genai.Content{}
	}
	req.Config.SystemInstruction.Parts = append(req.Config.SystemInstruction.Parts, genai.NewPartFromText(text))
}

// writtenThisTurn returns the text the model already showed since the last
// user message.
func writtenThisTurn(contents []*genai.Content) string {
	var parts []string
	for i := len(contents) - 1; i >= 0; i-- {
		c := contents[i]
		if c == nil {
			continue
		}
		if c.Role == genai.RoleUser && !hasFunctionResponse(c) {
			break
		}
		if c.Role == genai.RoleModel {
			if t := text(c); t != "" {
				parts = append([]string{t}, parts...)
			}
		}
	}
	return strings.Join(parts, "\n\n")
}

// answersAskContext reports whether content carries the questions returned by
// ask_context_agent. A skipped call returns no questions and is left alone.
func answersAskContext(content *genai.Content) bool {
	for _, p := range content.Parts {
		if p == nil || p.FunctionResponse == nil || p.FunctionResponse.Name != askContextTool {
			continue
		}
		if _, ok := p.FunctionResponse.Response["question_blocks"]; ok {
			return true
		}
	}
	return false
}

func hasFunctionResponse(content *genai.Content) bool {
	if content == nil {
		return false
	}
	for _, p := range content.Parts {
		if p != nil && p.FunctionResponse != nil {
			return true
		}
	}
	return false
}

func text(content *genai.Content) string {
	if content == nil {
		return ""
	}
	var sb strings.Builder
	for _, p := range content.Parts {
		if p != nil && !p.Thought {
			sb.WriteString(p.Text)
		}
	}
	return strings.TrimSpace(sb.String())
}
