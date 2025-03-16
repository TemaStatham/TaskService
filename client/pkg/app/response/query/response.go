package query

import (
	"context"
	"errors"
	"github.com/TemaStatham/TaskService/client/pkg/app/paginate"
	"github.com/TemaStatham/TaskService/client/pkg/app/response/model"
)

type ResponseQueryInterface interface {
	Show(
		ctx context.Context,
		taskId uint,
		pagination *paginate.Pagination,
	) (*paginate.Pagination, error)
}

type ResponseQuery struct {
	responseRepository model.ResponseRepositoryReadInterface
}

func NewResponseQuery(responseRepository model.ResponseRepositoryReadInterface) *ResponseQuery {
	return &ResponseQuery{
		responseRepository: responseRepository,
	}
}

func (r *ResponseQuery) Show(
	ctx context.Context,
	taskId uint,
	pagination *paginate.Pagination,
) (*paginate.Pagination, error) {
	if taskId < 0 {
		return nil, errors.New("invalid task id")
	}

	if pagination == nil {
		return nil, errors.New("pagination is required")
	}

	return r.responseRepository.Show(ctx, taskId, pagination)
}
