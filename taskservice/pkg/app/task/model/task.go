package model

import (
	"context"
	"github.com/TemaStatham/TaskService/taskservice/pkg/app/organization/model"
	"github.com/TemaStatham/TaskService/taskservice/pkg/app/task/data"
	"github.com/TemaStatham/TaskService/taskservice/pkg/infrastructure/lib/paginate"
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
		dto data.GetAllTasks,
		user uint,
		organizations []model.Organization,
	) (*paginate.Pagination, error)
	GetByOrganization(ctx context.Context, organizationID uint) (*paginate.Pagination, error)
}

type TaskRepositoryInterface interface {
	TaskReadRepositoryInterface
	Delete(ctx context.Context, id uint) error
	Update(ctx context.Context, task *data.UpdateTask) error
	Create(ctx context.Context, task *data.CreateTask) (uint, error)
}
