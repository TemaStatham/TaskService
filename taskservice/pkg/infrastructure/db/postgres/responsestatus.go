package postgres

import (
	"context"
	"github.com/TemaStatham/TaskService/taskservice/pkg/app/response/model"
	"gorm.io/gorm"
)

type ResponseStatusRepository struct {
	db *gorm.DB
}

func NewResponseStatusRepository(db *gorm.DB) ResponseStatusRepository {
	return ResponseStatusRepository{db: db}
}

func (r ResponseStatusRepository) GetStatus(ctx context.Context, name string) (*model.ResponseStatusModel, error) {
	responseStatus := model.ResponseStatusModel{}
	err := r.db.First(&responseStatus, "name = ?", name).Error
	return &responseStatus, err
}
