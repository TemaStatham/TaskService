package service

import (
	"context"
	"github.com/TemaStatham/TaskService/internal/model"
)

type UserRepository interface {
	Get(ctx context.Context, id uint) (*model.User, error)
	Create(ctx context.Context, user *model.User) error
}

type UserService struct {
	UserRepository
}

func NewUserService(repository UserRepository) UserService {
	return UserService{
		repository,
	}
}

func (u *UserService) CreateFromKafka(key, value []byte) error {

	return nil
}
