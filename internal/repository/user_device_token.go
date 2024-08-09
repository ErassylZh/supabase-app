package repository

import (
	"context"
	"gorm.io/gorm"
	"work-project/internal/model"
)

type UserDeviceToken interface {
	GetByUserId(ctx context.Context, userId string) ([]model.UserDeviceToken, error)
	Create(ctx context.Context, deviceToken model.UserDeviceToken) error
	Delete(ctx context.Context, userDeviceTokenId uint) error
}

type UserDeviceTokenDB struct {
	db *gorm.DB
}

func NewUserDeviceTokenDB(db *gorm.DB) *UserDeviceTokenDB {
	return &UserDeviceTokenDB{db: db}
}

func (r *UserDeviceTokenDB) GetByUserId(ctx context.Context, userId string) (deviceTokens []model.UserDeviceToken, err error) {
	db := r.db.WithContext(ctx)
	q := db.Model(&model.UserDeviceToken{})
	err = q.Where("user_id = ?", userId).
		Find(&deviceTokens).
		Error
	if err != nil {
		return deviceTokens, err
	}
	return deviceTokens, nil
}

func (r *UserDeviceTokenDB) Create(ctx context.Context, deviceToken model.UserDeviceToken) error {
	db := r.db.WithContext(ctx)
	q := db.Model(&model.UserDeviceToken{})
	err := q.Create(&deviceToken).Error
	return err
}

func (r *UserDeviceTokenDB) Delete(ctx context.Context, userDeviceTokenId uint) error {
	db := r.db.WithContext(ctx)
	q := db.Model(&model.UserDeviceToken{})
	err := q.Where("invited_user_id = ?", userDeviceTokenId).
		Delete(&model.UserDeviceToken{}).
		Error
	if err != nil {
		return err
	}
	return nil

}
