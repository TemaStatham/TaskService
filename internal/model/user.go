package model

type User struct {
	ID       uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Email    string `gorm:"unique;not null" json:"email" binding:"required"`
	PassHash []byte `gorm:"not null" json:"password" binding:"required"`
	Role     int16  `gorm:"not null" json:"role" binding:"required"`
}
