package repository

import "ril.api-ia/internal/domain/entity"

type UserRepository interface {
	FindByAiApiKey(aiApiKey string) (*entity.User, error)
	GetUserProfile(user *entity.User) (*entity.UserProfile, error)
}
