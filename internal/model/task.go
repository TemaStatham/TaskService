package model

import "time"

type Task struct {
	ID                uint           `gorm:"primaryKey;autoIncrement" json:"id"`
	OrganizationID    int            `gorm:"not null;index" json:"organization_id"`
	Title             string         `gorm:"not null" json:"title"`
	Type              int16          `gorm:"not null" json:"type"`
	Description       string         `gorm:"type:text;not null" json:"description"`
	Location          string         `gorm:"not null" json:"location"`
	TaskDate          time.Time      `gorm:"not null" json:"task_date"`
	ParticipantsCount int            `gorm:"not null;default:0" json:"participants_count"`
	MaxScore          int            `gorm:"not null;default:0" json:"max_score"`
	Status            *int16         `gorm:"default:NULL" json:"status"`
	Users             []User         `gorm:"many2many:tasks_users;" json:"users"`
	Organizations     []Organization `gorm:"many2many:tasks_organizations;" json:"organizations"`
	Categories        []Category     `gorm:"many2many:task_category;" json:"categories"`
	CreatedAt         time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt         time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
}
