package model

import "context"

type CategoryModel struct {
	ID   uint   `gorm:"column:id;type:SERIAL;primaryKey;autoIncrement"`
	Name string `gorm:"column:name;type:VARCHAR(100);unique;not null"`
}

func (CategoryModel) TableName() string {
	return "category"
}

type TaskCategory struct {
	ID         uint `gorm:"column:id;type:SERIAL;primaryKey;autoIncrement" json:"id"`
	TaskID     uint `gorm:"column:task_id;type:INTEGER;not null;index" json:"task_id"`
	CategoryID uint `gorm:"column:category_id;type:INTEGER;not null;index" json:"category_id"`
}

func (TaskCategory) TableName() string {
	return "task_category"
}

type TaskCategoryRepositoryInterface interface {
	Create(ctx context.Context, taskID uint, categoryID uint) error
}
