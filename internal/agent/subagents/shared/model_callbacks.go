package shared

import (
	"reflect"

	"google.golang.org/adk/agent"
	"google.golang.org/adk/model"
	"google.golang.org/genai"
)

func StripEmptyThoughtSignatureParts(_ agent.CallbackContext, req *model.LLMRequest) (*model.LLMResponse, error) {
	contents := make([]*genai.Content, 0, len(req.Contents))
	for _, content := range req.Contents {
		if content == nil {
			continue
		}
		parts := make([]*genai.Part, 0, len(content.Parts))
		for _, part := range content.Parts {
			if isPartWithoutUsableContent(part) {
				continue
			}
			parts = append(parts, part)
		}
		if len(parts) == 0 {
			continue
		}
		newContent := *content
		newContent.Parts = parts
		contents = append(contents, &newContent)
	}
	req.Contents = contents
	return nil, nil
}

func isPartWithoutUsableContent(p *genai.Part) bool {
	if p == nil {
		return true
	}
	stripped := *p
	stripped.ThoughtSignature = nil
	return reflect.ValueOf(stripped).IsZero()
}
