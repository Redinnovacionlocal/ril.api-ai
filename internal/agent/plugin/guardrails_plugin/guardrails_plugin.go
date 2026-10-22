package guardrails_plugin

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"strings"
	"sync"
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

const fragmentsPrompt = `Sos un corrector de respuestas de un asistente. Algunos fragmentos de la respuesta tienen problemas.

Mensaje del usuario:
<mensaje_usuario>
%s
</mensaje_usuario>

Problemas detectados:
%s
Paso 1: leé el mensaje del usuario. Si pide explícitamente algo que justifica los fragmentos tal como están (por ejemplo, responder en otro idioma o traducir un texto), devolvé keep en true.

Paso 2: si no, devolvé keep en false y cada fragmento corregido, en el mismo orden:
- Si el fragmento es razonamiento interno del asistente, devolvé un texto vacío para eliminarlo. Es razonamiento todo lo que el asistente se dice a sí mismo y no le diría al usuario: su plan, cómo va a armar la respuesta, el análisis del pedido, qué herramienta va a usar o qué información obtuvo, o cuando habla del usuario en tercera persona. Por ejemplo: "Let's synthesize an example for the user", "Now let's structure the response", "Great! We have retrieved documents from the RAG", "Wait, let's check the tool specifications".
- Si es contenido para el usuario, corregilo manteniendo el formato markdown, las URLs, el código y los nombres propios.
- Si una parte del fragmento es razonamiento y otra es contenido, eliminá solo el razonamiento.

Fragmentos:
%s`

var fragmentsSchema = &genai.Schema{
	Type: genai.TypeObject,
	Properties: map[string]*genai.Schema{
		"keep":      {Type: genai.TypeBoolean},
		"fragments": {Type: genai.TypeArray, Items: &genai.Schema{Type: genai.TypeString}},
	},
	Required: []string{"keep", "fragments"},
}

var errKept = errors.New("user explicitly asked for this response")

type guardrailsPlugin struct {
	client *genai.Client
	rules  []Rule
	// streams holds the *lineStream of each response being streamed.
	streams sync.Map
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
	key := streamKey(ctx)
	if respErr != nil || resp == nil {
		g.streams.Delete(key)
		return nil, nil
	}
	if resp.Partial {
		return g.afterPartial(ctx, key, resp), nil
	}
	if s, ok := g.streams.LoadAndDelete(key); ok {
		return g.afterStream(ctx, s.(*lineStream), resp), nil
	}
	if resp.Content == nil {
		return nil, nil
	}
	text, hasToolCall, ok := responseText(resp.Content)
	if !ok {
		return nil, nil
	}

	in := Input{UserText: contentText(ctx.UserContent()), ResponseText: text, State: ctx.State()}
	violations := g.check(in)
	if len(violations) == 0 {
		return nil, nil
	}

	attrs := logAttrs(ctx)
	for _, v := range violations {
		slog.WarnContext(ctx, "guardrails: rule failed",
			append([]any{"guardrails.rule", v.Rule, "guardrails.detail", v.Detail, "guardrails.tool_call", hasToolCall}, attrs...)...)
	}

	// Text next to a tool call is filler before the tool runs: drop it instead
	// of delaying the tool with a rewrite.
	if hasToolCall {
		slog.InfoContext(ctx, "guardrails: text removed from tool call",
			append([]any{"guardrails.rules", ruleNames(violations)}, attrs...)...)
		return replaceText(resp, ""), nil
	}

	start := time.Now()
	fixed, attempts, err := g.correct(ctx, in, violations)
	attrs = append(attrs,
		"guardrails.rules", ruleNames(violations),
		"guardrails.attempts", attempts,
		"guardrails.duration_ms", time.Since(start).Milliseconds())
	if errors.Is(err, errKept) {
		for _, v := range violations {
			if v.OnKeep != nil {
				v.OnKeep()
			}
		}
		slog.InfoContext(ctx, "guardrails: response kept at user request", attrs...)
		return nil, nil
	}
	if err != nil {
		slog.WarnContext(ctx, "guardrails: fallback to original response",
			append([]any{"guardrails.reason", err.Error()}, attrs...)...)
		return nil, nil
	}

	slog.InfoContext(ctx, "guardrails: response corrected", attrs...)
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
		var fixed string
		var err error
		if fragments := allFragments(violations); fragments != nil {
			fixed, err = g.rewriteFragments(ctx, in.UserText, text, fragments, violations)
		} else {
			fixed, err = g.rewrite(ctx, in.UserText, text, violations)
		}
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
	temperature := float32(0)
	result, err := g.client.Models.GenerateContent(ctx, rewriteModel,
		genai.Text(fmt.Sprintf(rewritePrompt, userText, instructions(violations), text)),
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

// rewriteFragments rewrites only the given fragments and puts them back in
// place, so long responses stay untouched and fast to fix.
func (g *guardrailsPlugin) rewriteFragments(ctx context.Context, userText, text string, fragments []string, violations []*Violation) (string, error) {
	fixedFragments, err := g.fixFragments(ctx, userText, fragments, violations)
	if err != nil {
		return "", err
	}
	for i, f := range fragments {
		fixed := fixedFragments[i]
		if fixed == "" {
			text = strings.Replace(text, f+"\n", "", 1)
			text = strings.Replace(text, f, "", 1)
			continue
		}
		text = strings.Replace(text, f, fixed, 1)
	}
	text = strings.TrimSpace(text)
	if text == "" {
		return "", errors.New("response is empty after removing fragments")
	}
	return text, nil
}

// fixFragments returns each fragment corrected, in order. An empty fragment
// means it must be removed.
func (g *guardrailsPlugin) fixFragments(ctx context.Context, userText string, fragments []string, violations []*Violation) ([]string, error) {
	var list strings.Builder
	for i, f := range fragments {
		fmt.Fprintf(&list, "<fragmento id=\"%d\">\n%s\n</fragmento>\n", i+1, f)
	}

	temperature := float32(0)
	result, err := g.client.Models.GenerateContent(ctx, rewriteModel,
		genai.Text(fmt.Sprintf(fragmentsPrompt, userText, instructions(violations), list.String())),
		&genai.GenerateContentConfig{
			Temperature:      &temperature,
			ResponseMIMEType: "application/json",
			ResponseSchema:   fragmentsSchema,
		},
	)
	if err != nil {
		return nil, err
	}
	var out struct {
		Keep      bool     `json:"keep"`
		Fragments []string `json:"fragments"`
	}
	if err := json.Unmarshal([]byte(result.Text()), &out); err != nil {
		return nil, fmt.Errorf("invalid rewrite model response: %w", err)
	}
	if out.Keep {
		return nil, errKept
	}
	if len(out.Fragments) != len(fragments) {
		return nil, fmt.Errorf("rewrite model returned %d fragments, want %d", len(out.Fragments), len(fragments))
	}
	for i := range out.Fragments {
		out.Fragments[i] = strings.TrimSpace(out.Fragments[i])
	}
	return out.Fragments, nil
}

// allFragments returns the fragments to rewrite, or nil when any violation
// needs the whole response rewritten.
func allFragments(violations []*Violation) []string {
	var fragments []string
	for _, v := range violations {
		if len(v.Fragments) == 0 {
			return nil
		}
		fragments = append(fragments, v.Fragments...)
	}
	return fragments
}

func instructions(violations []*Violation) string {
	var sb strings.Builder
	for _, v := range violations {
		fmt.Fprintf(&sb, "- %s\n", v.Instruction)
	}
	return sb.String()
}

// responseText returns the visible text of a model response and whether it
// also calls a tool.
func responseText(content *genai.Content) (string, bool, bool) {
	var sb strings.Builder
	hasToolCall := false
	for _, p := range content.Parts {
		if p == nil {
			continue
		}
		if p.FunctionResponse != nil {
			return "", false, false
		}
		if p.FunctionCall != nil {
			hasToolCall = true
		}
		if !p.Thought {
			sb.WriteString(p.Text)
		}
	}
	text := sb.String()
	return text, hasToolCall, strings.TrimSpace(text) != ""
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

// replaceText puts text in the first text part and drops the others. An empty
// text drops them all; their thought signature moves to the tool call so
// Gemini still accepts the history.
func replaceText(resp *model.LLMResponse, text string) *model.LLMResponse {
	parts := make([]*genai.Part, 0, len(resp.Content.Parts))
	var signature []byte
	replaced := false
	for _, p := range resp.Content.Parts {
		if p == nil || p.Thought || p.Text == "" {
			parts = append(parts, p)
			continue
		}
		if signature == nil {
			signature = p.ThoughtSignature
		}
		if !replaced && text != "" {
			parts = append(parts, &genai.Part{Text: text, ThoughtSignature: p.ThoughtSignature})
			replaced = true
		}
	}
	if text == "" && signature != nil {
		for i, p := range parts {
			if p != nil && p.FunctionCall != nil && p.ThoughtSignature == nil {
				moved := *p
				moved.ThoughtSignature = signature
				parts[i] = &moved
				break
			}
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
