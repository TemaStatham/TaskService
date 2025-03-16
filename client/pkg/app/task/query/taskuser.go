package query

import (
	"context"
	"errors"
	"github.com/TemaStatham/TaskService/client/pkg/app/paginate"
	"github.com/TemaStatham/TaskService/client/pkg/app/task/model"
)

type TaskUserQueryInterface interface {
	GetUsers(
		ctx context.Context,
		taskID uint,
		pagination *paginate.Pagination,
		isCoordinators *bool,
	) (*paginate.Pagination, error)
}

type TaskUserQuery struct {
	repo model.TaskUserReadRepositoryInterface
}

func NewTaskUserQuery(repo model.TaskUserReadRepositoryInterface) *TaskUserQuery {
	return &TaskUserQuery{
		repo: repo,
	}
}

func (tu *TaskUserQuery) GetUsers(
	ctx context.Context,
	taskID uint,
	pagination *paginate.Pagination,
	isCoordinators *bool,
) (*paginate.Pagination, error) {
	if pagination == nil {
		return nil, errors.New("pagination is required")
	}

	return tu.repo.GetUsers(ctx, taskID, pagination, isCoordinators)
}
