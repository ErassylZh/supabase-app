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
	DeleteManyByTitle(ctx context.Context, titles []string) error
	Create(ctx context.Context, story model.Stories) (model.Stories, error)
	GetByID(ctx context.Context, storyID uint) (model.Stories, error)
	Update(ctx context.Context, story model.Stories) (model.Stories, error)
	DeleteByID(ctx context.Context, id uint) error
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
		Order("start_time").
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

func (r *StoriesDB) UpdateMany(ctx context.Context, stories []model.Stories) ([]model.Stories, error) {
	err := r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		for _, story := range stories {
			if err := tx.Model(&model.Stories{}).Where("stories_id = ?", story.StoriesId).Updates(&story).Error; err != nil {
				return err // Если ошибка — транзакция отменяется
			}
		}
		return nil
	})

	if err != nil {
		return nil, err
	}
	return stories, nil
}

func (r *StoriesDB) DeleteManyByTitle(ctx context.Context, titles []string) error {
	db := r.db.WithContext(ctx)
	err := db.Model(&model.Stories{}).
		Where("title in (?)", titles).
		Delete(&model.Stories{}).
		Error
	if err != nil {
		return err
	}

	return nil
}

func (r *StoriesDB) Create(ctx context.Context, story model.Stories) (model.Stories, error) {
	db := r.db.WithContext(ctx)
	q := db.Model(&model.Stories{})
	err := q.Create(&story).Error
	if err != nil {
		return story, err
	}
	return story, nil
}

func (r *StoriesDB) GetByID(ctx context.Context, storyID uint) (story model.Stories, err error) {
	db := r.db.WithContext(ctx)
	q := db.Model(&model.Stories{})
	err = q.Where("stories_id = ?", storyID).
		Preload("StoryPages").
		First(&story).
		Error
	if err != nil {
		return story, err
	}
	return story, nil
}

func (r *StoriesDB) Update(ctx context.Context, story model.Stories) (model.Stories, error) {
	db := r.db.WithContext(ctx)
	q := db.Model(&model.Stories{})
	err := q.Where("stories_id = ?", story.StoriesId).
		Save(&story).
		Error
	if err != nil {
		return story, err
	}
	return story, nil
}

func (r *StoriesDB) DeleteByID(ctx context.Context, id uint) error {
	db := r.db.WithContext(ctx)
	q := db.Model(&model.Stories{})
	err := q.Where("stories_id = ?", id).
		Delete(&model.Stories{}).
		Error
	if err != nil {
		return err
	}
	return nil
}
