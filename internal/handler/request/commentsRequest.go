package request

import "github.com/TemaStatham/TaskService/pkg/paginate"

type CreateCommentRequest struct {
	TaskID  uint
	Comment string
}

type ShowCommentRequest struct {
	TaskID     uint
	Pagination paginate.Pagination
}
