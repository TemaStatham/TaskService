package service

import (
	"context"
	"github.com/TemaStatham/TaskService/taskservice/pkg/app/approve/data"
	"github.com/TemaStatham/TaskService/taskservice/pkg/app/approve/model"
	"github.com/TemaStatham/TaskService/taskservice/pkg/infrastructure/lib/paginate"
)

type ApproveServiceInterface interface {
	Create(ctx context.Context, dto data.CreateApprove) error
	Confirm(ctx context.Context, dto data.ConfirmApprove) error
	Reject(ctx context.Context, dto data.RejectApprove) error
	Show(ctx context.Context, dto data.ShowApproves) (paginate.Pagination, error)
}

type ApproveService struct {
	repository        model.ApproveRepositoryInterface
	approvestatusrepo model.ApproveTaskStatusReadRepositoryInterface
}

func NewApproveService(
	repository model.ApproveRepositoryInterface,
	approvestatusrepo model.ApproveTaskStatusReadRepositoryInterface,
) *ApproveService {
	return &ApproveService{
		repository:        repository,
		approvestatusrepo: approvestatusrepo,
	}
}

func (a *ApproveService) Create(ctx context.Context, dto data.CreateApprove) error {
	status, err := a.approvestatusrepo.Get(ctx, "На рассмотрении")
	if err != nil {
		return err
	}

	return a.repository.Create(ctx, dto, status)
}

func (a *ApproveService) Confirm(ctx context.Context, dto data.ConfirmApprove) error {
	status, err := a.approvestatusrepo.Get(ctx, "Принято")
	if err != nil {
		return err
	}

	return a.repository.Update(ctx, data.SetStatusApprove{
		ID:       dto.ID,
		Score:    dto.Score,
		Approved: dto.Approved,
		Status:   status,
	})
}

func (a *ApproveService) Reject(ctx context.Context, dto data.RejectApprove) error {
	status, err := a.approvestatusrepo.Get(ctx, "Отказано")
	if err != nil {
		return err
	}

	return a.repository.Update(ctx, data.SetStatusApprove{
		ID:       dto.ID,
		Score:    0,
		Approved: dto.Approved,
		Status:   status,
	})
}

func (a *ApproveService) Show(ctx context.Context, dto data.ShowApproves) (paginate.Pagination, error) {
	return a.repository.Show(ctx, dto)
}
