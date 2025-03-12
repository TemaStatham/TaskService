package model

/*
CREATE TABLE "organization" (

	"id" SERIAL PRIMARY KEY,
	"email" VARCHAR(150) UNIQUE NOT NULL,
	"phone" VARCHAR(18) NOT NULL,
	"name" VARCHAR(255) NOT NULL,
	"inn" VARCHAR(12) UNIQUE NOT NULL,
	"legal_address" VARCHAR(255) NOT NULL,
	"actual_address" VARCHAR(255) NOT NULL,
	"status_id" INTEGER NOT NULL,
	"full_name_owner" VARCHAR(255) NOT NULL,
	"created_at" TIMESTAMP DEFAULT (CURRENT_TIMESTAMP),
	"updated_at" TIMESTAMP DEFAULT (CURRENT_TIMESTAMP)

);
*/
type OrganizationModel struct {
	ID       int    `gorm:"column:id;type:SERIAL;primaryKey;autoIncrement" json:"id"`
	Name     string `gorm:"column:name;type:VARCHAR(255);not null;type:varchar(255)" json:"name" binding:"required"`
	StatusID int    `gorm:"column:status_id;type:INTEGER;not null" json:"status_id" binding:"required"`

	Status OrganizationStatusModel `gorm:"column:status;type:TEXT;not null" json:"status" binding:"required"`
}

func (OrganizationModel) TableName() string {
	return "organization"
}

/*
CREATE TABLE "organization_statuses" (

	"id" SERIAL PRIMARY KEY,
	"name" VARCHAR(20) UNIQUE NOT NULL

);
*/
type OrganizationStatusModel struct {
	ID   int    `gorm:"column:id;type:SERIAL;primaryKey;autoIncrement" json:"id"`
	Name string `gorm:"column:name;type:VARCHAR(20);not null" json:"name" binding:"required"`
}

func (OrganizationStatusModel) TableName() string {
	return "organization_statuses"
}
