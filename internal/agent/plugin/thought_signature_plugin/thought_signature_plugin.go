package thought_signature_plugin

import (
	"google.golang.org/adk/plugin"
	"ril.api-ia/internal/agent/subagents/shared"
)

func New(name string) (*plugin.Plugin, error) {
	if name == "" {
		name = "thought_signature_plugin"
	}

	return plugin.New(plugin.Config{
		Name:                name,
		BeforeModelCallback: shared.StripEmptyThoughtSignatureParts,
	})
}
