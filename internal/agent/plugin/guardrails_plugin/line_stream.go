package guardrails_plugin

import (
	"context"
	"errors"
	"log/slog"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/adk/agent"
	"google.golang.org/adk/model"
	"google.golang.org/genai"
	"ril.api-ia/internal/infrastructure/observability"
)

// lineStream holds back a streamed response so that only checked text reaches
// the user. Text is released line by line once every line rule passes; failed
// lines wait in pending and are corrected together, before the next good line
// or at the end of the response.
type lineStream struct {
	in          Input
	checkers    []func(line string) *Violation
	buf         string
	pending     []string
	violations  []*Violation
	released    strings.Builder
	inCode      bool
	hasToolCall bool
	kept        bool
}

func streamKey(ctx agent.CallbackContext) string {
	return ctx.InvocationID() + "/" + ctx.AgentName()
}

func (g *guardrailsPlugin) newLineStream(ctx agent.CallbackContext) *lineStream {
	s := &lineStream{in: Input{UserText: contentText(ctx.UserContent()), State: ctx.State()}}
	for _, r := range g.rules {
		if lr, ok := r.(LineRule); ok {
			if check := lr.LineChecker(s.in); check != nil {
				s.checkers = append(s.checkers, check)
			}
		}
	}
	return s
}

// afterPartial returns the streamed chunk with only the text that passed the
// line rules so far.
func (g *guardrailsPlugin) afterPartial(ctx agent.CallbackContext, key string, resp *model.LLMResponse) *model.LLMResponse {
	v, ok := g.streams.Load(key)
	if !ok {
		v, _ = g.streams.LoadOrStore(key, g.newLineStream(ctx))
	}
	s := v.(*lineStream)

	text, hasToolCall := chunkText(resp.Content)
	s.hasToolCall = s.hasToolCall || hasToolCall
	out := g.feed(ctx, s, text, resp.TurnComplete || resp.FinishReason != "")
	if out == text {
		return nil
	}
	return withText(resp, out)
}

// afterStream sets the final response to exactly the text the user saw, so
// the session history matches it.
func (g *guardrailsPlugin) afterStream(ctx agent.CallbackContext, s *lineStream, resp *model.LLMResponse) *model.LLMResponse {
	// The stream can end without a finish reason, leaving text held back.
	if s.buf != "" || len(s.pending) > 0 {
		g.feed(ctx, s, "", true)
	}
	text := s.released.String()

	// Rules also update session state, so they run on the final text anyway.
	if strings.TrimSpace(text) != "" {
		for _, v := range g.check(Input{UserText: s.in.UserText, ResponseText: text, State: ctx.State()}) {
			if s.kept && v.OnKeep != nil {
				v.OnKeep()
				continue
			}
			slog.WarnContext(ctx, "guardrails: rule failed after streaming",
				append([]any{"guardrails.rule", v.Rule, "guardrails.detail", v.Detail}, logAttrs(ctx)...)...)
		}
	}

	if resp.Content == nil {
		return nil
	}
	if original, _, _ := responseText(resp.Content); original == text {
		return nil
	}
	return replaceText(resp, text)
}

// feed adds streamed text and returns the text ready to release. last flushes
// everything held back.
func (g *guardrailsPlugin) feed(ctx agent.CallbackContext, s *lineStream, text string, last bool) string {
	if len(s.checkers) == 0 {
		s.released.WriteString(text)
		return text
	}
	s.buf += text
	var out strings.Builder
	for {
		i := strings.IndexByte(s.buf, '\n')
		if i < 0 {
			break
		}
		out.WriteString(g.feedLine(ctx, s, s.buf[:i+1]))
		s.buf = s.buf[i+1:]
	}
	if last {
		// The model sometimes stops mid-sentence to call a tool ("...el municipio.\n\nCon ").
		if s.buf != "" && s.hasToolCall && !endsSentence(s.buf) {
			slog.InfoContext(ctx, "guardrails: unfinished sentence removed before tool call", logAttrs(ctx)...)
			s.buf = ""
		}
		if s.buf != "" {
			out.WriteString(g.feedLine(ctx, s, s.buf))
			s.buf = ""
		}
		out.WriteString(g.resolve(ctx, s))
	}
	s.released.WriteString(out.String())
	return out.String()
}

func (g *guardrailsPlugin) feedLine(ctx agent.CallbackContext, s *lineStream, line string) string {
	if strings.HasPrefix(strings.TrimSpace(line), "```") {
		s.inCode = !s.inCode
	} else if !s.inCode {
		for _, check := range s.checkers {
			if v := check(strings.TrimSuffix(line, "\n")); v != nil {
				s.pending = append(s.pending, line)
				s.violations = append(s.violations, v)
				return ""
			}
		}
	}
	return g.resolve(ctx, s) + line
}

// resolve corrects the pending lines in a single call and returns them.
func (g *guardrailsPlugin) resolve(ctx agent.CallbackContext, s *lineStream) string {
	if len(s.pending) == 0 {
		return ""
	}
	pending, violations := s.pending, s.violations
	s.pending, s.violations = nil, nil

	attrs := append([]any{"guardrails.rules", ruleNames(violations), "guardrails.lines", len(pending)}, logAttrs(ctx)...)
	for _, v := range violations {
		slog.WarnContext(ctx, "guardrails: rule failed",
			append([]any{"guardrails.rule", v.Rule, "guardrails.detail", v.Detail, "guardrails.tool_call", s.hasToolCall}, logAttrs(ctx)...)...)
	}

	// Text next to a tool call is filler before the tool runs: drop it instead
	// of delaying the tool with a rewrite.
	if s.hasToolCall {
		slog.InfoContext(ctx, "guardrails: text removed from tool call", attrs...)
		return ""
	}

	fragments := make([]string, len(pending))
	for i, line := range pending {
		fragments[i] = strings.TrimSuffix(line, "\n")
	}
	start := time.Now()
	fixed, err := g.fixLines(ctx, s.in.UserText, fragments, violations)
	attrs = append(attrs, "guardrails.duration_ms", time.Since(start).Milliseconds())
	if errors.Is(err, errKept) {
		s.kept = true
		slog.InfoContext(ctx, "guardrails: response kept at user request", attrs...)
		return strings.Join(pending, "")
	}
	if err != nil {
		slog.WarnContext(ctx, "guardrails: fallback to original response",
			append([]any{"guardrails.reason", err.Error()}, attrs...)...)
		return strings.Join(pending, "")
	}

	var out strings.Builder
	for i, f := range fixed {
		if f == "" {
			continue
		}
		out.WriteString(f)
		if strings.HasSuffix(pending[i], "\n") {
			out.WriteString("\n")
		}
	}
	slog.InfoContext(ctx, "guardrails: response corrected", attrs...)
	return out.String()
}

func (g *guardrailsPlugin) fixLines(ctx context.Context, userText string, fragments []string, violations []*Violation) ([]string, error) {
	ctx, cancel := context.WithTimeout(ctx, correctTimeout)
	defer cancel()
	ctx, span := observability.Start(ctx, "guardrails.correct", observability.Attrs{
		"guardrails.rules":     ruleNames(violations),
		"gen_ai.request.model": rewriteModel,
	})
	defer span.End()

	fixed, err := g.fixFragments(ctx, userText, fragments, violations)
	if err != nil && !errors.Is(err, errKept) {
		span.Fail(err)
	}
	return fixed, err
}

func endsSentence(text string) bool {
	last, _ := utf8.DecodeLastRuneInString(strings.TrimSpace(text))
	return last == utf8.RuneError || strings.ContainsRune(".:!?)»\"*…", last)
}

func logAttrs(ctx agent.CallbackContext) []any {
	return []any{
		"agent", ctx.AgentName(),
		"invocation_id", ctx.InvocationID(),
		"session_id", ctx.SessionID(),
	}
}

// chunkText returns the visible text of a streamed chunk and whether it calls
// a tool.
func chunkText(content *genai.Content) (string, bool) {
	if content == nil {
		return "", false
	}
	var sb strings.Builder
	hasToolCall := false
	for _, p := range content.Parts {
		if p == nil {
			continue
		}
		if p.FunctionCall != nil {
			hasToolCall = true
		}
		if !p.Thought {
			sb.WriteString(p.Text)
		}
	}
	return sb.String(), hasToolCall
}

// withText replaces the visible text of a streamed chunk. A chunk left empty
// has nil content, so ADK does not send it.
func withText(resp *model.LLMResponse, text string) *model.LLMResponse {
	var parts []*genai.Part
	inserted := text == ""
	if resp.Content != nil {
		for _, p := range resp.Content.Parts {
			if p != nil && !p.Thought && p.Text != "" {
				if !inserted {
					parts = append(parts, &genai.Part{Text: text})
					inserted = true
				}
				continue
			}
			parts = append(parts, p)
		}
	}
	if !inserted {
		parts = append([]*genai.Part{{Text: text}}, parts...)
	}

	newResp := *resp
	newResp.Content = nil
	if len(parts) > 0 {
		newResp.Content = &genai.Content{Role: genai.RoleModel, Parts: parts}
	}
	return &newResp
}
