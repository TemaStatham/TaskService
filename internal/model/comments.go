package model

import (
	"gorm.io/gorm"
	"time"
)

type Comment struct {
	ID        uint           `gorm:"primaryKey;autoIncrement"`
	TaskID    uint           `gorm:"not null;index"`
	UserID    uint           `gorm:"not null;index"`
	Comment   string         `gorm:"type:text;not null"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
