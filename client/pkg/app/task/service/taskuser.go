package service

import "context"

type TaskUserServiceInterface interface {
	Add(ctx context.Context, userID, taskID uint, isCoordinator bool) error
	Delete(ctx context.Context, userID, taskID uint) error
}

type TaskUserService struct {
	taskservice TaskUserServiceInterface
}

func NewTaskUserService(taskservice TaskUserServiceInterface) *TaskUserService {
	return &TaskUserService{
		taskservice: taskservice,
	}
}

func (tu *TaskUserService) Add(ctx context.Context, userID, taskID uint, isCoordinator bool) error {
	return tu.taskservice.Add(ctx, userID, taskID, isCoordinator)
}

func (tu *TaskUserService) Delete(ctx context.Context, userID, taskID uint) error {
	return tu.taskservice.Delete(ctx, userID, taskID)
}
