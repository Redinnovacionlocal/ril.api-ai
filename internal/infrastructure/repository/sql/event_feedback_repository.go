package sql

import (
	"strconv"

	"github.com/jmoiron/sqlx"
	"ril.api-ia/internal/domain/entity"
)

type EventFeedbackRepository struct {
	db *sqlx.DB
}

func NewEventFeedbackRepository(db *sqlx.DB) *EventFeedbackRepository {
	return &EventFeedbackRepository{
		db: db,
	}
}

func (repo *EventFeedbackRepository) SaveFeedback(e *entity.EventFeedback) error {
	query := `
        INSERT INTO feedback_events (
            id, user_id, is_positive, comments, created_at, updated_at, error_type, invocation_id
        ) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
        ON CONFLICT (id) DO UPDATE SET
            user_id = EXCLUDED.user_id,
            is_positive = EXCLUDED.is_positive,
            comments = EXCLUDED.comments,
            updated_at = NOW(), 
            error_type = EXCLUDED.error_type,
            invocation_id = EXCLUDED.invocation_id;
    `
	userIdStr := strconv.FormatInt(e.UserId, 10)
	_, err := repo.db.Exec(query,
		e.Id.String(),
		userIdStr,
		e.IsPositive,
		e.Comments,
		e.CreatedAt,
		e.UpdatedAt,
		e.ErrorType,
		e.InvocationId,
	)
	return err
}

func (repo *EventFeedbackRepository) GetFeedbackByInvocationId(invocationId string) (*entity.EventFeedback, error) {
	query := `
		SELECT id, user_id, is_positive, comments, created_at, updated_at, error_type, invocation_id
		FROM feedback_events
		WHERE invocation_id = $1;
	`
	var feedback entity.EventFeedback
	err := repo.db.Get(&feedback, query, invocationId)
	if err != nil {
		return nil, nil
	}
	return &feedback, nil
}
