package repository

import (
	"context"
	"gorm.io/gorm"
	"work-project/internal/model"
)

type StoryPageUser interface {
	GetAllByStoryPageId(ctx context.Context, storyPageId uint) ([]model.StoryPageUser, error)
	GetAllByStoryIdUserID(ctx context.Context, storyId []uint, userId string) ([]model.StoryPageUser, error)
	Create(ctx context.Context, user model.StoryPageUser) error
}

type StoryPageUserDB struct {
	db *gorm.DB
}

func NewStoryPageUserDB(db *gorm.DB) *StoryPageUserDB {
	return &StoryPageUserDB{db: db}
}

func (r *StoryPageUserDB) GetAllByStoryPageId(ctx context.Context, storyPageId uint) (stories []model.StoryPageUser, err error) {
	db := r.db.WithContext(ctx)
	q := db.Model(&model.StoryPageUser{})
	err = q.Where("story_page_id=?", storyPageId).
		Find(&stories).
		Error
	if err != nil {
		return stories, err
	}
	return stories, nil
}

func (r *StoryPageUserDB) Create(ctx context.Context, data model.StoryPageUser) error {
	db := r.db.WithContext(ctx)
	q := db.Model(&model.StoryPageUser{})
	err := q.Create(&data).Error
	return err
}

func (r *StoryPageUserDB) GetAllByStoryIdUserID(ctx context.Context, storyIds []uint, userId string) (stories []model.StoryPageUser, err error) {
	db := r.db.WithContext(ctx)

	err = db.Table("story_page_user AS spu").
		Select("spu.*").
		Joins("JOIN story_page sp ON spu.story_page_id = sp.story_page_id").
		Joins("JOIN stories s ON s.stories_id = sp.stories_id and s.stories_id in (?)", storyIds).
		Where("spu.user_id = ?", userId).
		Find(&stories).
		Error
	if err != nil {
		return stories, err
	}
	return stories, nil
}
