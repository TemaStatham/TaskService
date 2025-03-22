package service

import (
	"context"
	"github.com/TemaStatham/TaskService/taskservice/pkg/app/response/model"
)

type ResponseServiceInterface interface {
	Update(ctx context.Context, id uint, statusName string) error
	Create(ctx context.Context, taskId, userId uint) (uint, error)
}

type ResponseService struct {
	responseRepository       model.ResponseRepositoryInterface
	responseStatusRepository model.ResponseStatusRepositoryInterface
}

func NewResponseService(responseRepository model.ResponseRepositoryInterface,
	responseStatusRepository model.ResponseStatusRepositoryInterface) *ResponseService {
	return &ResponseService{
		responseRepository:       responseRepository,
		responseStatusRepository: responseStatusRepository,
	}
}

func (r *ResponseService) Update(ctx context.Context, id uint, statusName string) error {
	status, err := r.responseStatusRepository.GetStatus(ctx, statusName)
	if err != nil {
		return err
	}

	return r.responseRepository.Update(ctx, id, status.ID)
}

func (r *ResponseService) Create(ctx context.Context, taskId, userId uint) (uint, error) {
	status, err := r.responseStatusRepository.GetStatus(ctx, "На рассмотрении")
	if err != nil {
		return 0, err
	}

	return r.responseRepository.Create(ctx, model.ResponseModel{UserID: userId, TaskID: taskId, StatusID: status.ID})
}
