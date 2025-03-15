package data

type CreateApprove struct {
	TaskID   uint
	UserID   uint
	StatusID uint
	Score    uint
	Approved *uint
}
