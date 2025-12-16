package sql

import (
	"github.com/jmoiron/sqlx"
	"ril.api-ia/internal/domain/entity"
)

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) FindByAiApiKey(aiApiKey string) (*entity.User, error) {
	var u entity.User
	err := r.db.Get(&u, "select * from user where api_ai_token = $1", aiApiKey)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func (r *UserRepository) GetUserProfile(user *entity.User) (*entity.UserProfile, error) {
	return &user.UserProfile, nil
}
