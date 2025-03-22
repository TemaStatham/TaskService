package postgres

import (
	"context"
	"github.com/TemaStatham/TaskService/taskservice/pkg/app/approve/data"
	"github.com/TemaStatham/TaskService/taskservice/pkg/app/approve/model"
	"github.com/TemaStatham/TaskService/taskservice/pkg/infrastructure/lib/paginate"
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

func (a *ApproveRepository) Create(ctx context.Context, approve data.CreateApprove, status uint) error {
	approveModel := model.ApproveTaskModel{
		TaskID:   approve.TaskID,
		UserID:   approve.UserID,
		StatusID: status,
	}

	res := a.db.WithContext(ctx).Create(&approveModel)
	if res.Error != nil {
		return res.Error
	}

	fileModel := model.File{
		SRC: approve.File,
	}

	res = a.db.WithContext(ctx).Create(&fileModel)
	if res.Error != nil {
		return res.Error
	}

	approveFileModel := model.ApproveFile{
		UserID:        approve.UserID,
		FileID:        fileModel.ID,
		ApproveTaskID: approve.TaskID,
	}

	res = a.db.WithContext(ctx).Create(&approveFileModel)
	if res.Error != nil {
		return res.Error
	}

	return nil
}

// получает все отправления потжтверждения для заявки
func (a *ApproveRepository) Show(ctx context.Context, dto data.ShowApproves) (paginate.Pagination, error) {
	approves := []*model.ApproveTaskModel{}

	query := a.db.WithContext(ctx).
		Model(&model.ApproveTaskModel{}).
		Joins("LEFT JOIN approve_file ON approve_file.approve_task_id = approve_task.id").
		Joins("LEFT JOIN file ON file.id = approve_file.file_id").
		Where("approve_task.task_id = ?", dto.TaskID)

	var total int64
	if err := query.Model(&model.ApproveTaskModel{}).Count(&total).Error; err != nil {
		return paginate.Pagination{}, err
	}

	offset := (dto.Page - 1) * dto.Limit
	if err := query.Offset(int(offset)).Limit(int(dto.Limit)).Find(&approves).Error; err != nil {
		return paginate.Pagination{}, err
	}

	pagination := paginate.Pagination{int(dto.Limit), int(dto.Page), total, approves}

	return pagination, nil
}

func (a *ApproveRepository) Update(ctx context.Context, dto data.SetStatusApprove) error {
	approve := &model.ApproveTaskModel{}

	res := a.db.WithContext(ctx).Where("id = ?", dto.ID).First(&approve)
	if res.Error != nil {
		return res.Error
	}

	approve.StatusID = dto.Status
	approve.Score = dto.Score
	approve.Approved = &dto.Approved

	if err := a.db.WithContext(ctx).Save(approve).Error; err != nil {
		return err
	}

	return nil
}
