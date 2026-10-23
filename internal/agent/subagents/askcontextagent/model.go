package askcontextagent

import (
	"context"
	"iter"
	"log/slog"
	"regexp"

	"google.golang.org/adk/model"
)

// brokenTextRe matches control characters, raw or as JSON escapes. The model
// sometimes returns them in place of accents ("ni\x11os" for "niños"), which
// breaks the form.
var brokenTextRe = regexp.MustCompile(`[\x00-\x08\x0b\x0c\x0e-\x1f]|\\u00(0[0-8bcef]|1[0-9a-fA-F])`)

// formModel always calls the model without streaming, since agenttool waits
// for the full answer anyway, and retries once when the text comes broken.
type formModel struct {
	model.LLM
}

func (m formModel) GenerateContent(ctx context.Context, req *model.LLMRequest, _ bool) iter.Seq2[*model.LLMResponse, error] {
	return func(yield func(*model.LLMResponse, error) bool) {
		var resp *model.LLMResponse
		var err error
		for attempt := 1; attempt <= 2; attempt++ {
			resp, err = generate(ctx, m.LLM, req)
			if err != nil || !isBroken(resp) {
				break
			}
			slog.WarnContext(ctx, "ask_context: broken characters in model output, retrying", "attempt", attempt)
		}
		yield(resp, err)
	}
}

func generate(ctx context.Context, llm model.LLM, req *model.LLMRequest) (*model.LLMResponse, error) {
	var last *model.LLMResponse
	for resp, err := range llm.GenerateContent(ctx, req, false) {
		if err != nil {
			return nil, err
		}
		last = resp
	}
	return last, nil
}

func isBroken(resp *model.LLMResponse) bool {
	if resp == nil || resp.Content == nil {
		return false
	}
	for _, p := range resp.Content.Parts {
		if p != nil && brokenTextRe.MatchString(p.Text) {
			return true
		}
	}
	return false
}
