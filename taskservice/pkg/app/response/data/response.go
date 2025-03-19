package data

type CreateResponse struct {
	TaskId uint `json:"task_id"`
	// userid достаю из токена
}

// todo: сделать два разных метода для отмены отклика и для подтверждения отклика
type UpdateResponse struct {
	ID uint `json:"response_id"`
}

// валдидировать что координатор - на фронте
type GetResponses struct {
	TaskId uint `json:"task_id"`
	Page   int  `json:"page,omitempty;query:page"`
	Limit  int  `json:"limit,omitempty;query:limit"`
}
