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
	GetAllListing(ctx context.Context, productTagIds []uint) ([]model.Product, error)
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
		Preload("ProductTags").
		Find(&products).
		Error
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (r *ProductDB) UpdateMany(ctx context.Context, products []model.Product) ([]model.Product, error) {
	err := r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		for _, product := range products {
			if err := tx.Model(&model.Product{}).Where("product_id = ?", product.ProductID).Updates(&product).Error; err != nil {
				return err // Если ошибка — транзакция отменяется
			}
		}
		return nil
	})

	if err != nil {
		return nil, err
	}
	return products, nil
}
func (r *ProductDB) GetAllListing(ctx context.Context, productTagIds []uint) (products []model.Product, err error) {
	db := r.db.WithContext(ctx)

	query := db.Model(&model.Product{})
	if len(productTagIds) > 0 {
		query = query.Joins("JOIN public.product_product_tag ON public.product_product_tag.product_id = public.product.product_id").
			Joins("JOIN public.product_tag ON public.product_tag.product_tag_id = public.product_product_tag.product_tag_id").
			Where("public.product_product_tag.product_tag_id IN (?)", productTagIds)
	}
	err = query.
		Where("status = ?", model.PRODUCT_STATUS_PUBLISH).
		Preload("Images").
		Preload("ProductTags").
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
