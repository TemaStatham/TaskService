package model

/*
CREATE TABLE "response" (

	"id" SERIAL PRIMARY KEY,
	"task_id" INTEGER NOT NULL,
	"user_id" INTEGER NOT NULL,
	"status_id" INTEGER DEFAULT 1

);
*/
type ResponseModel struct {
	ID       uint `gorm:"column:id;primaryKey;type:SERIAL;autoIncrement"`
	TaskID   uint `gorm:"column:task_id;not null;type:INTEGER;index"`
	UserID   uint `gorm:"column:user_id;not null;type:INTEGER;index"`
	StatusID uint `gorm:"column:status_id;type:INTEGER;default:1"`

	Task   TaskModel           `gorm:"column:task;type:TEXT;not null"`
	User   UserModel           `gorm:"column:user;type:TEXT;not null"`
	Status ResponseStatusModel `gorm:"column:status;type:TEXT;not null"`
}

/*
CREATE TABLE "response_status" (

	"id" SERIAL PRIMARY KEY,
	"name" VARCHAR(255)

);
*/
type ResponseStatusModel struct {
	ID   uint    `gorm:"column:id;primaryKey;type:SERIAL;autoIncrement"`
	Name *string `gorm:"column:name;type:VARCHAR(255)"`
}
