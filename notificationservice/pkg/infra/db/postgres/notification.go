package postgres

import (
	"context"

	"github.com/TemaStatham/TaskService/notificationservice/pkg/app/notification/data"
	"github.com/TemaStatham/TaskService/notificationservice/pkg/app/notification/model"
	"gorm.io/gorm"
)

type NotificationPostgres struct {
	db *gorm.DB
}

func NewNotificationPostgres(db *gorm.DB) *NotificationPostgres {
	return &NotificationPostgres{
		db: db,
	}
}

func (r *NotificationPostgres) GetMessages(ctx context.Context, id uint, page int, limit int) ([]model.Notification, error) {
	var notifications []model.Notification

	offset := (page - 1) * limit

	if err := r.db.WithContext(ctx).
		Where("user_id = ?", id).
		Order("created_at DESC").
		Limit(limit).
		Offset(offset).
		Find(&notifications).Error; err != nil {
		return nil, err
	}

	return notifications, nil
}

func (r *NotificationPostgres) CreateMessage(ctx context.Context, notification data.CreateNotification) error {
	newNotification := model.Notification{
		UserID:  notification.UserID,
		Message: notification.Message,
		IsRead:  false,
	}

	return r.db.WithContext(ctx).Create(&newNotification).Error
}

func (r *NotificationPostgres) SetIsRead(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).
		Model(&model.Notification{}).
		Where("id = ?", id).
		Update("is_read", true).Error
}
