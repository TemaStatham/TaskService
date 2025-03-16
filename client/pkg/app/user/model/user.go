package model

/*
CREATE TABLE "users" (

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

// todo: написать метод который будет получать UserOrganization:IsOwner
/*func (UserModel) TableName() string {
	return "user"
}
*/
/*type UserOrganization struct {
	ID             uint `gorm:"column:id;primaryKey;autoIncrement" json:"id" binding:"required"`
	IsOwner        bool `gorm:"column:is_owner;" json:"is_owner" binding:"required"`
	UserID         uint `gorm:"column:user_id;not null" json:"user_id" binding:"required"`
	OrganizationID uint `gorm:"column:organization_id;not null" json:"organization_id" binding:"required"`
}

func (UserOrganization) TableName() string {
	return "user_organization"
}
*/
