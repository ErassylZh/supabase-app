package repository

import (
	"context"
	"gorm.io/gorm"
	"work-project/internal/model"
)

type ProductProductTag interface {
	CreateMany(ctx context.Context, products []model.ProductProductTag) ([]model.ProductProductTag, error)
	DeleteByProductId(ctx context.Context, productId uint) error
	GetByProductId(ctx context.Context, productId uint) ([]model.ProductProductTag, error)
}

type ProductProductTagDB struct {
	db *gorm.DB
}

func NewProductProductTagDB(db *gorm.DB) *ProductProductTagDB {
	return &ProductProductTagDB{db: db}
}

func (r *ProductProductTagDB) CreateMany(ctx context.Context, products []model.ProductProductTag) ([]model.ProductProductTag, error) {
	db := r.db.WithContext(ctx)
	err := db.Model(&model.ProductProductTag{}).
		Create(&products).
		Error
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (r *ProductProductTagDB) DeleteByProductId(ctx context.Context, productId uint) error {
	db := r.db.WithContext(ctx)
	err := db.Model(&model.ProductProductTag{}).
		Where("product_id = ?", productId).
		Delete(&model.ProductProductTag{}).
		Error
	if err != nil {
		return err
	}
	return nil
}

func (r *ProductProductTagDB) GetByProductId(ctx context.Context, productId uint) ([]model.ProductProductTag, error) {
	db := r.db.WithContext(ctx)
	var data []model.ProductProductTag
	err := db.Model(&model.ProductProductTag{}).
		Where("product_id = ?", productId).
		Find(&data).
		Error
	if err != nil {
		return nil, err
	}
	return data, nil
}
