package repository

import (
	"context"
	"gorm.io/gorm"
	"work-project/internal/model"
)

type PrivacyTerms interface {
	GetAll(ctx context.Context) ([]model.PrivacyTerms, error)
}

type PrivacyTermsDB struct {
	db *gorm.DB
}

func NewPrivacyTermsDB(db *gorm.DB) *PrivacyTermsDB {
	return &PrivacyTermsDB{db: db}
}

func (r *PrivacyTermsDB) GetAll(ctx context.Context) (privacyTerms []model.PrivacyTerms, err error) {
	db := r.db.WithContext(ctx)
	q := db.Model(&model.PrivacyTerms{})
	err = q.Find(&privacyTerms).
		Error
	if err != nil {
		return privacyTerms, err
	}
	return privacyTerms, nil
}
