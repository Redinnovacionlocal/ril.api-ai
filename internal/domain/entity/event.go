package entity

import (
	"google.golang.org/adk/session"
	"google.golang.org/genai"
)

type EventActions struct {
	StateDelta    map[string]any   `json:"stateDelta"`
	ArtifactDelta map[string]int64 `json:"artifactDelta"`
}

// Event represents a single event in a session.
type Event struct {
	ID                 string                   `json:"id"`
	Time               int64                    `json:"time"`
	InvocationID       string                   `json:"invocationId"`
	Branch             string                   `json:"branch"`
	Author             string                   `json:"author"`
	Partial            bool                     `json:"partial"`
	LongRunningToolIDs []string                 `json:"longRunningToolIds"`
	Content            *genai.Content           `json:"content"`
	GroundingMetadata  *genai.GroundingMetadata `json:"groundingMetadata"`
	TurnComplete       bool                     `json:"turnComplete"`
	Interrupted        bool                     `json:"interrupted"`
	ErrorCode          string                   `json:"errorCode"`
	ErrorMessage       string                   `json:"errorMessage"`
	Actions            EventActions             `json:"actions"`
}

func FromSessionEvent(event session.Event) Event {
	return Event{
		ID:                 event.ID,
		Time:               event.Timestamp.Unix(),
		InvocationID:       event.InvocationID,
		Branch:             event.Branch,
		Author:             event.Author,
		Partial:            event.Partial,
		LongRunningToolIDs: event.LongRunningToolIDs,
		Content:            filterThoughts(event.LLMResponse.Content),
		GroundingMetadata:  event.LLMResponse.GroundingMetadata,
		TurnComplete:       event.LLMResponse.TurnComplete,
		Interrupted:        event.LLMResponse.Interrupted,
		ErrorCode:          event.LLMResponse.ErrorCode,
		ErrorMessage:       event.LLMResponse.ErrorMessage,
		Actions: EventActions{
			StateDelta:    event.Actions.StateDelta,
			ArtifactDelta: event.Actions.ArtifactDelta,
		},
	}
}

func filterThoughts(content *genai.Content) *genai.Content {
	if content == nil {
		return nil
	}
	filtered := make([]*genai.Part, 0, len(content.Parts))
	for _, p := range content.Parts {
		if p.Thought || len(p.ThoughtSignature) > 0 {
			continue
		}
		filtered = append(filtered, p)
	}
	return &genai.Content{
		Role:  content.Role,
		Parts: filtered,
	}
}
