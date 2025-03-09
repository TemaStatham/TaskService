package request

import (
	"github.com/TemaStatham/TaskService/pkg/paginate"
	"time"
)

type CreateTaskRequest struct {
	OrganizationID    int       `json:"organization_id" binding:"required"`
	Title             string    `json:"title" binding:"required"`
	Type              int16     `json:"type" binding:"required"`
	Description       string    `json:"description" binding:"required"`
	Location          string    `json:"location" binding:"required"`
	TaskDate          time.Time `json:"task-date" binding:"required"`
	ParticipantsCount int       `json:"participants_count" binding:"required"`
	MaxScore          int       `json:"max_score" binding:"required"`
	Status            *int16    `json:"status" binding:"required"`
}

type UpdateTaskRequest struct {
	ID                uint      `json:"id" binding:"required"`
	OrganizationID    int       `json:"organization_id" binding:"required"`
	Title             string    `json:"title" binding:"required"`
	Type              int16     `json:"type" binding:"required"`
	Description       string    `json:"description" binding:"required"`
	Location          string    `json:"location" binding:"required"`
	TaskDate          time.Time `json:"task-date" binding:"required"`
	ParticipantsCount int       `json:"participants_count" binding:"required"`
	MaxScore          int       `json:"max_score" binding:"required"`
	Status            *int16    `json:"status" binding:"required"`
}

type DeleteTaskRequest struct {
	ID uint `json:"id" binding:"required"`
}

type GetTaskRequest struct {
	ID uint `json:"id" binding:"required"`
}

type GetTasksRequest struct {
	Pagination paginate.Pagination `json:"pagination" binding:"required"`
}
