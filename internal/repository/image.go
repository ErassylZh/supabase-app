package repository

import (
	"context"
	"gorm.io/gorm"
	"work-project/internal/model"
)

type Image interface {
	CreateMany(ctx context.Context, images []model.Image) ([]model.Image, error)
	GetAllByProductId(ctx context.Context, productId uint) ([]model.Image, error)
	GetAllByPostId(ctx context.Context, postId uint) ([]model.Image, error)
	DeleteByPostId(ctx context.Context, postId uint) error
}

type ImageDb struct {
	db *gorm.DB
}

func NewImageDb(db *gorm.DB) *ImageDb {
	return &ImageDb{db: db}
}

func (r *ImageDb) CreateMany(ctx context.Context, images []model.Image) ([]model.Image, error) {
	db := r.db.WithContext(ctx)
	err := db.Model(&model.Image{}).
		Create(&images).
		Error
	if err != nil {
		return nil, err
	}
	return images, nil
}

func (r *ImageDb) GetAllByProductId(ctx context.Context, productId uint) (images []model.Image, err error) {
	db := r.db.WithContext(ctx)
	err = db.Model(&model.Image{}).
		Where("product_id = ?", productId).
		Find(&images).
		Error
	if err != nil {
		return nil, err
	}
	return images, nil
}

func (r *ImageDb) GetAllByPostId(ctx context.Context, postId uint) (images []model.Image, err error) {
	db := r.db.WithContext(ctx)
	err = db.Model(&model.Image{}).
		Where("post_id = ?", postId).
		Find(&images).
		Error
	if err != nil {
		return nil, err
	}
	return images, nil
}

func (r *ImageDb) DeleteByPostId(ctx context.Context, postId uint) error {
	db := r.db.WithContext(ctx)
	err := db.Model(&model.Image{}).
		Where("post_id = ?", postId).
		Delete(&model.Image{}).
		Error
	if err != nil {
		return err
	}
	return nil
}
