package repository

import (
	"context"
	"gorm.io/gorm"
	"work-project/internal/model"
)

type Post interface {
	CreateMany(ctx context.Context, posts []model.Post) ([]model.Post, error)
	GetAll(ctx context.Context) ([]model.Post, error)
	UpdateMany(ctx context.Context, posts []model.Post) ([]model.Post, error)
	GetAllForListing(ctx context.Context) ([]model.Post, error)
}

type PostDb struct {
	db *gorm.DB
}

func NewPostDb(db *gorm.DB) *PostDb {
	return &PostDb{db: db}
}

func (r *PostDb) CreateMany(ctx context.Context, posts []model.Post) ([]model.Post, error) {
	db := r.db.WithContext(ctx)
	err := db.Model(&model.Post{}).
		Create(&posts).
		Error
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func (r *PostDb) GetAll(ctx context.Context) (posts []model.Post, err error) {
	db := r.db.WithContext(ctx)
	err = db.Model(&model.Post{}).
		Find(&posts).
		Error
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func (r *PostDb) UpdateMany(ctx context.Context, posts []model.Post) ([]model.Post, error) {
	db := r.db.WithContext(ctx)

	for _, post := range posts {
		if err := db.Model(&model.Post{}).Where("post_id = ?", post.PostID).Updates(&post).Error; err != nil {
			return nil, err
		}
	}
	return posts, nil
}

func (r *PostDb) GetAllForListing(ctx context.Context) (posts []model.Post, err error) {
	db := r.db.WithContext(ctx)
	err = db.Model(&model.Post{}).
		Where("status = ?", model.PRODUCT_STATUS_PUBLISH).
		Preload("Images").
		Preload("Hashtags").
		Find(&posts).
		Error
	if err != nil {
		return nil, err
	}

	return posts, nil
}
