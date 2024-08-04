package repository

import (
	"context"
	"gorm.io/gorm"
	"work-project/internal/model"
)

type PushNotification interface {
	GetByID(ctx context.Context, profileId string) (model.PushNotification, error)
	Create(ctx context.Context, notification model.PushNotification) error
	GetByToken(ctx context.Context, token *string) ([]model.PushNotification, error)
}

type PushNotificationDB struct {
	db *gorm.DB
}

func NewPushNotificationDB(db *gorm.DB) *PushNotificationDB {
	return &PushNotificationDB{db: db}
}

func (r *PushNotificationDB) GetByID(ctx context.Context, notificationId string) (notification model.PushNotification, err error) {
	db := r.db.WithContext(ctx)
	q := db.Model(&model.PushNotification{})
	err = q.Where("push_notification_id	 = ?", notificationId).
		First(&notificationId).
		Error
	if err != nil {
		return notification, err
	}
	return notification, nil
}

func (r *PushNotificationDB) Create(ctx context.Context, notification model.PushNotification) error {
	db := r.db.WithContext(ctx)
	q := db.Model(&model.PushNotification{})
	err := q.Create(&notification).
		Error
	if err != nil {
		return err
	}
	return nil
}

func (r *PushNotificationDB) GetByToken(ctx context.Context, token *string) (notifications []model.PushNotification, err error) {
	db := r.db.WithContext(ctx)
	q := db.Model(&model.PushNotification{})
	err = q.Where("token = ?", token).
		Or("token is null and condition is null").
		Find(&notifications).
		Error
	if err != nil {
		return notifications, err
	}
	return notifications, nil
}
