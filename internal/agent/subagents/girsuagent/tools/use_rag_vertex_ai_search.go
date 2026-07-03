package tools

import (
	"google.golang.org/adk/tool"
	tools2 "ril.api-ia/internal/agent/subagents/securityagent/tools"
)

func UseRagVertexAISearchToolFunc(ctx tool.Context, args tools2.UseRagVertexAISearchToolArgs) (map[string]any, error) {
	return tools2.UseRagVertexAISearchWithDatastore(
		ctx,
		args,
		"rilia-girsu-rag-agent_1783019204455_gcs_store",
	)
}
