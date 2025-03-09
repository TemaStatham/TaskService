package request

import "github.com/TemaStatham/TaskService/pkg/paginate"

type CreateResponseRequest struct {
	TaskId uint   `json:"task_id"`
	UserId uint   `json:"user_id"`
	Status string `json:"status"`
}

type GetResponseRequest struct {
	TaskId     uint                `json:"task_id"`
	Pagination paginate.Pagination `json:"pagination"`
}

type UpdateResponseRequest struct {
	ID     uint   `json:"response_id"`
	Status string `json:"status"`
}
