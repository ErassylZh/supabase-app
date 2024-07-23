package repository

import (
	"context"
	"gorm.io/gorm"
	"work-project/internal/model"
)

type PushNotification interface {
	GetByID(ctx context.Context, profileId string) (model.PushNotification, error)
	Create(ctx context.Context, notification model.PushNotification) error
}

type PushNotificationDB struct {
	db *gorm.DB
}

func NewPushNotificationDB(db *gorm.DB) *PushNotificationDB {
	return &PushNotificationDB{db: db}
}

func (r *PushNotificationDB) GetByID(ctx context.Context, profileId string) (profile model.PushNotification, err error) {
	db := r.db.WithContext(ctx)
	q := db.Model(&model.PushNotification{})
	err = q.Where("id = ?", profileId).
		First(&profileId).
		Error
	if err != nil {
		return profile, err
	}
	return profile, nil
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
