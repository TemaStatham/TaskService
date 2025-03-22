package query

import (
	"context"
	"github.com/TemaStatham/TaskService/taskservice/pkg/app/task/model"
	"github.com/TemaStatham/TaskService/taskservice/pkg/infrastructure/lib/paginate"
)

type TaskUserQueryInterface interface {
	GetUsers(
		ctx context.Context,
		taskID uint,
		page int,
		limit int,
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
	page int,
	limit int,
	isCoordinators *bool,
) (*paginate.Pagination, error) {

	return tu.repo.GetUsers(ctx, taskID, page, limit, isCoordinators)
}
