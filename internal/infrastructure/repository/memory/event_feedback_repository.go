package memory

import (
	"sync"

	"ril.api-ia/internal/domain/entity"
)

type EventFeedbackRepository struct {
	mux  *sync.RWMutex
	data map[any]*entity.EventFeedback
}

func NewEventFeedbackRepository() *EventFeedbackRepository {
	return &EventFeedbackRepository{
		mux:  new(sync.RWMutex),
		data: make(map[any]*entity.EventFeedback),
	}
}

func (repo *EventFeedbackRepository) SaveFeedback(feedback *entity.EventFeedback) error {
	repo.mux.Lock()
	defer repo.mux.Unlock()
	repo.data[feedback.Id] = feedback
	return nil
}

func (repo *EventFeedbackRepository) GetFeedbackByInvocationId(invocationId string) (*entity.EventFeedback, error) {
	repo.mux.RLock()
	defer repo.mux.RUnlock()
	for _, feedback := range repo.data {
		if feedback.InvocationId == invocationId {
			return feedback, nil
		}
	}
	return nil, nil
}
