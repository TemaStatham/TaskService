package postgres

import (
	"context"
	"github.com/TemaStatham/TaskService/internal/model"
	"github.com/TemaStatham/TaskService/pkg/paginate"
	"gorm.io/gorm"
)

type CommentsRepository struct {
	db *gorm.DB
}

func NewCommentsRepository(db *gorm.DB) *CommentsRepository {
	return &CommentsRepository{
		db: db,
	}
}

func (c *CommentsRepository) Create(ctx context.Context, taskId, userId uint, comment string) (uint, error) {
	commentModel := &model.Comment{
		TaskID:  taskId,
		UserID:  userId,
		Comment: comment,
	}

	res := c.db.Create(comment)
	if res.Error != nil {
		return 0, res.Error
	}

	return commentModel.ID, nil
}

func (c *CommentsRepository) Show(
	ctx context.Context,
	taskId uint,
	pagination *paginate.Pagination,
) (*paginate.Pagination, error) {
	var responses []*model.TaskModel

	res := c.db.
		Where("task_id = ?", taskId).
		Scopes(paginate.Paginate(responses, pagination, c.db)).
		Find(&responses)
	if res.Error != nil {
		return &paginate.Pagination{}, res.Error
	}

	pagination.Rows = responses

	return pagination, nil
}
