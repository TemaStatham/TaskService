package model

import (
	"context"
	"github.com/TemaStatham/TaskService/taskservice/pkg/app/comment/data"
	"github.com/TemaStatham/TaskService/taskservice/pkg/infrastructure/lib/paginate"
	"time"
)

type CommentModel struct {
	ID        uint      `gorm:"column:id;type:SERIAL;primaryKey;autoIncrement"`
	TaskID    uint      `gorm:"column:task_id;type:INTEGER;not null;index"`
	UserID    uint      `gorm:"column:user_id;type:INTEGER;not null;index"`
	Comment   string    `gorm:"column:comment;type:TEXT;type:text;not null"`
	CreatedAt time.Time `gorm:"column:created_at;type:TIMESTAMP;autoCreateTime"`
}

func (CommentModel) TableName() string {
	return "comment"
}

type CommentReadRepositoryInterface interface {
	Show(
		ctx context.Context,
		taskId uint,
		page int,
		limit int,
	) (*paginate.Pagination, error)
}

type CommentRepositoryInterface interface {
	CommentReadRepositoryInterface
	Create(ctx context.Context, comment data.CreateComment, userId uint) (uint, error)
}
