package repository

import (
	"context"
	"gorm.io/gorm"
	"work-project/internal/model"
)

type ContestPrize interface {
	GetByContestID(ctx context.Context, contestId uint) ([]model.ContestPrize, error)
	CreateMany(ctx context.Context, contestPrizes []model.ContestPrize) ([]model.ContestPrize, error)
	Create(ctx context.Context, contestPrize model.ContestPrize) (model.ContestPrize, error)
	UpdateMany(ctx context.Context, contestPrizes []model.ContestPrize) ([]model.ContestPrize, error)
}

type ContestPrizeDB struct {
	db *gorm.DB
}

func NewContestPrizeDB(db *gorm.DB) *ContestPrizeDB {
	return &ContestPrizeDB{db: db}
}

func (r *ContestPrizeDB) GetByContestID(ctx context.Context, contestId uint) (contestPrizes []model.ContestPrize, err error) {
	db := r.db.WithContext(ctx)
	q := db.Model(&model.ContestPrize{})
	err = q.Where("contest_id = ?", contestId).
		Find(&contestPrizes).
		Error
	if err != nil {
		return contestPrizes, err
	}
	return contestPrizes, nil
}

func (r *ContestPrizeDB) CreateMany(ctx context.Context, contestPrizes []model.ContestPrize) ([]model.ContestPrize, error) {
	db := r.db.WithContext(ctx)
	q := db.Model(&model.Collection{})
	err := q.Create(&contestPrizes).
		Error
	if err != nil {
		return contestPrizes, err
	}
	return contestPrizes, nil
}

func (r *ContestPrizeDB) Create(ctx context.Context, contestPrize model.ContestPrize) (model.ContestPrize, error) {
	db := r.db.WithContext(ctx)
	q := db.Model(&model.Collection{})
	err := q.Create(&contestPrize).
		Error
	if err != nil {
		return contestPrize, err
	}
	return contestPrize, nil
}

func (r *ContestPrizeDB) UpdateMany(ctx context.Context, contestPrizes []model.ContestPrize) ([]model.ContestPrize, error) {
	db := r.db.WithContext(ctx)

	for _, contestPrize := range contestPrizes {
		if err := db.Model(&model.ContestPrize{}).
			Where("contest_id = ?", contestPrize.ContestPrizeID).
			Updates(&contestPrize).
			Error; err != nil {
			return nil, err
		}
	}
	return contestPrizes, nil
}
