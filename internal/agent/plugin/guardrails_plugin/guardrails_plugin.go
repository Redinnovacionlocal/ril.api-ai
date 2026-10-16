package guardrails_plugin

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"strings"
	"time"

	"google.golang.org/adk/agent"
	"google.golang.org/adk/model"
	"google.golang.org/adk/plugin"
	"google.golang.org/genai"
	"ril.api-ia/internal/infrastructure/env"
	"ril.api-ia/internal/infrastructure/observability"
)

const (
	rewriteModel   = "gemini-2.5-flash"
	maxRetries     = 1
	correctTimeout = 20 * time.Second
)

const keepSentinel = "SIN_CAMBIOS"

const rewritePrompt = `Sos un corrector de respuestas de un asistente.

Mensaje del usuario:
<mensaje_usuario>
%s
</mensaje_usuario>

Problemas detectados en la respuesta:
%s
Paso 1: leé el mensaje del usuario. Si pide explícitamente algo que justifica la respuesta tal como está (por ejemplo, responder en otro idioma o traducir un texto), respondé únicamente ` + keepSentinel + ` y nada más.

Paso 2: si no, corregí los problemas detectados:
- Mantené exactamente la misma información, estructura y formato markdown (títulos, listas, tablas, links).
- No agregues ni quites contenido, ni sumes comentarios o explicaciones sobre la corrección.
- No modifiques URLs, código ni nombres propios.
- Devolvé únicamente la respuesta corregida, sin las etiquetas <respuesta>.

<respuesta>
%s
</respuesta>`

var errKept = errors.New("user explicitly asked for this response")

type guardrailsPlugin struct {
	client *genai.Client
	rules  []Rule
}

// GUARDRAILS_ENABLED         (default: true)
// GUARDRAILS_DISABLED_RULES  comma-separated rule names
// GUARDRAILS_LANGUAGES       ISO 639-1 codes (default: es,pt,en)
func New(ctx context.Context, name string) (*plugin.Plugin, error) {
	if name == "" {
		name = "guardrails_plugin"
	}

	if env.GetOrDefault("GUARDRAILS_ENABLED", "true") != "true" {
		slog.Info("guardrails disabled (GUARDRAILS_ENABLED != true)")
		return plugin.New(plugin.Config{Name: name})
	}

	rules, err := enabledRules(env.GetOrDefault("GUARDRAILS_DISABLED_RULES", ""))
	if err != nil {
		return nil, err
	}

	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		Backend:  genai.BackendVertexAI,
		Location: "global",
	})
	if err != nil {
		return nil, fmt.Errorf("error creating GenAI client: %w", err)
	}

	g := &guardrailsPlugin{client: client, rules: rules}

	return plugin.New(plugin.Config{
		Name:               name,
		AfterModelCallback: g.afterModel,
	})
}

func enabledRules(disabled string) ([]Rule, error) {
	off := make(map[string]bool)
	for _, n := range strings.Split(disabled, ",") {
		if n = strings.TrimSpace(n); n != "" {
			off[n] = true
		}
	}
	var rules []Rule
	for _, f := range registeredRules() {
		if off[f.name] {
			slog.Info("guardrails: rule disabled", "guardrails.rule", f.name)
			continue
		}
		r, err := f.new()
		if err != nil {
			return nil, fmt.Errorf("error creating rule %q: %w", f.name, err)
		}
		rules = append(rules, r)
	}
	return rules, nil
}

func (g *guardrailsPlugin) afterModel(ctx agent.CallbackContext, resp *model.LLMResponse, respErr error) (*model.LLMResponse, error) {
	if respErr != nil || resp == nil || resp.Partial || resp.Content == nil {
		return nil, nil
	}
	responseText, ok := finalText(resp.Content)
	if !ok {
		return nil, nil
	}

	in := Input{UserText: contentText(ctx.UserContent()), ResponseText: responseText, State: ctx.State()}
	violations := g.check(in)
	if len(violations) == 0 {
		return nil, nil
	}

	logAttrs := []any{
		"agent", ctx.AgentName(),
		"invocation_id", ctx.InvocationID(),
		"session_id", ctx.SessionID(),
	}
	for _, v := range violations {
		slog.WarnContext(ctx, "guardrails: rule failed",
			append([]any{"guardrails.rule", v.Rule, "guardrails.detail", v.Detail}, logAttrs...)...)
	}

	start := time.Now()
	fixed, attempts, err := g.correct(ctx, in, violations)
	logAttrs = append(logAttrs,
		"guardrails.rules", ruleNames(violations),
		"guardrails.attempts", attempts,
		"guardrails.duration_ms", time.Since(start).Milliseconds())
	if errors.Is(err, errKept) {
		for _, v := range violations {
			if v.OnKeep != nil {
				v.OnKeep()
			}
		}
		slog.InfoContext(ctx, "guardrails: response kept at user request", logAttrs...)
		return nil, nil
	}
	if err != nil {
		slog.WarnContext(ctx, "guardrails: fallback to original response",
			append([]any{"guardrails.reason", err.Error()}, logAttrs...)...)
		return nil, nil
	}

	slog.InfoContext(ctx, "guardrails: response corrected", logAttrs...)
	return replaceText(resp, fixed), nil
}

func (g *guardrailsPlugin) check(in Input) []*Violation {
	var violations []*Violation
	for _, r := range g.rules {
		if v := r.Check(in); v != nil {
			violations = append(violations, v)
		}
	}
	return violations
}

func (g *guardrailsPlugin) correct(ctx context.Context, in Input, violations []*Violation) (string, int, error) {
	ctx, cancel := context.WithTimeout(ctx, correctTimeout)
	defer cancel()
	ctx, span := observability.Start(ctx, "guardrails.correct", observability.Attrs{
		"guardrails.rules":     ruleNames(violations),
		"gen_ai.request.model": rewriteModel,
	})
	defer span.End()

	text := in.ResponseText
	for attempt := 1; attempt <= maxRetries; attempt++ {
		fixed, err := g.rewrite(ctx, in.UserText, text, violations)
		if errors.Is(err, errKept) {
			return "", attempt, err
		}
		if err != nil {
			span.Fail(err)
			return "", attempt, fmt.Errorf("attempt %d failed: %w", attempt, err)
		}
		text = fixed
		violations = g.check(Input{UserText: in.UserText, ResponseText: text, State: in.State})
		if len(violations) == 0 {
			span.Set(observability.Attrs{"guardrails.attempts": attempt})
			return text, attempt, nil
		}
	}
	return "", maxRetries, fmt.Errorf("retries exhausted, still failing: %s", ruleNames(violations))
}

func (g *guardrailsPlugin) rewrite(ctx context.Context, userText, text string, violations []*Violation) (string, error) {
	var instructions strings.Builder
	for _, v := range violations {
		fmt.Fprintf(&instructions, "- %s\n", v.Instruction)
	}

	temperature := float32(0)
	result, err := g.client.Models.GenerateContent(ctx, rewriteModel,
		genai.Text(fmt.Sprintf(rewritePrompt, userText, instructions.String(), text)),
		&genai.GenerateContentConfig{Temperature: &temperature},
	)
	if err != nil {
		return "", err
	}
	out := strings.TrimSpace(result.Text())
	out = strings.TrimSpace(strings.TrimSuffix(strings.TrimPrefix(out, "<respuesta>"), "</respuesta>"))
	if out == "" {
		return "", errors.New("rewrite model returned an empty response")
	}
	if strings.Trim(out, " .`*\"'") == keepSentinel {
		return "", errKept
	}
	return out, nil
}

func finalText(content *genai.Content) (string, bool) {
	var sb strings.Builder
	for _, p := range content.Parts {
		if p == nil {
			continue
		}
		if p.FunctionCall != nil || p.FunctionResponse != nil {
			return "", false
		}
		if !p.Thought {
			sb.WriteString(p.Text)
		}
	}
	text := sb.String()
	return text, strings.TrimSpace(text) != ""
}

func contentText(content *genai.Content) string {
	if content == nil {
		return ""
	}
	var sb strings.Builder
	for _, p := range content.Parts {
		if p != nil && !p.Thought {
			sb.WriteString(p.Text)
		}
	}
	return sb.String()
}

func replaceText(resp *model.LLMResponse, text string) *model.LLMResponse {
	parts := make([]*genai.Part, 0, len(resp.Content.Parts))
	replaced := false
	for _, p := range resp.Content.Parts {
		if p == nil || p.Thought || p.Text == "" {
			parts = append(parts, p)
			continue
		}
		if !replaced {
			parts = append(parts, &genai.Part{Text: text, ThoughtSignature: p.ThoughtSignature})
			replaced = true
		}
	}

	newResp := *resp
	newResp.Content = &genai.Content{Role: resp.Content.Role, Parts: parts}
	return &newResp
}

func ruleNames(violations []*Violation) string {
	names := make([]string, len(violations))
	for i, v := range violations {
		names[i] = v.Rule
	}
	return strings.Join(names, ",")
}
