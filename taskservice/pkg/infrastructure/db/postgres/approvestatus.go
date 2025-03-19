package postgres

import (
	"context"
	"github.com/TemaStatham/TaskService/taskservice/pkg/app/approve/model"
	"gorm.io/gorm"
)

type ApproveStatusRepository struct {
	db *gorm.DB
}

func (a ApproveStatusRepository) Get(ctx context.Context, status string) (uint, error) {
	var result model.ApproveTaskStatusModel
	res := a.db.WithContext(ctx).
		Where("name = ?", status).
		First(&result)

	if res.Error != nil {
		return 0, res.Error
	}
	return result.ID, nil
}

func NewApproveStatusRepository(db *gorm.DB) model.ApproveTaskStatusReadRepositoryInterface {
	return &ApproveStatusRepository{
		db: db,
	}
}
