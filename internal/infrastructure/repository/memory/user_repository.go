package memory

import (
	"sync"

	"ril.api-ia/internal/domain/entity"
)

type UserRepository struct {
	mux  *sync.RWMutex
	data map[int64]*entity.User
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		mux:  new(sync.RWMutex),
		data: make(map[int64]*entity.User),
	}
}

func (repo *UserRepository) Save(user *entity.User) error {
	repo.mux.Lock()
	defer repo.mux.Unlock()
	repo.data[user.Id] = user
	return nil
}

func (repo *UserRepository) FindByAiApiKey(aiApiKey string) (*entity.User, error) {
	for _, user := range repo.data {
		if *user.ApiAiToken == aiApiKey {
			return user, nil
		}
	}
	return nil, nil
}
func (repo *UserRepository) GetUserProfile(user *entity.User) (*entity.UserProfile, error) {
	return &user.UserProfile, nil
}
