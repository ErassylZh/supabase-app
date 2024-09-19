package repository

import (
	"context"
	"gorm.io/gorm"
	"work-project/internal/model"
)

type Collection interface {
	GetByID(ctx context.Context, id uint) (model.Collection, error)
	GetByName(ctx context.Context, collectionName string) (model.Collection, error)
	GetAll(ctx context.Context) ([]model.Collection, error)
	GetAllCollection(ctx context.Context) ([]model.Collection, error)
	GetAllRecommendation(ctx context.Context) ([]model.Collection, error)
	CreateMany(ctx context.Context, collections []model.Collection) ([]model.Collection, error)
	UpdateMany(ctx context.Context, collections []model.Collection) ([]model.Collection, error)
	DeleteMany(ctx context.Context, collections []uint) error
}

type CollectionDB struct {
	db *gorm.DB
}

func NewCollectionDB(db *gorm.DB) *CollectionDB {
	return &CollectionDB{db: db}
}

func (r *CollectionDB) GetByID(ctx context.Context, id uint) (collection model.Collection, err error) {
	db := r.db.WithContext(ctx)
	q := db.Model(&model.Collection{})
	err = q.Where("collection_id = ?", id).
		First(&collection).
		Error
	if err != nil {
		return collection, err
	}
	return collection, nil
}

func (r *CollectionDB) GetByName(ctx context.Context, collectionName string) (collection model.Collection, err error) {
	db := r.db.WithContext(ctx)
	q := db.Model(&model.Collection{})
	err = q.Where("name = ?", collectionName).
		First(&collection).
		Error
	if err != nil {
		return collection, err
	}
	return collection, nil
}
func (r *CollectionDB) GetAll(ctx context.Context) (collections []model.Collection, err error) {
	db := r.db.WithContext(ctx)
	q := db.Model(&model.Collection{})
	err = q.Preload("Posts").
		Find(&collections).
		Error
	if err != nil {
		return collections, err
	}
	return collections, nil
}

func (r *CollectionDB) GetAllCollection(ctx context.Context) (collections []model.Collection, err error) {
	db := r.db.WithContext(ctx)
	q := db.Model(&model.Collection{})
	err = q.Where("not is_recommendation").
		Preload("Posts").
		Preload("Posts.Images").
		Preload("Posts.Hashtags").
		Preload("Posts.Collections").
		Find(&collections).
		Error
	if err != nil {
		return collections, err
	}
	return collections, nil
}

func (r *CollectionDB) GetAllRecommendation(ctx context.Context) (collections []model.Collection, err error) {
	db := r.db.WithContext(ctx)
	q := db.Model(&model.Collection{})
	err = q.Where("is_recommendation").
		Preload("Posts").
		Preload("Posts.Images").
		Preload("Posts.Hashtags").
		Preload("Posts.Collections").
		Find(&collections).
		Error
	if err != nil {
		return collections, err
	}
	return collections, nil
}

func (r *CollectionDB) CreateMany(ctx context.Context, collections []model.Collection) ([]model.Collection, error) {
	db := r.db.WithContext(ctx)
	q := db.Model(&model.Collection{})
	err := q.Create(&collections).
		Error
	if err != nil {
		return collections, err
	}
	return collections, nil
}

func (r *CollectionDB) UpdateMany(ctx context.Context, collections []model.Collection) ([]model.Collection, error) {
	db := r.db.WithContext(ctx)

	for _, post := range collections {
		if err := db.Model(&model.Collection{}).Where("collection_id = ?", post.CollectionID).Updates(&post).Error; err != nil {
			return nil, err
		}
	}
	return collections, nil
}

func (r *CollectionDB) DeleteMany(ctx context.Context, collectionIds []uint) error {
	db := r.db.WithContext(ctx)
	q := db.Model(&model.Collection{})
	err := q.Where("collection_id in (?)", collectionIds).
		Delete(&model.Collection{}).
		Error
	if err != nil {
		return err
	}
	return nil
}
