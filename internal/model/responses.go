package model

import "time"

type Response struct {
	ID        uint      `gorm:"primaryKey;autoIncrement"`
	TaskID    uint      `gorm:"not null;index"`
	UserID    uint      `gorm:"not null;index"`
	Status    string    `gorm:"size:50;default:'Откликнулся'"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}
