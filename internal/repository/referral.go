package repository

import (
	"context"
	"gorm.io/gorm"
	"work-project/internal/model"
)

type Referral interface {
	GetByUserId(ctx context.Context, userId string) (model.Referral, error)
	GetByInvitedUserId(ctx context.Context, userId string) (model.Referral, error)
	Create(ctx context.Context, referral model.Referral) error
}

type ReferralDB struct {
	db *gorm.DB
}

func NewReferralDB(db *gorm.DB) *ReferralDB {
	return &ReferralDB{db: db}
}

func (r *ReferralDB) GetByUserId(ctx context.Context, userId string) (referral model.Referral, err error) {
	db := r.db.WithContext(ctx)
	q := db.Model(&model.Referral{})
	err = q.Where("user_id = ?", userId).
		First(&referral).
		Error
	if err != nil {
		return referral, err
	}
	return referral, nil
}

func (r *ReferralDB) Create(ctx context.Context, referral model.Referral) error {
	db := r.db.WithContext(ctx)
	q := db.Model(&model.Referral{})
	err := q.Create(&referral).Error
	return err
}

func (r *ReferralDB) GetByInvitedUserId(ctx context.Context, userId string) (referral model.Referral, err error) {
	db := r.db.WithContext(ctx)
	q := db.Model(&model.Referral{})
	err = q.Where("invited_user_id = ?", userId).
		First(&referral).
		Error
	if err != nil {
		return referral, err
	}
	return referral, nil

}
