package data

type CreateApprove struct {
	TaskID   uint  `json:"task_id"`
	UserID   uint  `json:"user_id"`
	StatusID uint  `json:"status_id"`
	Score    uint  `json:"score"`
	Approved *uint `json:"approved"`
}
