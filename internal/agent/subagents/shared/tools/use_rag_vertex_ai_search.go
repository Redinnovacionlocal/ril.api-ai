package tools

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"golang.org/x/oauth2/google"
	"google.golang.org/adk/tool"
	"google.golang.org/adk/tool/functiontool"
)

type UseRagVertexAISearchToolArgs struct {
	Query  string            `json:"query" jsonSchema:"The search query to perform using Vertex AI."`
	Filter map[string]string `json:"filter,omitempty" jsonSchema:"Optional filters to apply to the search results."`
}

func useRagVertexAISearchWithDatastore(ctx tool.Context, args UseRagVertexAISearchToolArgs, datastore string) (map[string]any, error) {
	apiEndpoint := fmt.Sprintf(
		"https://discoveryengine.googleapis.com/v1/projects/ril-admin/locations/global/collections/default_collection/dataStores/%s/servingConfigs/default_search:search",
		datastore,
	)

	client, err := google.DefaultClient(ctx, "https://www.googleapis.com/auth/cloud-platform")
	if err != nil {
		return nil, fmt.Errorf("failed to create authenticated client: %w", err)
	}

	payload := map[string]any{
		"query":    args.Query,
		"pageSize": 25,
		"contentSearchSpec": map[string]any{
			"snippetSpec": map[string]any{
				"returnSnippet": true,
			},
			"summarySpec": map[string]any{
				"summaryResultCount": 3,
			},
		},
	}

	if filterStr := buildFilter(args.Filter); filterStr != "" {
		payload["filter"] = filterStr
	}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, apiEndpoint, bytes.NewBuffer(jsonPayload))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		bodyStr := string(body)

		if resp.StatusCode == http.StatusBadRequest && strings.Contains(bodyStr, "Unsupported field") {
			fmt.Printf("Aviso RAG: Filtro ignorado (campo no existe en Schema). Devolviendo 0 resultados. Detalle: %s\n", bodyStr)

			return map[string]any{
				"results": []any{},
			}, nil
		}

		return nil, fmt.Errorf("error %d: %s", resp.StatusCode, bodyStr)
	}

	var result map[string]any
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	return result, nil
}

func NewRagTool(datastoreID string) (tool.Tool, error) {
	if datastoreID == "" {
		return nil, fmt.Errorf("datastoreID is required")
	}
	return functiontool.New(functiontool.Config{
		Name:        "rilia_rag_agent",
		Description: "Permite al agente utilizar un documento recuperado del RAG como parte de su respuesta al usuario. El agente puede extraer información relevante del documento para enriquecer sus recomendaciones.",
	}, func(ctx tool.Context, args UseRagVertexAISearchToolArgs) (map[string]any, error) {
		return useRagVertexAISearchWithDatastore(ctx, args, datastoreID)
	})
}

func buildFilter(filter map[string]string) string {
	if len(filter) == 0 {
		return ""
	}
	parts := make([]string, 0, len(filter))
	for k, v := range filter {
		parts = append(parts, fmt.Sprintf("%s: ANY(\"%s\")", k, v))
	}
	return strings.Join(parts, " AND ")
}
