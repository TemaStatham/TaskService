package request

import (
	"github.com/TemaStatham/TaskService/internal/model"
	"github.com/TemaStatham/TaskService/pkg/paginate"
	"time"
)

type CreateTaskRequest struct {
	Name              string    `json:"name"`
	Description       string    `json:"description"`
	Location          string    `json:"location"`
	TaskDate          time.Time `json:"task_date"`
	ParticipantsCount *int      `json:"participants_count"`
	MaxScore          *int      `json:"max_score"`

	Users        []model.UserModel       `json:"users"`
	Categories   []model.CategoryModel   `json:"categories"`
	Organization model.OrganizationModel `json:"organization"`
	TaskType     model.TaskTypeModel     `json:"task_type"`
	TaskStatus   model.TaskStatusModel   `json:"task_status"`
}

type UpdateTaskRequest struct {
	ID                uint      `json:"id"`
	Name              string    `json:"name"`
	Description       string    `json:"description"`
	Location          string    `json:"location"`
	TaskDate          time.Time `json:"task_date"`
	ParticipantsCount *int      `json:"participants_count"`
	MaxScore          *int      `json:"max_score"`

	Users        []model.UserModel       `json:"users"`
	Categories   []model.CategoryModel   `json:"categories"`
	Organization model.OrganizationModel `json:"organization"`
	TaskType     model.TaskTypeModel     `json:"task_type"`
	TaskStatus   model.TaskStatusModel   `json:"task_status"`
}

type DeleteTaskRequest struct {
	ID uint `json:"id" binding:"required"`
}

type GetTaskRequest struct {
	ID uint `json:"id" binding:"required"`
}

type GetTasksRequest struct {
	Pagination paginate.Pagination `json:"pagination"`
}
