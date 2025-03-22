package postgres

import (
	"context"
	"github.com/TemaStatham/TaskService/taskservice/pkg/app/comment/data"
	"github.com/TemaStatham/TaskService/taskservice/pkg/app/comment/model"
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

func (c *CommentsRepository) Show(ctx context.Context, taskId uint, page int, limit int) (*paginate2.Pagination, error) {
	var responses []*model.CommentModel
	query := c.db.WithContext(ctx).Where("task_id = ?", taskId)

	var total int64
	if err := query.Model(&model.CommentModel{}).Count(&total).Error; err != nil {
		return nil, err
	}

	offset := (page - 1) * limit
	if err := query.Offset(offset).Limit(limit).Find(&responses).Error; err != nil {
		return nil, err
	}

	pagination := paginate2.Pagination{limit, page, total, responses}
	return &pagination, nil
}
