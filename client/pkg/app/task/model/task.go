package model

import (
	"context"
	"github.com/TemaStatham/TaskService/client/pkg/app/organization/model"
	"github.com/TemaStatham/TaskService/client/pkg/app/paginate"
	"github.com/TemaStatham/TaskService/client/pkg/app/task/data"
	"time"
)

type TaskModel struct {
	ID                uint      `gorm:"column:id;type:SERIAL;primaryKey;autoIncrement" json:"id"`
	OrganizationID    uint      `gorm:"column:organization_id;type:INTEGER;not null" json:"organization_id"`
	Name              string    `gorm:"column:name;type:VARCHAR(255);not null" json:"name"`
	TypeID            uint      `gorm:"column:type_id;type:INTEGER;not null;index" json:"type_id"`
	Description       string    `gorm:"column:description;type:TEXT;not null" json:"description"`
	Location          string    `gorm:"column:location;type:VARCHAR(255);not null" json:"location"`
	TaskDate          time.Time `gorm:"column:task_date;type:TIMESTAMP;not null" json:"task_date"`
	ParticipantsCount *int      `gorm:"column:participants_count;type:INTEGER" json:"participants_count"`
	MaxScore          *int      `gorm:"column:max_score;type:INTEGER" json:"max_score"`
	StatusID          uint      `gorm:"column:status_id;type:INTEGER;default:1;index" json:"status_id"`
	CreatedAt         time.Time `gorm:"column:created_at;type:TIMESTAMP;default:CURRENT_TIMESTAMP;autoCreateTime" json:"created_at"`
	UpdatedAt         time.Time `gorm:"column:updated_at;type:TIMESTAMP;default:CURRENT_TIMESTAMP;autoUpdateTime" json:"updated_at"`
}

func (TaskModel) TableName() string {
	return "task"
}

type TaskUser struct {
	ID            uint `gorm:"column:id;type:SERIAL;primaryKey;autoIncrement"`
	TaskID        uint `gorm:"column:task_id;type:INTEGER;not null;index"`
	UserID        uint `gorm:"column:user_id;type:INTEGER;not null;index"`
	IsCoordinator bool `gorm:"column:is_coordinator;type:BOOLEAN;default:false"`
}

func (TaskUser) TableName() string {
	return "task_user"
}

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

type TaskTypeModel struct {
	ID   uint    `gorm:"column:id;type:SERIAL;primaryKey;autoIncrement" json:"id"`
	Name *string `gorm:"column:name;type:VARCHAR(255)" json:"name"`
}

func (TaskTypeModel) TableName() string {
	return "task_type"
}

type TaskStatusModel struct {
	ID   uint    `gorm:"column:id;type:SERIAL;primaryKey;autoIncrement" json:"id"`
	Name *string `gorm:"column:name;type:VARCHAR(255)" json:"name"`
}

func (TaskStatusModel) TableName() string {
	return "task_status"
}

type TaskReadRepositoryInterface interface {
	Get(ctx context.Context, id uint) (*TaskModel, error)
	GetAll(
		ctx context.Context,
		pagination *paginate.Pagination,
		user uint,
		organizations []model.OrganizationModel,
	) (*paginate.Pagination, error)
}

type TaskRepositoryInterface interface {
	TaskReadRepositoryInterface
	Delete(ctx context.Context, id uint) error
	Update(ctx context.Context, task *data.UpdateTask) error
	Create(ctx context.Context, task *data.CreateTask) (uint, error)
}

type TaskUserReadRepositoryInterface interface {
	GetUsers(ctx context.Context, taskID uint, pagination *paginate.Pagination, isCoordinators *bool) (*paginate.Pagination, error)
}

type TaskUserRepositoryInterface interface {
	TaskUserReadRepositoryInterface
	Add(ctx context.Context, userID, taskID uint, isCoordinator bool) error
	Delete(ctx context.Context, userID, taskID uint) error
}
