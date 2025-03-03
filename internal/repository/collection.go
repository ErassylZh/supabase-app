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
	GetAllCollection(ctx context.Context, language string, withoutPosts bool) ([]model.Collection, error)
	GetAllRecommendation(ctx context.Context, language string) ([]model.Collection, error)
	CreateMany(ctx context.Context, collections []model.Collection) ([]model.Collection, error)
	UpdateMany(ctx context.Context, collections []model.Collection) ([]model.Collection, error)
	DeleteMany(ctx context.Context, collections []uint) error
	Create(ctx context.Context, data model.Collection) (model.Collection, error)
	Update(ctx context.Context, data model.Collection) (model.Collection, error)
	Delete(ctx context.Context, id uint) error
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

func (r *CollectionDB) GetAllCollection(ctx context.Context, language string, withoutPosts bool) (collections []model.Collection, err error) {
	db := r.db.WithContext(ctx)
	q := db.Model(&model.Collection{})
	if !withoutPosts {
		q = q.Preload("Posts", "language = ?", language).
			Preload("Posts.Images").
			Preload("Posts.Hashtags")
	}

	err = q.Where("not is_recommendation").
		Find(&collections).
		Error
	if err != nil {
		return collections, err
	}
	return collections, nil
}

func (r *CollectionDB) GetAllRecommendation(ctx context.Context, language string) (collections []model.Collection, err error) {
	db := r.db.WithContext(ctx)
	q := db.Model(&model.Collection{})
	err = q.Preload("Posts", "language = ?", language).
		Preload("Posts.Images").
		Preload("Posts.Hashtags").
		Preload("Posts.Collections").
		Where("is_recommendation").
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
	err := r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		for _, collection := range collections {
			if err := tx.Model(&model.Collection{}).Where("collection_id = ?", collection.CollectionID).Updates(map[string]interface{}{
				"name":              collection.Name,
				"name_ru":           collection.NameRu,
				"name_kz":           collection.NameKz,
				"image_path":        collection.ImagePath,
				"image_path_kz":     collection.ImagePathKz,
				"image_path_ru":     collection.ImagePathRu,
				"is_recommendation": collection.IsRecommendation, // Явное указание поля
			}).Error; err != nil {
				return err // Если ошибка — транзакция отменяется
			}
		}
		return nil
	})

	if err != nil {
		return nil, err
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

func (r *CollectionDB) Create(ctx context.Context, collection model.Collection) (model.Collection, error) {
	err := r.db.WithContext(ctx).Create(&collection).Error
	return collection, err
}

// Обновление коллекции
func (r *CollectionDB) Update(ctx context.Context, collection model.Collection) (model.Collection, error) {
	err := r.db.WithContext(ctx).Save(&collection).Error
	return collection, err
}

// Удаление коллекции
func (r *CollectionDB) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&model.Collection{}, id).Error
}
