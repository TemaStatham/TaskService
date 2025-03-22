package model

type Organization struct {
	ID       uint   `json:"id"`
	Name     string `json:"name" binding:"required"`
	StatusID uint   `json:"status_id" binding:"required"`
}
