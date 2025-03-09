package request

import "github.com/TemaStatham/TaskService/pkg/paginate"

type CreateCommentRequest struct {
	TaskID  uint   `gorm:"not null;index"`
	UserID  uint   `gorm:"not null;index"`
	Comment string `gorm:"type:text;not null"`
}

type ShowCommentRequest struct {
	TaskID uint                `gorm:"not null;index"`
	Pag    paginate.Pagination `gorm:"not null;index"`
}
