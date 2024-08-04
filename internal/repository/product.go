package repository

import (
	"context"
	"gorm.io/gorm"
	"work-project/internal/model"
)

type Product interface {
	CreateMany(ctx context.Context, products []model.Product) ([]model.Product, error)
}

type ProductDb struct {
	db *gorm.DB
}

func NewProductDb(db *gorm.DB) *ProductDb {
	return &ProductDb{db: db}
}

func (r *ProductDb) CreateMany(ctx context.Context, products []model.Product) ([]model.Product, error) {
	db := r.db.WithContext(ctx)
	err := db.Model(&model.Product{}).
		Create(&products).
		Error
	if err != nil {
		return nil, err
	}
	return products, nil
}
