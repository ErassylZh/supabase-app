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
	GetByContestIDAndNumber(ctx context.Context, contestId uint, number int) (model.ContestPrize, error)
	GetAll(ctx context.Context) ([]model.ContestPrize, error)
	GetByID(ctx context.Context, id uint) (model.ContestPrize, error)
	Update(ctx context.Context, prize model.ContestPrize) (model.ContestPrize, error)
	Delete(ctx context.Context, id uint) error
}

type ContestPrizeDB struct {
	db *gorm.DB
}

func NewContestPrizeDB(db *gorm.DB) *ContestPrizeDB {
	return &ContestPrizeDB{db: db}
}

func (r *ContestPrizeDB) GetByContestIDAndNumber(ctx context.Context, contestId uint, number int) (contestPrize model.ContestPrize, err error) {
	db := r.db.WithContext(ctx)
	q := db.Model(&model.ContestPrize{})
	err = q.Where("contest_id = ?", contestId).
		Where("number = ?", number).
		First(&contestPrize).
		Error
	if err != nil {
		return contestPrize, err
	}
	return contestPrize, nil
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
	err := r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		for _, contestPrize := range contestPrizes {
			if err := tx.Model(&model.ContestBook{}).Where("contest_prize_id = ?", contestPrize.ContestPrizeID).Updates(&contestPrize).Error; err != nil {
				return err // Если ошибка — транзакция отменяется
			}
		}
		return nil
	})

	if err != nil {
		return nil, err
	}
	return contestPrizes, nil
}

func (r *ContestPrizeDB) GetAll(ctx context.Context) (contestPrizes []model.ContestPrize, err error) {
	db := r.db.WithContext(ctx)
	q := db.Model(&model.ContestPrize{})
	err = q.Find(&contestPrizes).
		Error
	if err != nil {
		return contestPrizes, err
	}
	return contestPrizes, nil
}

func (r *ContestPrizeDB) GetByID(ctx context.Context, id uint) (contestPrizes model.ContestPrize, err error) {
	db := r.db.WithContext(ctx)
	q := db.Model(&model.ContestPrize{})
	err = q.Where("contest_prize_id = ?", id).
		First(&contestPrizes).
		Error
	if err != nil {
		return contestPrizes, err
	}
	return contestPrizes, nil
}

func (r *ContestPrizeDB) Update(ctx context.Context, prize model.ContestPrize) (model.ContestPrize, error) {
	db := r.db.WithContext(ctx)
	q := db.Model(&model.ContestPrize{})
	err := q.Where("contest_prize_id = ?", prize.ContestPrizeID).
		Save(&prize).
		Error
	if err != nil {
		return prize, err
	}
	return prize, nil
}

func (r *ContestPrizeDB) Delete(ctx context.Context, id uint) error {
	db := r.db.WithContext(ctx)
	q := db.Model(&model.ContestPrize{})
	err := q.Where("contest_prize_id = ?", id).
		Delete(&model.ContestPrize{}).
		Error
	if err != nil {
		return err
	}
	return nil
}
