package data

import "github.com/TemaStatham/TaskService/client/pkg/app/paginate"

type CreateComment struct {
	TaskID  uint
	Comment string
}

type ShowComment struct {
	TaskID     uint
	Pagination paginate.Pagination
}
