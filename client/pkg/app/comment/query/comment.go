package query

import (
	"context"
	"github.com/TemaStatham/TaskService/client/pkg/app/comment/data"
	"github.com/TemaStatham/TaskService/client/pkg/app/comment/model"
	"github.com/TemaStatham/TaskService/client/pkg/app/paginate"
)

type CommentQueryInterface interface {
	Show(
		ctx context.Context,
		comment data.ShowComment,
	) (*paginate.Pagination, error)
}

type CommentQuery struct {
	commentRepo model.CommentReadRepositoryInterface
}

func NewCommentQuery(commentRepo model.CommentReadRepositoryInterface) *CommentQuery {
	return &CommentQuery{
		commentRepo: commentRepo,
	}
}

func (c *CommentQuery) Show(ctx context.Context, comment data.ShowComment) (*paginate.Pagination, error) {
	return nil, nil
}
