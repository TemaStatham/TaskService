package model

type User struct {
	ID            uint           `gorm:"primaryKey;autoIncrement" json:"id"`
	Email         string         `gorm:"unique;not null" json:"email" binding:"required"`
	Role          int16          `gorm:"not null" json:"role" binding:"required"`
	Organizations []Organization `gorm:"many2many:users_organizations" json:"organizations" binding:"required"`
}
