package query

import (
	"context"
	"errors"
	"github.com/TemaStatham/TaskService/client/pkg/app/organization/query"
	"github.com/TemaStatham/TaskService/client/pkg/app/paginate"

	"github.com/TemaStatham/TaskService/client/pkg/app/task/model"
)

type TaskQueryInterface interface {
	Get(ctx context.Context, id uint) (*model.TaskModel, error)
	Show(
		ctx context.Context,
		pagination *paginate.Pagination,
		user uint,
	) (*paginate.Pagination, error)
}

type TaskQuery struct {
	readRepository    model.TaskReadRepositoryInterface
	organizationQuery query.OrganizationQueryInterface
}

func NewTaskQuery(
	readRepository model.TaskReadRepositoryInterface,
	organizationQuery query.OrganizationQueryInterface,
) *TaskQuery {
	return &TaskQuery{
		readRepository:    readRepository,
		organizationQuery: organizationQuery,
	}
}

func (t *TaskQuery) Get(ctx context.Context, id uint) (*model.TaskModel, error) {
	task, err := t.readRepository.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return task, nil
}

func (t *TaskQuery) Show(
	ctx context.Context,
	pagination *paginate.Pagination,
	user uint,
) (*paginate.Pagination, error) {
	if pagination == nil {
		return nil, errors.New("paginate is required")
	}

	organizations, err := t.organizationQuery.GetOrganizationsByUserID(ctx, uint64(user))
	if err != nil {
		return nil, err
	}

	taskPagination, err := t.readRepository.GetAll(ctx, pagination, user, organizations)
	if err != nil {
		return nil, err
	}

	return taskPagination, nil
}
