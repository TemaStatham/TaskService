package model

/*
CREATE TABLE "user" (

	"id" SERIAL PRIMARY KEY,
	"surname" VARCHAR(255),
	"name" VARCHAR(255) NOT NULL,
	"is_admin" BOOLEAN DEFAULT false

);
*/
type UserModel struct {
	ID      uint    `gorm:"column:id;primaryKey;autoIncrement" json:"id" binding:"required"`
	Surname *string `gorm:"column:surname;" json:"surname" binding:"required"`
	Name    string  `gorm:"column:name;not null" json:"name" binding:"required"`
	IsAdmin bool    `gorm:"column:is_admin;" json:"is_admin" binding:"required"`
}

func (UserModel) TableName() string {
	return "user"
}
