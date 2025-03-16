package data

import "github.com/TemaStatham/TaskService/client/pkg/app/paginate"

type GetTasksUsers struct {
	TaskID         uint                `json:"task_id"`
	Pagination     paginate.Pagination `json:"pagination"`
	IsCoordinators *bool               `json:"is_coordinators"`
}

type AddTasksUsers struct {
	TaskID        uint `json:"task_id"`
	UserID        uint `json:"user_id"`
	IsCoordinator bool `json:"is_coordinator"`
}

type DeleteTasksUsers struct {
	TaskID uint `json:"task_id"`
	UserID uint `json:"user_id"`
}
