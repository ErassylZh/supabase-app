package repository

import (
	"context"
	"gorm.io/gorm"
	"work-project/internal/model"
)

type Order interface {
	Create(ctx context.Context, order model.Order) (model.Order, error)
}

type OrderDB struct {
	db *gorm.DB
}

func NewOrderDB(db *gorm.DB) *OrderDB {
	return &OrderDB{db: db}
}

func (r *OrderDB) Create(ctx context.Context, order model.Order) (model.Order, error) {
	db := r.db.WithContext(ctx)
	err := db.Model(&model.Order{}).
		Create(&order).
		Error
	if err != nil {
		return order, err
	}
	return order, nil
}
