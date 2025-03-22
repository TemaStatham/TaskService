package data

type GetTasksUsers struct {
	TaskID         uint  `json:"task_id"`
	Page           int   `json:"page,omitempty;query:page"`
	Limit          int   `json:"limit,omitempty;query:limit"`
	IsCoordinators *bool `json:"is_coordinators"`
}

type GetTasksUsersResponse struct {
	// users paginations
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
