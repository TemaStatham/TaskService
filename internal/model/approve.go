package model

import "time"

/*
CREATE TABLE "approve_task_status" (

	"id" SERIAL PRIMARY KEY,
	"name" VARCHAR(255)

);
*/
type ApproveTaskStatusModel struct {
	ID   uint    `gorm:"column:id;type:SERIAL;primaryKey;autoIncrement"`
	Name *string `gorm:"column:name;type:TEXT;not null"`
}

func (ApproveTaskStatusModel) TableName() string {
	return "approve_task_status"
}

/*
CREATE TABLE "approve_task" (

	"id" SERIAL PRIMARY KEY,
	"task_id" INTEGER NOT NULL,
	"user_id" INTEGER NOT NULL,
	"status_id" INTEGER DEFAULT 1,
	"score" INTEGER DEFAULT 0,
	"approved" INTEGER,
	"created_at" TIMESTAMP DEFAULT (NOW())

);
*/
type ApproveTaskModel struct {
	ID        uint      `gorm:"column:id;type:SERIAL;primaryKey;autoIncrement"`
	TaskID    uint      `gorm:"column:task_id;type:INTEGER;not null;index"`
	UserID    uint      `gorm:"column:user_id;type:INTEGER;not null;index"`
	StatusID  uint      `gorm:"column:status_id;type:INTEGER;not null;index"`
	Score     uint      `gorm:"column:score;type:INTEGER;not null;index"`
	Approved  *uint     `gorm:"column:approved;type:INTEGER;index"`
	CreatedAt time.Time `gorm:"column:created_at;type:TIMESTAMP;not null"`

	Status ApproveTaskStatusModel `gorm:"column:status;type:TEXT;not null"`
}

func (ApproveTaskModel) TableName() string {
	return "approve_task"
}

/*
CREATE TABLE "approve_file" (

	"id" SERIAL PRIMARY KEY,
	"user_id" INTEGER NOT NULL,
	"file_id" INTEGER NOT NULL,
	"approve_task_id" INTEGER NOT NULL

);
*/
type ApproveFile struct {
	ID            uint `gorm:"column:id;type:SERIAL;primaryKey;autoIncrement"`
	UserID        uint `gorm:"column:user_id;type:INTEGER;not null;index"`
	FileID        uint `gorm:"column:file_id;type:INTEGER;not null;index"`
	ApproveTaskID uint `gorm:"column:approve_task_id;type:INTEGER;not null;index"`
}

/*
CREATE TABLE "file" (

	"id" SERIAL PRIMARY KEY,
	"src" TEXT NOT NULL,
	"uploaded_at" TIMESTAMP DEFAULT (CURRENT_TIMESTAMP)

);
*/
type File struct {
	ID         uint      `gorm:"column:id;type:SERIAL;primaryKey;autoIncrement"`
	SRC        string    `gorm:"column:src;type:TEXT;not null"`
	UploadedAt time.Time `gorm:"column:uploaded_at;type:TIMESTAMP;not null"`
}
