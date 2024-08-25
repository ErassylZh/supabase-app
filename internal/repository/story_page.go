package repository

import (
	"context"
	"gorm.io/gorm"
	"work-project/internal/model"
)

type StoryPage interface {
	GetAllByStoryId(ctx context.Context, storyId uint) ([]model.StoryPage, error)
	CreateMany(ctx context.Context, stories []model.StoryPage) error
}

type StoryPageDB struct {
	db *gorm.DB
}

func NewStoryPageDB(db *gorm.DB) *StoryPageDB {
	return &StoryPageDB{db: db}
}

func (r *StoryPageDB) GetAllByStoryId(ctx context.Context, storyId uint) (stories []model.StoryPage, err error) {
	db := r.db.WithContext(ctx)
	q := db.Model(&model.StoryPage{})
	err = q.Where("story_id=?", storyId).
		Find(&stories).
		Error
	if err != nil {
		return stories, err
	}
	return stories, nil
}

func (r *StoryPageDB) CreateMany(ctx context.Context, stories []model.StoryPage) error {
	db := r.db.WithContext(ctx)
	q := db.Model(&model.StoryPage{})
	err := q.Create(&stories).Error
	return err
}
