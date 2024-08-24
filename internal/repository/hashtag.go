package repository

import (
	"context"
	"gorm.io/gorm"
	"work-project/internal/model"
)

type Hashtag interface {
	GetByID(ctx context.Context, id uint) (model.Hashtag, error)
	GetByName(ctx context.Context, hashtagName string) (model.Hashtag, error)
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
