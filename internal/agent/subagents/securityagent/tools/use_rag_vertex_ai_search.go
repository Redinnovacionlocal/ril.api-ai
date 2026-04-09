package tools

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os/exec"
	"strings"

	"google.golang.org/adk/tool"
)

type UseRagVertexAISearchToolArgs struct {
	Query  string            `json:"query" jsonSchema:"The search query to perform using Vertex AI."`
	Filter map[string]string `json:"filter,omitempty" jsonSchema:"Optional filters to apply to the search results."`
}

// getGCloudToken obtiene el access token del CLI de gcloud en runtime.
func getGCloudToken() (string, error) {
	out, err := exec.Command("gcloud", "auth", "print-access-token").Output()
	if err != nil {
		return "", fmt.Errorf("failed to get gcloud token: %w", err)
	}
	return strings.TrimSpace(string(out)), nil
}

// buildFilter convierte el mapa de filtros al formato de la Discovery Engine API.
// Ejemplo: {"category": "security"} → "category: security"
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
	const (
		apiEndpoint = "https://discoveryengine.googleapis.com/v1/"
		datastore   = "projects/ril-admin/locations/global/collections/default_collection/dataStores/ril-security-knowledge_1775562649372_gcs_store/servingConfigs/default_search:search"
	)

	// 1. Obtener token de gcloud
	token, err := getGCloudToken()
	if err != nil {
		return nil, err
	}

	// 2. Construir payload según la Discovery Engine API
	payload := map[string]any{
		"query":    args.Query,
		"pageSize": 10,
	}
	if filterStr := buildFilter(args.Filter); filterStr != "" {
		payload["filter"] = filterStr
	}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("failed to serialize payload: %w", err)
	}

	// 3. Construir request con Authorization header
	req, err := http.NewRequest(http.MethodPost, apiEndpoint+datastore, bytes.NewBuffer(jsonPayload))
	if err != nil {
		return nil, fmt.Errorf("failed to build request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	// 4. Ejecutar request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to perform search: %w", err)
	}
	defer resp.Body.Close()

	// 5. Leer body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("search request failed with status %s: %s", resp.Status, string(body))
	}

	// 6. Parsear respuesta JSON
	var result map[string]any
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	return result, nil
}
