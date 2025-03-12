package service

import (
	"context"
	"github.com/TemaStatham/TaskService/internal/model"
)

type UserRepository interface {
	Get(ctx context.Context, id uint) (*model.UserModel, error)
	IsCoordinatorForTask(ctx context.Context, userID, taskID uint) (bool, error)
	IsAdmin(ctx context.Context, userID uint) (bool, error)
	IsOwner(ctx context.Context, userID uint) (bool, error)
	Create(ctx context.Context, user *model.UserModel) error
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
