package model

import (
	"context"
	"github.com/TemaStatham/TaskService/client/pkg/app/approve/data"
	"time"
)

type ApproveTaskStatusModel struct {
	ID   uint    `gorm:"column:id;type:SERIAL;primaryKey;autoIncrement"`
	Name *string `gorm:"column:name;type:TEXT;not null"`
}

func (ApproveTaskStatusModel) TableName() string {
	return "approve_task_status"
}

type ApproveTaskModel struct {
	ID        uint      `gorm:"column:id;type:SERIAL;primaryKey;autoIncrement"`
	TaskID    uint      `gorm:"column:task_id;type:INTEGER;not null;index"`
	UserID    uint      `gorm:"column:user_id;type:INTEGER;not null;index"`
	StatusID  uint      `gorm:"column:status_id;type:INTEGER;not null;index"`
	Score     uint      `gorm:"column:score;type:INTEGER;not null;index"`
	Approved  *uint     `gorm:"column:approved;type:INTEGER;index"`
	CreatedAt time.Time `gorm:"column:created_at;type:TIMESTAMP;not null"`
}

func (ApproveTaskModel) TableName() string {
	return "approve_task"
}

type ApproveFile struct {
	ID            uint `gorm:"column:id;type:SERIAL;primaryKey;autoIncrement"`
	UserID        uint `gorm:"column:user_id;type:INTEGER;not null;index"`
	FileID        uint `gorm:"column:file_id;type:INTEGER;not null;index"`
	ApproveTaskID uint `gorm:"column:approve_task_id;type:INTEGER;not null;index"`
}

func (ApproveFile) TableName() string {
	return "approve_file"
}

type File struct {
	ID         uint      `gorm:"column:id;type:SERIAL;primaryKey;autoIncrement"`
	SRC        string    `gorm:"column:src;type:TEXT;not null"`
	UploadedAt time.Time `gorm:"column:uploaded_at;type:TIMESTAMP;not null"`
}

func (File) TableName() string {
	return "file"
}

type ApproveRepositoryInterface interface {
	Create(ctx context.Context, approve data.CreateApprove) error
}
