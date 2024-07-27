package repository

import (
	"context"
	"gorm.io/gorm"
	"work-project/internal/model"
)

type Transaction interface {
	GetAllByUserID(ctx context.Context, userId string) ([]model.Transaction, error)
	Create(ctx context.Context, transaction model.Transaction) error
}

type TransactionDB struct {
	db *gorm.DB
}

func NewTransactionDB(db *gorm.DB) *TransactionDB {
	return &TransactionDB{db: db}
}

func (r *TransactionDB) GetAllByUserID(ctx context.Context, userId string) (transactions []model.Transaction, err error) {
	db := r.db.WithContext(ctx)
	q := db.Model(&model.Transaction{})
	err = q.Where("user_id = ?", userId).
		Order("created_at").
		Find(&transactions).
		Error
	if err != nil {
		return transactions, err
	}
	return transactions, nil
}
func (r *TransactionDB) Create(ctx context.Context, transaction model.Transaction) error {
	db := r.db.WithContext(ctx)
	q := db.Model(&model.Transaction{})
	err := q.Create(&transaction).
		Error
	if err != nil {
		return err
	}
	return nil
}
