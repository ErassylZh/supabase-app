package repository

import (
	"context"
	"gorm.io/gorm"
	"work-project/internal/model"
)

type PostHashtag interface {
	CreateMany(ctx context.Context, posts []model.PostHashtag) ([]model.PostHashtag, error)
	DeleteByPostId(ctx context.Context, postId uint) error
	GetByPostId(ctx context.Context, postId uint) ([]model.PostHashtag, error)
}

type PostHashtagDb struct {
	db *gorm.DB
}

func NewPostHashtagDb(db *gorm.DB) *PostHashtagDb {
	return &PostHashtagDb{db: db}
}

func (r *PostHashtagDb) CreateMany(ctx context.Context, posts []model.PostHashtag) ([]model.PostHashtag, error) {
	db := r.db.WithContext(ctx)
	err := db.Model(&model.PostHashtag{}).
		Create(&posts).
		Error
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func (r *PostHashtagDb) DeleteByPostId(ctx context.Context, postId uint) error {
	db := r.db.WithContext(ctx)
	err := db.Model(&model.PostHashtag{}).
		Where("post_id = ?", postId).
		Delete(&model.PostHashtag{}).
		Error
	if err != nil {
		return err
	}
	return nil
}

func (r *PostHashtagDb) GetByPostId(ctx context.Context, postId uint) ([]model.PostHashtag, error) {
	db := r.db.WithContext(ctx)
	var data []model.PostHashtag
	err := db.Model(&model.PostHashtag{}).
		Where("post_id = ?", postId).
		Find(&data).
		Error
	if err != nil {
		return nil, err
	}
	return data, nil
}