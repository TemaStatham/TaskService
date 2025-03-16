package postgres

import (
	"context"
	"errors"
	"github.com/TemaStatham/TaskService/client/pkg/app/paginate"
	"github.com/TemaStatham/TaskService/client/pkg/app/task/model"
	"gorm.io/gorm"
)

type TaskUserRepository struct {
	db *gorm.DB
}

func NewTaskUserPostgresRepository(db *gorm.DB) *TaskUserRepository {
	return &TaskUserRepository{db: db}
}

func (tu *TaskUserRepository) GetUsers(ctx context.Context, taskID uint, pagination *paginate.Pagination, isCoordinators *bool) (*paginate.Pagination, error) {
	var tasksusers []model.TaskUser
	query := tu.db.WithContext(ctx).Where("task_id = ?", taskID)

	if isCoordinators != nil {
		query = query.Where("is_coordinator = ?", *isCoordinators)
	}

	query = query.Scopes(paginate.Paginate(tasksusers, pagination, tu.db))

	res := query.Find(&tasksusers)
	if res.Error != nil {
		return nil, res.Error
	}

	pagination.Rows = tasksusers

	return pagination, nil
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
