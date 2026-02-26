package session

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/adk/session"
	"gorm.io/gorm"
)

type Service interface {
	session.Service
	UpdateTitle(context.Context, *UpdateTitleRequest) error
}
type UpdateTitleRequest struct {
	AppName   string
	UserID    string
	SessionID string
	Title     string
}

type MyDatabaseService struct {
	session.Service
	db *gorm.DB
}

type storageSession struct {
	AppName    string    `gorm:"primaryKey;"`
	UserID     string    `gorm:"primaryKey;"`
	ID         string    `gorm:"primaryKey;"`
	CreateTime time.Time `gorm:"precision:6"`
	UpdateTime time.Time `gorm:"precision:6"`
}

func (storageSession) TableName() string {
	return "sessions"
}

func NewMyDatabaseService(base session.Service, dialector gorm.Dialector, opts ...gorm.Option) *MyDatabaseService {
	db, err := gorm.Open(dialector, opts...)
	if err != nil {
		panic(fmt.Sprintf("failed to connect database: %v", err))
	}
	return &MyDatabaseService{
		Service: base,
		db:      db,
	}
}

func (s *MyDatabaseService) UpdateTitle(ctx context.Context, req *UpdateTitleRequest) error {
	appName, userID, sessionID := req.AppName, req.UserID, req.SessionID
	if appName == "" || userID == "" || sessionID == "" {
		return fmt.Errorf("app_name, user_id, session_id are required, got app_name: %q, user_id: %q, session_id: %q", appName, userID, sessionID)
	}

	err := s.db.WithContext(ctx).
		Table("sessions").
		Where("app_name = ? AND user_id = ? AND id = ?",
			appName, userID, sessionID,
		).
		Update("state", gorm.Expr(
			"jsonb_set(COALESCE(state, '{}'::jsonb), '{title}', to_jsonb(?::text), true)",
			req.Title,
		)).
		Error

	if err != nil {
		return fmt.Errorf("database error while updating title: %w", err)
	}

	return nil
}
