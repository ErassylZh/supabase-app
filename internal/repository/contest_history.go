package repository

import (
	"context"
	"gorm.io/gorm"
	"work-project/internal/model"
)

type ContestHistory interface {
	GetByContestID(ctx context.Context, contestId uint) ([]model.ContestHistory, error)
	CreateMany(ctx context.Context, contestHistories []model.ContestHistory) ([]model.ContestHistory, error)
	Create(ctx context.Context, contestHistory model.ContestHistory) (model.ContestHistory, error)
	UpdateMany(ctx context.Context, contestHistories []model.ContestHistory) ([]model.ContestHistory, error)
	GetByContestBookAndUserID(ctx context.Context, contestBookId uint, userId string) (model.ContestHistory, error)
}

type ContestHistoryDB struct {
	db *gorm.DB
}

func NewContestHistoryDB(db *gorm.DB) *ContestHistoryDB {
	return &ContestHistoryDB{db: db}
}

func (r *ContestHistoryDB) GetByContestBookAndUserID(ctx context.Context, contestBookId uint, userId string) (contestHistories model.ContestHistory, err error) {
	db := r.db.WithContext(ctx)
	q := db.Model(&model.ContestHistory{})
	err = q.Where("contest_book_id = ?", contestBookId).
		Where("user_id = ?", userId).
		First(&contestHistories).
		Error
	if err != nil {
		return contestHistories, err
	}
	return contestHistories, nil
}

func (r *ContestHistoryDB) GetByContestID(ctx context.Context, contestId uint) (contestHistories []model.ContestHistory, err error) {
	db := r.db.WithContext(ctx)
	q := db.Model(&model.ContestHistory{})
	err = q.Where("contest_id = ?", contestId).
		Find(&contestHistories).
		Error
	if err != nil {
		return contestHistories, err
	}
	return contestHistories, nil
}

func (r *ContestHistoryDB) CreateMany(ctx context.Context, contestHistories []model.ContestHistory) ([]model.ContestHistory, error) {
	db := r.db.WithContext(ctx)
	q := db.Model(&model.Collection{})
	err := q.Create(&contestHistories).
		Error
	if err != nil {
		return contestHistories, err
	}
	return contestHistories, nil
}

func (r *ContestHistoryDB) Create(ctx context.Context, contestHistory model.ContestHistory) (model.ContestHistory, error) {
	db := r.db.WithContext(ctx)
	q := db.Model(&model.Collection{})
	err := q.Create(&contestHistory).
		Error
	if err != nil {
		return contestHistory, err
	}
	return contestHistory, nil
}

func (r *ContestHistoryDB) UpdateMany(ctx context.Context, contestHistories []model.ContestHistory) ([]model.ContestHistory, error) {
	db := r.db.WithContext(ctx)

	for _, contestHistory := range contestHistories {
		if err := db.Model(&model.ContestHistory{}).
			Where("contest_id = ?", contestHistory.ContestHistoryID).
			Updates(&contestHistory).
			Error; err != nil {
			return nil, err
		}
	}
	return contestHistories, nil
}
