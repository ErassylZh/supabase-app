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
	Create(ctx context.Context, collection model.PostHashtag) (model.PostHashtag, error)
	DeleteByPostAndHashtagId(ctx context.Context, postId, hashtagId uint) error
}

type PostHashtagDB struct {
	db *gorm.DB
}

func NewPostHashtagDB(db *gorm.DB) *PostHashtagDB {
	return &PostHashtagDB{db: db}
}

func (r *PostHashtagDB) CreateMany(ctx context.Context, posts []model.PostHashtag) ([]model.PostHashtag, error) {
	db := r.db.WithContext(ctx)
	err := db.Model(&model.PostHashtag{}).
		Create(&posts).
		Error
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func (r *PostHashtagDB) DeleteByPostId(ctx context.Context, postId uint) error {
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

func (r *PostHashtagDB) GetByPostId(ctx context.Context, postId uint) ([]model.PostHashtag, error) {
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

func (r *PostHashtagDB) Create(ctx context.Context, collection model.PostHashtag) (model.PostHashtag, error) {
	db := r.db.WithContext(ctx)
	err := db.Model(&model.PostHashtag{}).
		Create(&collection).
		Error
	if err != nil {
		return collection, err
	}
	return collection, nil
}

func (r *PostHashtagDB) DeleteByPostAndHashtagId(ctx context.Context, postId, hashtagId uint) error {
	db := r.db.WithContext(ctx)
	err := db.Model(&model.PostHashtag{}).
		Where("post_id = ? and hashtag_id = ?", postId, hashtagId).
		Delete(&model.PostHashtag{}).
		Error
	if err != nil {
		return err
	}
	return nil
}
