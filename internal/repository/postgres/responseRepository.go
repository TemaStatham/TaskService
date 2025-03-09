package postgres

import (
	"context"
	"github.com/TemaStatham/TaskService/internal/model"
	"github.com/TemaStatham/TaskService/pkg/paginate"
	"gorm.io/gorm"
)

type ResponseRepository struct {
	db *gorm.DB
}

func NewResponseRepository(db *gorm.DB) *ResponseRepository {
	return &ResponseRepository{
		db: db,
	}
}

func (r *ResponseRepository) Get(ctx context.Context, id uint) (*model.Response, error) {
	var response model.Response

	res := r.db.First(&response, "id = ?", id)
	if res.Error != nil {
		return nil, res.Error
	}

	return &response, nil
}

func (r *ResponseRepository) Create(ctx context.Context, taskId, userId uint, status string) (uint, error) {
	response := &model.Response{
		TaskID: taskId,
		UserID: userId,
		Status: status,
	}

	res := r.db.Create(response)
	if res.Error != nil {
		return 0, res.Error
	}

	return response.ID, res.Error
}

func (r *ResponseRepository) Show(ctx context.Context, taskId uint, pagination *paginate.Pagination) (*paginate.Pagination, error) {
	var responses []*model.Task

	res := r.db.
		Where("task_id = ?", taskId).
		Scopes(paginate.Paginate(responses, pagination, r.db)).
		Find(&responses)
	if res.Error != nil {
		return &paginate.Pagination{}, res.Error
	}

	pagination.Rows = responses

	return pagination, nil
}

func (r *ResponseRepository) Update(ctx context.Context, id uint, status string) error {
	var response model.Response

	res := r.db.First(&response, "id = ?", id)
	if res.Error != nil {
		return res.Error
	}

	response.Status = status

	res = r.db.Model(&response).Update("status", status)
	if res.Error != nil {
		return res.Error
	}

	return nil
}
