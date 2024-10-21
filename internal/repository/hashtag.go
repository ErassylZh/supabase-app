package repository

import (
	"context"
	"gorm.io/gorm"
	"work-project/internal/model"
)

type Hashtag interface {
	GetByID(ctx context.Context, id uint) (model.Hashtag, error)
	GetByName(ctx context.Context, hashtagName string) (model.Hashtag, error)
	GetAll(ctx context.Context) ([]model.Hashtag, error)
	CreateMany(ctx context.Context, hashtags []model.Hashtag) ([]model.Hashtag, error)
	UpdateMany(ctx context.Context, hashtags []model.Hashtag) ([]model.Hashtag, error)
	DeleteMany(ctx context.Context, hashtagIds []uint) error
	GetVisible(ctx context.Context) ([]model.Hashtag, error)
}

type HashtagDB struct {
	db *gorm.DB
}

func NewHashtagDB(db *gorm.DB) *HashtagDB {
	return &HashtagDB{db: db}
}

func (r *HashtagDB) GetByID(ctx context.Context, id uint) (hashtag model.Hashtag, err error) {
	db := r.db.WithContext(ctx)
	q := db.Model(&model.Hashtag{})
	err = q.Where("hashtag_id = ?", id).
		First(&hashtag).
		Error
	if err != nil {
		return hashtag, err
	}
	return hashtag, nil
}

func (r *HashtagDB) GetByName(ctx context.Context, hashtagName string) (hashtag model.Hashtag, err error) {
	db := r.db.WithContext(ctx)
	q := db.Model(&model.Hashtag{})
	err = q.Where("name = ?", hashtagName).
		First(&hashtag).
		Error
	if err != nil {
		return hashtag, err
	}
	return hashtag, nil
}

func (r *HashtagDB) GetAll(ctx context.Context) (hashtags []model.Hashtag, err error) {
	db := r.db.WithContext(ctx)
	q := db.Model(&model.Hashtag{})
	err = q.Find(&hashtags).
		Error
	if err != nil {
		return hashtags, err
	}
	return hashtags, nil
}

func (r *HashtagDB) GetVisible(ctx context.Context) (hashtags []model.Hashtag, err error) {
	db := r.db.WithContext(ctx)
	q := db.Model(&model.Hashtag{})
	err = q.Where("is_visible").
		Find(&hashtags).
		Error
	if err != nil {
		return hashtags, err
	}
	return hashtags, nil
}

func (r *HashtagDB) CreateMany(ctx context.Context, hashtags []model.Hashtag) ([]model.Hashtag, error) {
	db := r.db.WithContext(ctx)
	q := db.Model(&model.Hashtag{})
	err := q.Create(&hashtags).
		Error
	if err != nil {
		return hashtags, err
	}
	return hashtags, nil
}

func (r *HashtagDB) UpdateMany(ctx context.Context, hashtags []model.Hashtag) ([]model.Hashtag, error) {
	db := r.db.WithContext(ctx)

	for _, hashtag := range hashtags {
		if err := db.Model(&model.Hashtag{}).Where("hashtag_id = ?", hashtag.HashtagID).Updates(map[string]interface{}{
			"name":       hashtag.Name,
			"name_ru":    hashtag.NameRu,
			"name_kz":    hashtag.NameKz,
			"image_path": hashtag.ImagePath,
			"is_visible": hashtag.IsVisible,
		}).Error; err != nil {
			return nil, err
		}
	}
	return hashtags, nil
}

func (r *HashtagDB) DeleteMany(ctx context.Context, hashtagIds []uint) error {
	db := r.db.WithContext(ctx)
	q := db.Model(&model.Hashtag{})
	err := q.Where("hashtag_id in (?)", hashtagIds).
		Delete(&model.Hashtag{}).
		Error
	if err != nil {
		return err
	}
	return nil
}
