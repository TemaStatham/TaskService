package postgres

import (
	"context"
	"github.com/TemaStatham/TaskService/taskservice/pkg/app/comment/data"
	"github.com/TemaStatham/TaskService/taskservice/pkg/app/comment/model"
	"github.com/TemaStatham/TaskService/taskservice/pkg/app/paginate"
	paginate2 "github.com/TemaStatham/TaskService/taskservice/pkg/infrastructure/lib/paginate"
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

func (c *CommentsRepository) Create(ctx context.Context, comment data.CreateComment, userID uint) (uint, error) {
	commentModel := &model.CommentModel{
		TaskID:  comment.TaskID,
		UserID:  userID,
		Comment: comment.Comment,
	}

	res := c.db.WithContext(ctx).Create(&commentModel)
	if res.Error != nil {
		return 0, res.Error
	}

	return commentModel.ID, nil
}

func (c *CommentsRepository) Show(
	ctx context.Context,
	taskId uint,
	pagination *paginate2.Pagination,
) (*paginate2.Pagination, error) {
	var responses []*model.CommentModel

	res := c.db.
		WithContext(ctx).
		Where("task_id = ?", taskId).
		Scopes(paginate.Paginate(responses, pagination, c.db)).
		Find(&responses)
	if res.Error != nil {
		return &paginate2.Pagination{}, res.Error
	}

	pagination.Rows = responses

	return pagination, nil
}
