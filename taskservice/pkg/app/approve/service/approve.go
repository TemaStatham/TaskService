package service

import (
	"context"
	"github.com/TemaStatham/TaskService/taskservice/pkg/app/approve/data"
	"github.com/TemaStatham/TaskService/taskservice/pkg/app/approve/model"
)

type ApproveServiceInterface interface {
	Create(ctx context.Context, approve data.CreateApprove) error
}

type ApproveService struct {
	repository model.ApproveRepositoryInterface
}

func NewApproveService(repository model.ApproveRepositoryInterface) *ApproveService {
	return &ApproveService{
		repository: repository,
	}
}

func (a *ApproveService) Create(ctx context.Context, approve data.CreateApprove) error {
	return a.repository.Create(ctx, approve)
}
