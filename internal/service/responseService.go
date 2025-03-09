package service

import (
	"context"
	"errors"
	"github.com/TemaStatham/TaskService/internal/handler/request"
	"github.com/TemaStatham/TaskService/pkg/paginate"
)

type ResponseRepository interface {
	Create(ctx context.Context, taskId, userId uint, status string) (uint, error)
	Show(ctx context.Context, taskId uint, pagination *paginate.Pagination) (*paginate.Pagination, error)
	Update(ctx context.Context, id uint, status string) error
}

type ResponseService struct {
	ResponseRepository
}

func NewResponseService(repository ResponseRepository) *ResponseService {
	return &ResponseService{
		repository,
	}
}

func (r *ResponseService) Create(ctx context.Context, dto request.CreateResponseRequest, user uint) (uint, error) {
	if dto.UserId < 0 {
		return 0, errors.New("invalid user id")
	}

	id, err := r.ResponseRepository.Create(ctx, dto.TaskId, dto.UserId, dto.Status)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *ResponseService) Show(ctx context.Context, dto request.GetResponseRequest, user uint) (*paginate.Pagination, error) {
	if dto.TaskId < 0 {
		return &paginate.Pagination{}, errors.New("invalid task id")
	}

	pag, err := r.ResponseRepository.Show(ctx, dto.TaskId, &dto.Pagination)
	if err != nil {
		return &paginate.Pagination{}, err
	}

	return pag, nil
}

func (r *ResponseService) Update(ctx context.Context, dto request.UpdateResponseRequest, user uint) error {
	if dto.ID < 0 {
		return errors.New("invalid id")
	}

	return r.ResponseRepository.Update(ctx, dto.ID, dto.Status)
}
