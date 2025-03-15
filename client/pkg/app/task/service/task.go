package service

import (
	"context"
	"github.com/TemaStatham/TaskService/client/pkg/app/organization/query"
	"github.com/TemaStatham/TaskService/client/pkg/app/task/data"
	"github.com/TemaStatham/TaskService/client/pkg/app/task/model"
)

type TaskServiceInterface interface {
	Update(ctx context.Context, task *data.UpdateTask) error
	Delete(ctx context.Context, id uint) error
	Create(ctx context.Context, task *data.CreateTask) (uint, error)
}

type TaskService struct {
	taskRepository    model.TaskRepositoryInterface
	organizationQuery query.OrganizationQueryInterface
}

func NewTaskService(
	rep model.TaskRepositoryInterface,
	organizationQuery query.OrganizationQueryInterface,
) *TaskService {
	return &TaskService{
		taskRepository:    rep,
		organizationQuery: organizationQuery,
	}
}

func (t *TaskService) Update(ctx context.Context, task *data.UpdateTask) error {
	err := t.taskRepository.Update(ctx, task)

	return err
}

func (t *TaskService) Delete(ctx context.Context, id uint) error {
	err := t.taskRepository.Delete(ctx, id)
	return err
}

func (t *TaskService) Create(ctx context.Context, task *data.CreateTask) (uint, error) {
	_, err := t.organizationQuery.GetOrganization(ctx, uint64(task.Organization))
	if err != nil {
		return 0, err
	}

	return t.taskRepository.Create(ctx, task)
}
