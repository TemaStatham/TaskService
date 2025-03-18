package data

import (
	"time"
)

type CreateTask struct {
	Name              string    `json:"name"`
	Description       string    `json:"description"`
	Location          string    `json:"location"`
	TaskDate          time.Time `json:"task_date"`
	ParticipantsCount *int      `json:"participants_count"`
	MaxScore          *int      `json:"max_score"`

	Organization uint `json:"organization"`
	TaskType     uint `json:"task_type"`
	TaskStatus   uint `json:"task_status"`

	Categories   []uint `json:"categories"`
	Coordinators []uint `json:"coordinators"`
}

type UpdateTask struct {
	ID                uint      `json:"id"`
	Name              string    `json:"name"`
	Description       string    `json:"description"`
	Location          string    `json:"location"`
	TaskDate          time.Time `json:"task_date"`
	ParticipantsCount *int      `json:"participants_count"`
	MaxScore          *int      `json:"max_score"`

	Organization uint `json:"organization"`
	TaskType     uint `json:"task_type"`
	TaskStatus   uint `json:"task_status"`

	Categories []uint `json:"categories"`
}

type DeleteTask struct {
	ID uint `json:"id" binding:"required"`
}

type GetTask struct {
	ID uint `json:"id" binding:"required"`
}

type GetAllTasks struct {
	Page  int `json:"page,omitempty;query:page"`
	Limit int `json:"limit,omitempty;query:limit"`
}
