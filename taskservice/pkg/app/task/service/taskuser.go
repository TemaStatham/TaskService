package service

import (
	"context"
	"github.com/TemaStatham/TaskService/taskservice/pkg/app/task/model"
)

type TaskUserServiceInterface interface {
	Add(ctx context.Context, userID, taskID uint, isCoordinator bool) error
	Delete(ctx context.Context, userID, taskID uint) error
}

type TaskUserService struct {
	taskuserrepo model.TaskUserRepositoryInterface
}

func NewTaskUserService(taskuserrepo model.TaskUserRepositoryInterface) *TaskUserService {
	return &TaskUserService{
		taskuserrepo: taskuserrepo,
	}
}

func (tu *TaskUserService) Add(ctx context.Context, userID, taskID uint, isCoordinator bool) error {
	return tu.taskuserrepo.Add(ctx, userID, taskID, isCoordinator)
}

func (tu *TaskUserService) Delete(ctx context.Context, userID, taskID uint) error {
	return tu.taskuserrepo.Delete(ctx, userID, taskID)
}
