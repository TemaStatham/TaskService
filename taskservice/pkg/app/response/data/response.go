package data

type CreateResponse struct {
	TaskId uint `json:"task_id"`
	Status uint `json:"status"`
}

type UpdateResponse struct {
	ID     uint `json:"response_id"`
	Status uint `json:"status"`
}

type GetResponses struct {
	TaskId uint `json:"task_id"`
	Page   int  `json:"page,omitempty;query:page"`
	Limit  int  `json:"limit,omitempty;query:limit"`
}
