package service

import (
	"context"
	"github.com/TemaStatham/TaskService/client/pkg/app/response/model"
)

type ResponseServiceInterface interface {
	Update(ctx context.Context, id uint, status uint) error
	Create(ctx context.Context, taskId, userId uint, status uint) (uint, error)
}

type ResponseService struct {
	responseRepository model.ResponseRepositoryInterface
}

func NewResponseService(responseRepository model.ResponseRepositoryInterface) *ResponseService {
	return &ResponseService{
		responseRepository: responseRepository,
	}
}

func (r *ResponseService) Update(ctx context.Context, id uint, status uint) error {
	return r.responseRepository.Update(ctx, id, status)
}

func (r *ResponseService) Create(ctx context.Context, taskId, userId uint, status uint) (uint, error) {
	// TODO: написать логику получения статуса, и какой id передать нужно понять. Вынести в data
	return r.responseRepository.Create(ctx, model.ResponseModel{UserID: userId, TaskID: taskId, StatusID: status})
}
