package repository

import (
	"context"
	"gorm.io/gorm"
	"work-project/internal/model"
)

type ReferralCode interface {
	GetByUserId(ctx context.Context, userId string) (model.ReferralCode, error)
	GetByReferralCode(ctx context.Context, referralCode string) (model.ReferralCode, error)
}

type ReferralCodeDB struct {
	db *gorm.DB
}

func NewReferralCodeDB(db *gorm.DB) *ReferralCodeDB {
	return &ReferralCodeDB{db: db}
}

func (r *ReferralCodeDB) GetByUserId(ctx context.Context, userId string) (referralCode model.ReferralCode, err error) {
	db := r.db.WithContext(ctx)
	q := db.Model(&model.ReferralCode{})
	err = q.Where("user_id = ?", userId).
		First(&referralCode).
		Error
	if err != nil {
		return referralCode, err
	}
	return referralCode, nil
}

func (r *ReferralCodeDB) GetByReferralCode(ctx context.Context, code string) (referralCode model.ReferralCode, err error) {
	db := r.db.WithContext(ctx)
	q := db.Model(&model.ReferralCode{})
	err = q.Where("referral_code = ?", code).
		First(&referralCode).
		Error
	if err != nil {
		return referralCode, err
	}
	return referralCode, nil
}
