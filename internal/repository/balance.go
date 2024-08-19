package repository

import (
	"context"
	"gorm.io/gorm"
	"work-project/internal/model"
)

type Balance interface {
	GetByUserID(ctx context.Context, userId string) (model.Balance, error)
	Update(ctx context.Context, balance model.Balance) error
}

type BalanceDB struct {
	db *gorm.DB
}

func NewBalanceDB(db *gorm.DB) *BalanceDB {
	return &BalanceDB{db: db}
}

func (r *BalanceDB) GetByUserID(ctx context.Context, userId string) (balance model.Balance, err error) {
	db := r.db.WithContext(ctx)
	q := db.Model(&model.Balance{})
	err = q.Where("user_id = ?", userId).
		First(&balance).
		Error
	if err != nil {
		return balance, err
	}
	return balance, nil
}
func (r *BalanceDB) Update(ctx context.Context, balance model.Balance) error {
	db := r.db.WithContext(ctx)
	q := db.Model(&model.Balance{})
	err := q.Save(&balance).
		Error
	if err != nil {
		return err
	}
	return nil
}