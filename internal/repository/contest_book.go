package repository

import (
	"context"
	"gorm.io/gorm"
	"work-project/internal/model"
)

type ContestBook interface {
	GetByContestID(ctx context.Context, contestId uint) ([]model.ContestBook, error)
	CreateMany(ctx context.Context, contestBooks []model.ContestBook) ([]model.ContestBook, error)
	UpdateMany(ctx context.Context, contestBooks []model.ContestBook) ([]model.ContestBook, error)
}

type ContestBookDB struct {
	db *gorm.DB
}

func NewContestBookDB(db *gorm.DB) *ContestBookDB {
	return &ContestBookDB{db: db}
}

func (r *ContestBookDB) GetByContestID(ctx context.Context, contestId uint) (contestBooks []model.ContestBook, err error) {
	db := r.db.WithContext(ctx)
	q := db.Model(&model.ContestBook{})
	err = q.Where("contest_id = ?", contestId).
		Find(&contestBooks).
		Error
	if err != nil {
		return contestBooks, err
	}
	return contestBooks, nil
}

func (r *ContestBookDB) CreateMany(ctx context.Context, contestBooks []model.ContestBook) ([]model.ContestBook, error) {
	db := r.db.WithContext(ctx)
	q := db.Model(&model.Collection{})
	err := q.Create(&contestBooks).
		Error
	if err != nil {
		return contestBooks, err
	}
	return contestBooks, nil
}

func (r *ContestBookDB) UpdateMany(ctx context.Context, contestBooks []model.ContestBook) ([]model.ContestBook, error) {
	db := r.db.WithContext(ctx)

	for _, contestBook := range contestBooks {
		if err := db.Model(&model.ContestBook{}).
			Where("contest_id = ?", contestBook.ContestBookID).
			Updates(&contestBook).
			Error; err != nil {
			return nil, err
		}
	}
	return contestBooks, nil
}
