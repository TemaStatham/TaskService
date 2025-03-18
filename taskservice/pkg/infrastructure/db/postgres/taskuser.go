package postgres

import (
	"context"
	"errors"
	"github.com/TemaStatham/TaskService/taskservice/pkg/app/task/model"
	"github.com/TemaStatham/TaskService/taskservice/pkg/infrastructure/lib/paginate"
	"gorm.io/gorm"
)

type TaskUserRepository struct {
	db *gorm.DB
}

func NewTaskUserPostgresRepository(db *gorm.DB) *TaskUserRepository {
	return &TaskUserRepository{db: db}
}

func (tu *TaskUserRepository) GetUsers(
	ctx context.Context,
	taskID uint,
	page int,
	limit int,
	isCoordinators *bool,
) (*paginate.Pagination, error) {
	var taskUsers []model.TaskUser
	query := tu.db.WithContext(ctx).Where("task_id = ?", taskID)

	if isCoordinators != nil {
		query = query.Where("is_coordinator = ?", *isCoordinators)
	}

	var total int64
	if err := query.Model(&model.TaskUser{}).Count(&total).Error; err != nil {
		return nil, err
	}

	// todo: вынести в сервис
	if limit <= paginate.MinLimit {
		limit = paginate.DefaultLimit
	}

	if page <= paginate.MinPage {
		page = paginate.DefaultPage
	}

	offset := (page - 1) * limit
	if err := query.Offset(offset).Limit(limit).Find(&taskUsers).Error; err != nil {
		return nil, err
	}

	pagination := paginate.Pagination{limit, page, total, taskUsers}
	return &pagination, nil
}

func (tu *TaskUserRepository) Add(ctx context.Context, userID, taskID uint, isCoordinator bool) error {
	if isExist, err := tu.isExist(ctx, userID, taskID); err != nil || !isExist {
		return err
	}
	taskuser := model.TaskUser{
		TaskID:        taskID,
		UserID:        userID,
		IsCoordinator: isCoordinator,
	}

	res := tu.db.WithContext(ctx).Create(&taskuser)

	return res.Error
}

func (tu *TaskUserRepository) Delete(ctx context.Context, userID, taskID uint) error {
	if isExist, err := tu.isExist(ctx, userID, taskID); err != nil || !isExist {
		return err
	}
	taskuser := model.TaskUser{
		TaskID: taskID,
		UserID: userID,
	}
	res := tu.db.WithContext(ctx).Delete(&taskuser)

	return res.Error
}

func (tu *TaskUserRepository) isExist(ctx context.Context, userID, taskID uint) (bool, error) {
	var existing model.TaskUser

	err := tu.db.WithContext(ctx).
		Where("task_id = ? AND user_id = ?", taskID, userID).
		First(&existing).Error

	if err == nil {
		return true, nil
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, nil
	}

	return false, err
}

func (tu *TaskUserRepository) DeleteAllByUserID(ctx context.Context, userID uint) error {
	if err := tu.db.WithContext(ctx).Where("user_id = ?", userID).Delete(&model.TaskUser{}).Error; err != nil {
		return err
	}
	return nil
}
