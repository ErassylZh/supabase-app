package repository

import (
	"context"
	"gorm.io/gorm"
	"work-project/internal/model"
)

type StoryPage interface {
	GetAllByStoryId(ctx context.Context, storyId uint) ([]model.StoryPage, error)
	CreateMany(ctx context.Context, stories []model.StoryPage) error
	UpdateMany(ctx context.Context, pages []model.StoryPage) ([]model.StoryPage, error)
	DeleteManyByUuid(ctx context.Context, uuids []string) error
	Create(ctx context.Context, page model.StoryPage) (model.StoryPage, error)
	Update(ctx context.Context, page model.StoryPage) (model.StoryPage, error)
	GetByID(ctx context.Context, id uint) (model.StoryPage, error)
	Delete(ctx context.Context, id uint) error
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

func (r *StoryPageDB) UpdateMany(ctx context.Context, pages []model.StoryPage) ([]model.StoryPage, error) {
	err := r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		for _, page := range pages {
			if err := tx.Model(&model.StoryPage{}).Where("story_page_id = ?", page.StoryPageId).Updates(&page).Error; err != nil {
				return err // Если ошибка — транзакция отменяется
			}
		}
		return nil
	})

	if err != nil {
		return nil, err
	}
	return pages, nil

}

func (r *StoryPageDB) DeleteManyByUuid(ctx context.Context, uuids []string) error {
	db := r.db.WithContext(ctx)
	q := db.Model(&model.StoryPage{})
	err := q.Where("uuid not in (?)", uuids).Delete(&model.StoryPage{}).Error
	return err
}

func (r *StoryPageDB) Create(ctx context.Context, page model.StoryPage) (model.StoryPage, error) {
	db := r.db.WithContext(ctx)
	q := db.Model(&model.StoryPage{})
	err := q.Create(&page).Error
	return page, err
}

func (r *StoryPageDB) Update(ctx context.Context, page model.StoryPage) (model.StoryPage, error) {
	db := r.db.WithContext(ctx)
	q := db.Model(&model.StoryPage{})
	err := q.Where("story_page_id = ?", page.StoryPageId).
		Save(&page).
		Error
	return page, err
}

func (r *StoryPageDB) GetByID(ctx context.Context, id uint) (story model.StoryPage, err error) {
	db := r.db.WithContext(ctx)
	q := db.Model(&model.StoryPage{})
	err = q.Where("story_page_id=?", id).
		First(&story).
		Error
	if err != nil {
		return story, err
	}
	return story, nil
}

func (r *StoryPageDB) Delete(ctx context.Context, id uint) error {
	db := r.db.WithContext(ctx)
	q := db.Model(&model.StoryPage{})
	err := q.Where("story_page_id = ?", id).
		Delete(&model.StoryPage{}).
		Error
	return err
}
