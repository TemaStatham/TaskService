package query

import (
	"context"
	"errors"
	"github.com/TemaStatham/TaskService/taskservice/pkg/app/organization/query"
	"github.com/TemaStatham/TaskService/taskservice/pkg/app/task/data"
	"github.com/TemaStatham/TaskService/taskservice/pkg/app/task/model"
	userquery "github.com/TemaStatham/TaskService/taskservice/pkg/app/user/query"
	"github.com/TemaStatham/TaskService/taskservice/pkg/infrastructure/lib/paginate"
)

type TaskQueryInterface interface {
	Get(ctx context.Context, id uint) (*model.TaskModel, error)
	Show(
		ctx context.Context,
		dto data.GetAllTasks,
		authUser uint,
	) (*paginate.Pagination, error)
}

type TaskQuery struct {
	readRepository    model.TaskReadRepositoryInterface
	organizationQuery query.OrganizationQueryInterface
	userQuery         userquery.UserQueryInterface
}

func NewTaskQuery(
	readRepository model.TaskReadRepositoryInterface,
	organizationQuery query.OrganizationQueryInterface,
	userQuery userquery.UserQueryInterface,
) *TaskQuery {
	return &TaskQuery{
		readRepository:    readRepository,
		organizationQuery: organizationQuery,
		userQuery:         userQuery,
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
	dto data.GetAllTasks,
	authUser uint,
) (*paginate.Pagination, error) {
	user, err := t.userQuery.GetUser(ctx, uint64(authUser))
	if err != nil {
		return nil, err
	}

	// Получаем все организации, обращаясь к сервису профилей
	organizations, err := t.organizationQuery.GetOrganizationsByUserID(ctx, uint64(authUser))
	if err != nil {
		return nil, err
	}

	if user.IsAdmin {
		if len(organizations) == 0 {
			return &paginate.Pagination{}, errors.New("user is admin but organization not found")
		}
		return t.readRepository.GetByOrganization(ctx, organizations[0].ID)
	}

	var orgIDs []uint
	for _, org := range organizations {
		orgIDs = append(orgIDs, org.ID)
	}

	taskPagination, err := t.readRepository.GetAll(ctx, dto, authUser, organizations)
	if err != nil {
		return nil, err
	}

	return taskPagination, nil
}
