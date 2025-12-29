package usecase

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.com/redis/go-redis/v9"
	"ril.api-ia/internal/domain/entity"
	"ril.api-ia/internal/domain/repository"
)

type UserUseCase struct {
	userRepository repository.UserRepository
	rdb            *redis.Client
	ctx            context.Context
}

func NewUserUseCase(ctx context.Context, userRepository repository.UserRepository, rdb *redis.Client) *UserUseCase {
	return &UserUseCase{
		userRepository: userRepository,
		rdb:            rdb,
		ctx:            ctx,
	}
}

func (uuc *UserUseCase) GetUserByApiAiToken(apiAiToken string) (*entity.User, error) {
	var user *entity.User
	val, err := uuc.rdb.Get(uuc.ctx, "user:api_ai_token:"+apiAiToken).Result()
	if errors.Is(err, redis.Nil) {
		user, err = uuc.userRepository.FindByAiApiKey(apiAiToken)
		if err != nil {
			return nil, err
		}
		if user != nil {
			userJson, err := json.Marshal(user)
			if err != nil {
				return nil, err
			}
			err = uuc.rdb.Set(uuc.ctx, "user:api_ai_token:"+apiAiToken, userJson, time.Hour).Err()
			if err != nil {
				return nil, err
			}
			return user, nil
		}
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	if err := json.Unmarshal([]byte(val), &user); err != nil {
		return nil, err
	}
	return user, nil
}
