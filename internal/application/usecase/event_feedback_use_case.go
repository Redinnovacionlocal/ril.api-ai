package usecase

import (
	"context"
	"time"

	ulid2 "github.com/oklog/ulid"
	"github.com/oklog/ulid/v2"
	"ril.api-ia/internal/domain/entity"
	"ril.api-ia/internal/domain/repository"
)

type EventFeedbackUseCase struct {
	ctx                     context.Context
	eventFeedbackRepository repository.EventFeedbackRepository
}

func NewEventFeedbackUseCase(ctx context.Context, eventFeedbackRepository repository.EventFeedbackRepository) *EventFeedbackUseCase {
	return &EventFeedbackUseCase{
		ctx:                     ctx,
		eventFeedbackRepository: eventFeedbackRepository,
	}
}

func (efuc *EventFeedbackUseCase) SaveFeedback(invocationId string, user *entity.User, isPositive bool, comments *string, errorType *string) error {
	feedback, err := efuc.eventFeedbackRepository.GetFeedbackByInvocationId(invocationId)
	if err != nil {
		return err
	}
	if feedback == nil {
		feedback = &entity.EventFeedback{
			Id:           ulid2.ULID(ulid.Make()),
			InvocationId: invocationId,
			UserId:       user.Id,
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		}
	}
	feedback.IsPositive = isPositive
	feedback.Comments = comments
	feedback.ErrorType = errorType
	feedback.UpdatedAt = time.Now()
	err = efuc.eventFeedbackRepository.SaveFeedback(feedback)
	if err != nil {
		return err
	}
	return nil
}
