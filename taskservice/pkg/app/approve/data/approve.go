package data

type CreateApprove struct {
	TaskID uint   `json:"task_id"`
	UserID uint   `json:"user_id"`
	File   string `json:"file"`
}

type RejectApprove struct {
	ID       uint `json:"id"`
	Approved uint `json:"approved"`
}

type ConfirmApprove struct {
	ID       uint `json:"id"`
	Score    uint `json:"score"`
	Approved uint `json:"approved"`
}

type SetStatusApprove struct {
	ID       uint `json:"id"`
	Score    uint `json:"score"`
	Approved uint `json:"approved"`
	Status   uint `json:"status"`
}

type ShowApproves struct {
	TaskID uint `json:"task_id"`
	Page   uint `json:"page"`
	Limit  uint `json:"limit"`
}

type ShowApprovesResponse struct {
	UserID uint `json:"userID"`
	Score  uint `json:"score"`
}
