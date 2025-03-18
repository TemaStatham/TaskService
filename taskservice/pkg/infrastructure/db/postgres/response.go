package postgres

import (
	"context"
	"github.com/TemaStatham/TaskService/taskservice/pkg/app/paginate"
	"github.com/TemaStatham/TaskService/taskservice/pkg/app/response/model"
	paginate2 "github.com/TemaStatham/TaskService/taskservice/pkg/infrastructure/lib/paginate"
	"gorm.io/gorm"
)

type ResponseRepository struct {
	db *gorm.DB
}

func NewResponsePostgresRepository(db *gorm.DB) *ResponseRepository {
	return &ResponseRepository{
		db: db,
	}
}

func (r *ResponseRepository) Create(
	ctx context.Context,
	response model.ResponseModel,
) (uint, error) {
	res := r.db.WithContext(ctx).Create(&response)
	if res.Error != nil {
		return 0, res.Error
	}

	return response.ID, res.Error
}

func (r *ResponseRepository) Show(
	ctx context.Context,
	taskId uint,
	pagination *paginate2.Pagination,
) (*paginate2.Pagination, error) {
	var responses []*model.ResponseModel

	res := r.db.
		WithContext(ctx).
		Where("task_id = ?", taskId).
		Scopes(paginate.Paginate(responses, pagination, r.db)).
		Find(&responses)
	if res.Error != nil {
		return &paginate2.Pagination{}, res.Error
	}

	pagination.Rows = responses

	return pagination, nil
}

func (r *ResponseRepository) Update(ctx context.Context, id uint, status uint) error {
	var response model.ResponseModel

	res := r.db.WithContext(ctx).First(&response, "id = ?", id)
	if res.Error != nil {
		return res.Error
	}

	response.StatusID = status

	res = r.db.Model(&response).Update("status_id", status)
	if res.Error != nil {
		return res.Error
	}

	return nil
}
