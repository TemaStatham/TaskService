package service

import (
	"context"
	"errors"
	"github.com/TemaStatham/TaskService/internal/handler/request"
	"github.com/TemaStatham/TaskService/pkg/paginate"
)

type CommentRepository interface {
	Create(ctx context.Context, taskId, userId uint, comment string) (uint, error)
	Show(
		ctx context.Context,
		taskId uint,
		pagination *paginate.Pagination,
	) (*paginate.Pagination, error)
}

type CommentService struct {
	CommentRepository
}

func NewCommentService(repository CommentRepository) *CommentService {
	return &CommentService{
		repository,
	}
}

func (c *CommentService) Create(
	ctx context.Context,
	dto *request.CreateCommentRequest,
	user uint,
) (uint, error) {
	if dto.TaskID < 0 {
		return 0, errors.New("invalid task id")
	}

	id, err := c.CommentRepository.Create(ctx, dto.TaskID, user, dto.Comment)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (c *CommentService) Show(
	ctx context.Context,
	dto *request.ShowCommentRequest,
) (*paginate.Pagination, error) {
	if dto.TaskID < 0 {
		return &paginate.Pagination{}, errors.New("invalid task id")
	}

	pagination, err := c.CommentRepository.Show(ctx, dto.TaskID, &dto.Pagination)
	if err != nil {
		return &paginate.Pagination{}, err
	}

	return pagination, nil
}
