package data

import "github.com/TemaStatham/TaskService/client/pkg/app/paginate"

type CreateResponse struct {
	TaskId uint `json:"task_id"`
	Status uint `json:"status"`
}

type UpdateResponse struct {
	ID     uint `json:"response_id"`
	Status uint `json:"status"`
}

type GetResponses struct {
	TaskId     uint
	Pagination *paginate.Pagination
}
