package repository

import (
	"ril.api-ia/internal/domain/entity"
)

type EventFeedbackRepository interface {
	SaveFeedback(feedback *entity.EventFeedback) error
	GetFeedbackByInvocationId(invocationId string) (*entity.EventFeedback, error)
}
