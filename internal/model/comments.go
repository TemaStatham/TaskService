package model

import (
	"time"
)

/*
CREATE TABLE "comment" (

	"id" SERIAL PRIMARY KEY,
	"task_id" INTEGER NOT NULL,
	"user_id" INTEGER NOT NULL,
	"comment" TEXT NOT NULL,
	"created_at" TIMESTAMP DEFAULT (CURRENT_TIMESTAMP)

);
*/
type Comment struct {
	ID        uint      `gorm:"column:id;type:SERIAL;primaryKey;autoIncrement"`
	TaskID    uint      `gorm:"column:task_id;type:INTEGER;not null;index"`
	UserID    uint      `gorm:"column:user_id;type:INTEGER;not null;index"`
	Comment   string    `gorm:"column:comment;type:TEXT;type:text;not null"`
	CreatedAt time.Time `gorm:"column:created_at;type:TIMESTAMP;autoCreateTime"`

	Task TaskModel `gorm:"column:task;type:TEXT;not null"`
	User UserModel `gorm:"column:user;type:TEXT;not null"`
}

func (Comment) TableName() string {
	return "comment"
}
