package repository

import (
	"context"
	"gorm.io/gorm"
	"time"
	"work-project/internal/model"
)

type Stories interface {
	GetAll(ctx context.Context) ([]model.Stories, error)
	CreateMany(ctx context.Context, story []model.Stories) ([]model.Stories, error)
	GetAllActive(ctx context.Context) ([]model.Stories, error)
	GetAllActiveByUser(ctx context.Context, userId string) ([]model.Stories, error)
	UpdateMany(ctx context.Context, posts []model.Stories) ([]model.Stories, error)
}

type StoriesDB struct {
	db *gorm.DB
}

func NewStoriesDB(db *gorm.DB) *StoriesDB {
	return &StoriesDB{db: db}
}

func (r *StoriesDB) GetAll(ctx context.Context) (stories []model.Stories, err error) {
	db := r.db.WithContext(ctx)
	q := db.Model(&model.Stories{})
	err = q.Preload("StoryPages").
		Find(&stories).
		Error
	if err != nil {
		return stories, err
	}
	return stories, nil
}

func (r *StoriesDB) CreateMany(ctx context.Context, stories []model.Stories) ([]model.Stories, error) {
	db := r.db.WithContext(ctx)
	q := db.Model(&model.Stories{})
	err := q.Create(&stories).Error
	if err != nil {
		return nil, err
	}
	return stories, nil
}

func (r *StoriesDB) GetAllActive(ctx context.Context) (stories []model.Stories, err error) {
	db := r.db.WithContext(ctx)
	q := db.Model(&model.Stories{})
	err = q.Preload("StoryPages", func(db *gorm.DB) *gorm.DB {
		return db.Order("page_order ASC")
	}).
		Where("start_time < ? and end_time > ?", time.Now(), time.Now()).
		Find(&stories).
		Error
	if err != nil {
		return stories, err
	}
	return stories, nil
}

func (r *StoriesDB) GetAllActiveByUser(ctx context.Context, userId string) (stories []model.Stories, err error) {
	db := r.db.WithContext(ctx)
	var ids []uint
	if len(userId) > 0 {
		db.Table("story_page_user").
			Select("story_page_id").
			Where("user_id = ?", userId).Scan(&ids)
	}

	db.Table("stories AS s").
		Select("s.*").
		Joins("JOIN story_page sp ON s.stories_id = sp.stories_id").
		Where("sp.story_page_id NOT IN (?)", ids).
		Preload("StoryPages").
		Find(&stories)
	if err != nil {
		return stories, err
	}
	return stories, nil
}

func (r *StoriesDB) UpdateMany(ctx context.Context, posts []model.Stories) ([]model.Stories, error) {
	db := r.db.WithContext(ctx)

	for _, post := range posts {
		if err := db.Model(&model.Stories{}).Where("stories_id = ?", post.StoriesId).Updates(&post).Error; err != nil {
			return nil, err
		}
	}
	return posts, nil
}
