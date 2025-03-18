package postgres

import (
	"context"
	"github.com/TemaStatham/TaskService/taskservice/pkg/app/task/model"
	"gorm.io/gorm"
)

type TaskCategoryRepository struct {
	db *gorm.DB
}

func NewTaskCategoryRepository(db *gorm.DB) *TaskCategoryRepository {
	return &TaskCategoryRepository{db: db}
}

func (repo *TaskCategoryRepository) Create(ctx context.Context, taskID, categoryID uint) error {
	taskCategory := model.TaskCategory{
		TaskID:     taskID,
		CategoryID: categoryID,
	}

	if err := repo.db.WithContext(ctx).Create(&taskCategory).Error; err != nil {
		return err
	}

	return nil
}
