package model

type User struct {
	ID      uint    `json:"id" binding:"required"`
	Surname *string `json:"surname" binding:"required"`
	Name    string  `json:"name" binding:"required"`
	IsAdmin bool    `json:"is_admin" binding:"required"`
}
