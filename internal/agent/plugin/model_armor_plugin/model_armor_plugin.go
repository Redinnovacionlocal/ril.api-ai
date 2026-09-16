// Package model_armor_plugin sanitizes the user's first message of each turn
// through Google Cloud Model Armor before it reaches the model, to catch
// prompt injection/jailbreak attempts and other disallowed content.
package model_armor_plugin

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	modelarmor "cloud.google.com/go/modelarmor/apiv1"
	"cloud.google.com/go/modelarmor/apiv1/modelarmorpb"
	"google.golang.org/adk/agent"
	"google.golang.org/adk/plugin"
	"google.golang.org/api/option"
	"google.golang.org/genai"
	"ril.api-ia/internal/infrastructure/env"
	"ril.api-ia/internal/infrastructure/observability"
)

type modelArmorPlugin struct {
	name         string
	client       *modelarmor.Client
	templateName string
}

// New builds the Model Armor plugin. Configured via env vars:
//
//	MODEL_ARMOR_PROJECT   (default: GOOGLE_CLOUD_PROJECT)
//	MODEL_ARMOR_LOCATION  (default: us-central1)
//	MODEL_ARMOR_TEMPLATE
func New(ctx context.Context, name string) (*plugin.Plugin, error) {
	if name == "" {
		name = "model_armor_plugin"
	}

	project := env.GetOrDefault("MODEL_ARMOR_PROJECT", os.Getenv("GOOGLE_CLOUD_PROJECT"))
	location := env.GetOrDefault("MODEL_ARMOR_LOCATION", "us-central1")
	templateID := env.GetOrDefault("MODEL_ARMOR_TEMPLATE", "prueba")

	endpoint := fmt.Sprintf("modelarmor.%s.rep.googleapis.com:443", location)
	client, err := modelarmor.NewClient(ctx, option.WithEndpoint(endpoint))
	if err != nil {
		return nil, fmt.Errorf("error creando cliente de Model Armor: %w", err)
	}

	a := &modelArmorPlugin{
		name:         name,
		client:       client,
		templateName: fmt.Sprintf("projects/%s/locations/%s/templates/%s", project, location, templateID),
	}

	return plugin.New(plugin.Config{
		Name:                  a.name,
		OnUserMessageCallback: a.onUserMessage,
		CloseFunc:             client.Close,
	})
}

func (a *modelArmorPlugin) onUserMessage(ctx agent.InvocationContext, userMessage *genai.Content) (*genai.Content, error) {
	text := extractText(userMessage)
	if text == "" {
		return nil, nil
	}

	spanCtx, span := observability.Start(ctx, "model_armor.sanitize_user_prompt", observability.Attrs{
		"model_armor.template":   a.templateName,
		"model_armor.prompt_len": len(text),
	})
	defer span.End()

	resp, err := a.client.SanitizeUserPrompt(spanCtx, &modelarmorpb.SanitizeUserPromptRequest{
		Name: a.templateName,
		UserPromptData: &modelarmorpb.DataItem{
			DataItem: &modelarmorpb.DataItem_Text{Text: text},
		},
	})
	if err != nil {
		span.Fail(err)
		log.Printf("model_armor_plugin: error llamando a Model Armor: %v", err)
		return nil, nil
	}

	result := resp.GetSanitizationResult()
	span.Set(observability.Attrs{
		"model_armor.filter_match_state": result.GetFilterMatchState().String(),
		"model_armor.invocation_result":  result.GetInvocationResult().String(),
	})
	log.Printf("model_armor_plugin: filter_match_state=%s invocation_result=%s",
		result.GetFilterMatchState(), result.GetInvocationResult())

	if result.GetFilterMatchState() != modelarmorpb.FilterMatchState_MATCH_FOUND {
		return nil, nil
	}

	if pij := result.GetFilterResults()["pi_and_jailbreak"].GetPiAndJailbreakFilterResult(); pij.GetMatchState() == modelarmorpb.FilterMatchState_MATCH_FOUND {
		return nil, fmt.Errorf("mensaje bloqueado por Model Armor: posible prompt injection/jailbreak (confianza: %s)", pij.GetConfidenceLevel())
	}

	return nil, fmt.Errorf("mensaje bloqueado por Model Armor: contenido no permitido detectado")
}

func extractText(content *genai.Content) string {
	if content == nil {
		return ""
	}
	var sb strings.Builder
	for _, part := range content.Parts {
		sb.WriteString(part.Text)
	}
	return sb.String()
}
