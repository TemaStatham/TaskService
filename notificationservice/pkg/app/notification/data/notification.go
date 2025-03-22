package data

type CreateNotification struct {
	UserID  uint64 `json:"user_id"`
	Message string `json:"message"`
}

type GetNotifications struct {
	UserID uint64 `json:"user_id"`
}
