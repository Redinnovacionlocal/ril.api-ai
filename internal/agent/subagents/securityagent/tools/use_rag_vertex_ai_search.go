package tools

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"golang.org/x/oauth2/google"
	"google.golang.org/adk/tool"
)

type UseRagVertexAISearchToolArgs struct {
	Query  string            `json:"query" jsonSchema:"The search query to perform using Vertex AI."`
	Filter map[string]string `json:"filter,omitempty" jsonSchema:"Optional filters to apply to the search results."`
}

func getADCToken(ctx context.Context) (string, error) {
	tokenSource, err := google.DefaultTokenSource(ctx,
		"https://www.googleapis.com/auth/cloud-platform",
	)
	if err != nil {
		return "", fmt.Errorf("failed to get token source: %w", err)
	}

	token, err := tokenSource.Token()
	if err != nil {
		return "", fmt.Errorf("failed to get token: %w", err)
	}

	return token.AccessToken, nil
}

func buildFilter(filter map[string]string) string {
	if len(filter) == 0 {
		return ""
	}
	parts := make([]string, 0, len(filter))
	for k, v := range filter {
		parts = append(parts, fmt.Sprintf("%s: \"%s\"", k, v))
	}
	return strings.Join(parts, " AND ")
}

func UseRagVertexAISearchToolFunc(ctx tool.Context, args UseRagVertexAISearchToolArgs) (map[string]any, error) {
	const apiEndpoint = "https://discoveryengine.googleapis.com/v1/projects/ril-admin/locations/global/collections/default_collection/dataStores/ril-security-knowledge_1775562649372_gcs_store/servingConfigs/default_search:search"

	// 1. Obtener cliente autenticado (Maneja el cacheo de tokens por ti)
	// Es mucho más eficiente que obtener el token manualmente cada vez
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
		return nil, fmt.Errorf("error %d: %s", resp.StatusCode, string(body))
	}

	var result map[string]any
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	return result, nil
}
