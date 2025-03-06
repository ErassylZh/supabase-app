package repository

import (
	"context"
	"gorm.io/gorm"
	"work-project/internal/model"
)

type PostCollection interface {
	CreateMany(ctx context.Context, posts []model.PostCollection) ([]model.PostCollection, error)
	Create(ctx context.Context, posts model.PostCollection) (model.PostCollection, error)
	DeleteByPostAndCollectionId(ctx context.Context, postId, collectionId uint) error
	GetByPostId(ctx context.Context, postId uint) ([]model.PostCollection, error)
	GetByPostAndCollectionId(ctx context.Context, postId, collectionId uint) (model.PostCollection, error)
}

type PostCollectionDB struct {
	db *gorm.DB
}

func NewPostCollectionDB(db *gorm.DB) *PostCollectionDB {
	return &PostCollectionDB{db: db}
}

func (r *PostCollectionDB) Create(ctx context.Context, posts model.PostCollection) (model.PostCollection, error) {
	db := r.db.WithContext(ctx)
	err := db.Model(&model.PostCollection{}).
		Create(&posts).
		Error
	if err != nil {
		return posts, err
	}
	return posts, nil
}

func (r *PostCollectionDB) CreateMany(ctx context.Context, posts []model.PostCollection) ([]model.PostCollection, error) {
	db := r.db.WithContext(ctx)
	err := db.Model(&model.PostCollection{}).
		Create(&posts).
		Error
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func (r *PostCollectionDB) DeleteByPostId(ctx context.Context, postId uint) error {
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

func (r *PostCollectionDB) DeleteByPostAndCollectionId(ctx context.Context, postId, collectionId uint) error {
	db := r.db.WithContext(ctx)
	err := db.Model(&model.PostCollection{}).
		Where("post_id = ? and collection_id", postId, collectionId).
		Delete(&model.PostCollection{}).
		Error
	if err != nil {
		return err
	}
	return nil
}

func (r *PostCollectionDB) GetByPostId(ctx context.Context, postId uint) ([]model.PostCollection, error) {
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

func (r *PostCollectionDB) GetByPostAndCollectionId(ctx context.Context, postId, collectionId uint) (model.PostCollection, error) {
	db := r.db.WithContext(ctx)
	var data model.PostCollection
	err := db.Model(&model.PostCollection{}).
		Where("post_id = ? and collection_id = ?", postId, collectionId).
		First(&data).
		Error
	if err != nil {
		return data, err
	}
	return data, nil
}
