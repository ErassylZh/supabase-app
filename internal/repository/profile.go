package repository

import (
	"context"
	"gorm.io/gorm"
	"work-project/internal/model"
)

type Profile interface {
	GetByID(ctx context.Context, profileId string) (model.Profile, error)
	DeleteByID(ctx context.Context, profileId string) error
	Update(ctx context.Context, profile model.Profile) error
}

type ProfileDB struct {
	db *gorm.DB
}

func NewProfileDB(db *gorm.DB) *ProfileDB {
	return &ProfileDB{db: db}
}

func (r *ProfileDB) GetByID(ctx context.Context, profileId string) (profile model.Profile, err error) {
	db := r.db.WithContext(ctx)
	q := db.Model(&model.Profile{})
	err = q.Where("id = ?", profileId).
		First(&profile).
		Error
	if err != nil {
		return profile, err
	}
	return profile, nil
}

func (r *ProfileDB) DeleteByID(ctx context.Context, userId string) (err error) {
	db := r.db.WithContext(ctx)
	q := db.Model(&model.Profile{})
	err = q.Where("id = ?", userId).
		Delete(&model.Profile{}).
		Error
	if err != nil {
		return err
	}
	return nil
}

func (r *ProfileDB) Update(ctx context.Context, profile model.Profile) (err error) {
	db := r.db.WithContext(ctx)
	q := db.Model(&model.Profile{})
	err = q.Where("id = ?", profile.ID).
		Save(&profile).
		Error
	if err != nil {
		return err
	}
	return nil
}
