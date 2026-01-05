package repository

import (
	"context"
	"gorm.io/gorm"
	"work-project/internal/model"
)

type ProductTag interface {
	GetByID(ctx context.Context, id uint) (model.ProductTag, error)
	GetByName(ctx context.Context, productTagName string) (model.ProductTag, error)
	GetAll(ctx context.Context) ([]model.ProductTag, error)
	CreateMany(ctx context.Context, productTags []model.ProductTag) ([]model.ProductTag, error)
	UpdateMany(ctx context.Context, productTags []model.ProductTag) ([]model.ProductTag, error)
	DeleteMany(ctx context.Context, productTagIds []uint) error
}

type ProductTagDB struct {
	db *gorm.DB
}

func NewProductTagDB(db *gorm.DB) *ProductTagDB {
	return &ProductTagDB{db: db}
}

func (r *ProductTagDB) GetByID(ctx context.Context, id uint) (productTag model.ProductTag, err error) {
	db := r.db.WithContext(ctx)
	q := db.Model(&model.ProductTag{})
	err = q.Where("product_tag_id = ?", id).
		First(&productTag).
		Error
	if err != nil {
		return productTag, err
	}
	return productTag, nil
}

func (r *ProductTagDB) GetByName(ctx context.Context, productTagName string) (productTag model.ProductTag, err error) {
	db := r.db.WithContext(ctx)
	q := db.Model(&model.ProductTag{})
	err = q.Where("name = ?", productTagName).
		First(&productTag).
		Error
	if err != nil {
		return productTag, err
	}
	return productTag, nil
}

func (r *ProductTagDB) GetAll(ctx context.Context) (productTags []model.ProductTag, err error) {
	db := r.db.WithContext(ctx)
	q := db.Model(&model.ProductTag{})
	err = q.Find(&productTags).
		Error
	if err != nil {
		return productTags, err
	}
	return productTags, nil
}

func (r *ProductTagDB) CreateMany(ctx context.Context, productTags []model.ProductTag) ([]model.ProductTag, error) {
	db := r.db.WithContext(ctx)
	q := db.Model(&model.ProductTag{})
	err := q.Create(&productTags).
		Error
	if err != nil {
		return productTags, err
	}
	return productTags, nil
}

func (r *ProductTagDB) UpdateMany(ctx context.Context, productTags []model.ProductTag) ([]model.ProductTag, error) {
	db := r.db.WithContext(ctx)

	for _, productTag := range productTags {
		if err := db.Model(&model.ProductTag{}).Where("product_tag_id = ?", productTag.ProductTagID).Updates(map[string]interface{}{
			"name":       productTag.Name,
			"name_ru":    productTag.NameRu,
			"name_kz":    productTag.NameKz,
			"image_path": productTag.ImagePath,
		}).Error; err != nil {
			return nil, err
		}
	}
	return productTags, nil
}

func (r *ProductTagDB) DeleteMany(ctx context.Context, product_tagIds []uint) error {
	db := r.db.WithContext(ctx)
	q := db.Model(&model.ProductTag{})
	err := q.Where("product_tag_id in (?)", product_tagIds).
		Delete(&model.ProductTag{}).
		Error
	if err != nil {
		return err
	}
	return nil
}
