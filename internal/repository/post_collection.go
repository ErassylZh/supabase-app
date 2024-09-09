package repository

import (
	"context"
	"gorm.io/gorm"
	"work-project/internal/model"
)

type PostCollection interface {
	CreateMany(ctx context.Context, posts []model.PostCollection) ([]model.PostCollection, error)
	DeleteByPostId(ctx context.Context, postId uint) error
	GetByPostId(ctx context.Context, postId uint) ([]model.PostCollection, error)
}

type PostCollectionDb struct {
	db *gorm.DB
}

func NewPostCollectionDb(db *gorm.DB) *PostCollectionDb {
	return &PostCollectionDb{db: db}
}

func (r *PostCollectionDb) CreateMany(ctx context.Context, posts []model.PostCollection) ([]model.PostCollection, error) {
	db := r.db.WithContext(ctx)
	err := db.Model(&model.PostCollection{}).
		Create(&posts).
		Error
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func (r *PostCollectionDb) DeleteByPostId(ctx context.Context, postId uint) error {
	db := r.db.WithContext(ctx)
	err := db.Model(&model.PostCollection{}).
		Where("post_id = ?", postId).
		Delete(&model.PostCollection{}).
		Error
	if err != nil {
		return err
	}
	return nil
}

func (r *PostCollectionDb) GetByPostId(ctx context.Context, postId uint) ([]model.PostCollection, error) {
	db := r.db.WithContext(ctx)
	var data []model.PostCollection
	err := db.Model(&model.PostCollection{}).
		Where("post_id = ?", postId).
		Find(&data).
		Error
	if err != nil {
		return nil, err
	}
	return data, nil
}
