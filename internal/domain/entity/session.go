package entity

import (
	"maps"

	"google.golang.org/adk/session"
)

type Session struct {
	ID        string         `json:"id"`
	AppName   string         `json:"appName"`
	UserID    string         `json:"userId"`
	UpdatedAt int64          `json:"lastUpdateTime"`
	Events    []Event        `json:"events"`
	State     map[string]any `json:"state"`
}

func FromSession(session session.Session) (Session, error) {
	state := map[string]any{}
	maps.Insert(state, session.State().All())
	events := []Event{}
	for event := range session.Events().All() {
		events = append(events, FromSessionEvent(*event))
	}
	mappedSession := Session{
		ID:        session.ID(),
		AppName:   session.AppName(),
		UserID:    session.UserID(),
		UpdatedAt: session.LastUpdateTime().Unix(),
		Events:    events,
		State:     state,
	}
	return mappedSession, nil
}
