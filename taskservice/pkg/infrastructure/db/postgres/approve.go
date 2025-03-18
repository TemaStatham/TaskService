package postgres

import (
	"context"
	"github.com/TemaStatham/TaskService/taskservice/pkg/app/approve/data"
	"github.com/TemaStatham/TaskService/taskservice/pkg/app/approve/model"
	"gorm.io/gorm"
)

type ApproveRepository struct {
	db *gorm.DB
}

func NewApproveRepository(db *gorm.DB) *ApproveRepository {
	return &ApproveRepository{
		db: db,
	}
}

func (a *ApproveRepository) Create(ctx context.Context, approve data.CreateApprove) error {
	approveModel := model.ApproveTaskModel{
		TaskID:   approve.TaskID,
		UserID:   approve.UserID,
		StatusID: approve.StatusID,
		Score:    approve.Score,
		Approved: approve.Approved,
	}

	res := a.db.WithContext(ctx).Create(&approveModel)
	if res.Error != nil {
		return res.Error
	}

	return nil
}
