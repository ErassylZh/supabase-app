package repository

import (
	"context"
	"gorm.io/gorm"
	"work-project/internal/model"
)

type Product interface {
	CreateMany(ctx context.Context, products []model.Product) ([]model.Product, error)
	GetAll(ctx context.Context) ([]model.Product, error)
	UpdateMany(ctx context.Context, products []model.Product) ([]model.Product, error)
	GetAllListing(ctx context.Context) ([]model.Product, error)
	GetById(ctx context.Context, id uint) (model.Product, error)
}

type ProductDB struct {
	db *gorm.DB
}

func NewProductDB(db *gorm.DB) *ProductDB {
	return &ProductDB{db: db}
}

func (r *ProductDB) CreateMany(ctx context.Context, products []model.Product) ([]model.Product, error) {
	db := r.db.WithContext(ctx)
	err := db.Model(&model.Product{}).
		Create(&products).
		Error
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (r *ProductDB) GetAll(ctx context.Context) (products []model.Product, err error) {
	db := r.db.WithContext(ctx)
	err = db.Model(&model.Product{}).
		Find(&products).
		Error
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (r *ProductDB) UpdateMany(ctx context.Context, products []model.Product) ([]model.Product, error) {
	db := r.db.WithContext(ctx)

	for _, product := range products {
		if err := db.Model(&model.Product{}).Where("product_id = ?", product.ProductID).Updates(&product).Error; err != nil {
			return nil, err
		}
	}
	return products, nil
}
func (r *ProductDB) GetAllListing(ctx context.Context) (products []model.Product, err error) {
	db := r.db.WithContext(ctx)
	err = db.Model(&model.Product{}).
		Where("status = ?", model.PRODUCT_STATUS_PUBLISH).
		Preload("Images").
		Find(&products).
		Error
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (r *ProductDB) GetById(ctx context.Context, id uint) (product model.Product, err error) {
	db := r.db.WithContext(ctx)
	err = db.Model(&model.Product{}).
		Where("status = ?", model.PRODUCT_STATUS_PUBLISH).
		Where("product_id = ? ", id).
		First(&product).
		Error
	if err != nil {
		return model.Product{}, err
	}
	return product, nil
}
