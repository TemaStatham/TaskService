package model

import "time"

/*
CREATE TABLE "task" (

	"id" SERIAL PRIMARY KEY,
	"organization_id" INTEGER NOT NULL,
	"name" VARCHAR(255) NOT NULL,
	"type_id" INTEGER NOT NULL,
	"description" TEXT NOT NULL,
	"location" VARCHAR(255) NOT NULL,
	"task_date" TIMESTAMP NOT NULL,
	"participants_count" INTEGER,
	"max_score" INTEGER,
	"status_id" INTEGER DEFAULT 1,
	"created_at" TIMESTAMP DEFAULT (CURRENT_TIMESTAMP),
	"updated_at" TIMESTAMP DEFAULT (CURRENT_TIMESTAMP)

);
*/
type TaskModel struct {
	ID                uint      `gorm:"column:id;type:SERIAL;primaryKey;autoIncrement" json:"id"`
	OrganizationID    int       `gorm:"column:organization_id;type:INTEGER;not null;index" json:"organization_id"`
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

	Users        []UserModel       `gorm:"many2many:task_user;" json:"users"`
	Categories   []CategoryModel   `gorm:"many2many:task_category;" json:"categories"`
	Organization OrganizationModel `gorm:"foreignKey:OrganizationID" json:"organization"`
	TaskType     TaskTypeModel     `gorm:"foreignKey:TypeID" json:"task_type"`
	TaskStatus   TaskStatusModel   `gorm:"foreignKey:StatusID" json:"task_status"`
}

func (TaskModel) TableName() string {
	return "task"
}

/*
CREATE TABLE "task_user" (

	"id" SERIAL PRIMARY KEY,
	"task_id" INTEGER NOT NULL,
	"user_id" INTEGER NOT NULL,
	"is_coordinator" BOOLEAN DEFAULT false

);
*/
type TaskUser struct {
	ID            uint `gorm:"column:id;type:SERIAL;primaryKey;autoIncrement"`
	TaskID        uint `gorm:"column:task_id;type:INTEGER;not null;index"`
	UserID        uint `gorm:"column:user_id;type:INTEGER;not null;index"`
	IsCoordinator bool `gorm:"column:is_coordinator;type:BOOLEAN;default:false"`
}

func (TaskUser) TableName() string {
	return "task_user"
}

/*
CREATE TABLE "category" (

	"id" SERIAL PRIMARY KEY,
	"name" VARCHAR(100) UNIQUE NOT NULL

);
*/
type CategoryModel struct {
	ID   uint   `gorm:"column:id;type:SERIAL;primaryKey;autoIncrement"`
	Name string `gorm:"column:name;type:VARCHAR(100);unique;not null"`
}

func (CategoryModel) TableName() string {
	return "category"
}

/*
CREATE TABLE "task_category" (

	"id" SERIAL PRIMARY KEY,
	"task_id" INTEGER NOT NULL,
	"category_id" INTEGER NOT NULL

);
*/

type TaskCategory struct {
	ID         uint `gorm:"column:id;type:SERIAL;primaryKey;autoIncrement" json:"id"`
	TaskID     uint `gorm:"column:task_id;type:INTEGER;not null;index" json:"task_id"`
	CategoryID uint `gorm:"column:category_id;type:INTEGER;not null;index" json:"category_id"`
}

func (TaskCategory) TableName() string {
	return "task_category"
}

/*
CREATE TABLE "task_type" (

	"id" SERIAL PRIMARY KEY,
	"name" VARCHAR(255)

);
*/
type TaskTypeModel struct {
	ID   uint    `gorm:"column:id;type:SERIAL;primaryKey;autoIncrement" json:"id"`
	Name *string `gorm:"column:name;type:VARCHAR(255)" json:"name"`
}

func (TaskTypeModel) TableName() string {
	return "task_type"
}

/*
CREATE TABLE "task_status" (

	"id" SERIAL PRIMARY KEY,
	"name" VARCHAR(255)

);
*/
type TaskStatusModel struct {
	ID   uint    `gorm:"column:id;type:SERIAL;primaryKey;autoIncrement" json:"id"`
	Name *string `gorm:"column:name;type:VARCHAR(255)" json:"name"`
}

func (TaskStatusModel) TableName() string {
	return "task_status"
}
