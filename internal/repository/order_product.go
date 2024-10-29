package repository

import (
	"context"
	"gorm.io/gorm"
	"work-project/internal/model"
)

type OrderProduct interface {
	Create(ctx context.Context, orderProducts []model.OrderProduct) ([]model.OrderProduct, error)
}

type OrderProductDB struct {
	db *gorm.DB
}

func NewOrderProductDB(db *gorm.DB) *OrderProductDB {
	return &OrderProductDB{db: db}
}

func (r *OrderProductDB) Create(ctx context.Context, orderProducts []model.OrderProduct) ([]model.OrderProduct, error) {
	db := r.db.WithContext(ctx)
	err := db.Model(&model.OrderProduct{}).
		Create(&orderProducts).
		Error
	if err != nil {
		return orderProducts, err
	}
	return orderProducts, nil
}
