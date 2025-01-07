package repository

import (
	"context"
	"gorm.io/gorm"
	"time"
	"work-project/internal/model"
)

type Contest interface {
	GetActive(ctx context.Context) (model.Contest, error)
	Create(ctx context.Context, contest model.Contest) (model.Contest, error)
	Update(ctx context.Context, contest model.Contest) (model.Contest, error)
}

type ContestDB struct {
	db *gorm.DB
}

func NewContestDB(db *gorm.DB) *ContestDB {
	return &ContestDB{db: db}
}

func (r *ContestDB) GetActive(ctx context.Context) (contest model.Contest, err error) {
	db := r.db.WithContext(ctx)
	q := db.Model(&model.Contest{})
	err = q.Where("is_active and start_time > ? and end_time < ?", time.Now()).
		First(&contest).
		Error
	if err != nil {
		return contest, err
	}
	return contest, nil
}

func (r *ContestDB) Create(ctx context.Context, contest model.Contest) (model.Contest, error) {
	db := r.db.WithContext(ctx)
	q := db.Model(&model.Collection{})
	err := q.Create(&contest).
		Error
	if err != nil {
		return contest, err
	}
	return contest, nil
}

func (r *ContestDB) Update(ctx context.Context, contest model.Contest) (model.Contest, error) {
	db := r.db.WithContext(ctx)
	q := db.Model(&model.Collection{})
	err := q.Where("contest_id = ?", contest.ContestID).
		Save(&contest).
		Error
	if err != nil {
		return contest, err
	}
	return contest, nil
}
