package repository

import (
	"context"
	"gorm.io/gorm"
	"work-project/internal/model"
)

type Mark interface {
	CreateMark(ctx context.Context, mark *model.Mark) error
	FindByUserID(ctx context.Context, userID string) ([]model.Mark, error)
	DeleteMark(ctx context.Context, markID uint) error
	FindPostsByUserID(ctx context.Context, userID string) ([]model.Post, error)
}

type MarkDb struct {
	db *gorm.DB
}

func NewMarkDb(db *gorm.DB) *MarkDb {
	return &MarkDb{db: db}
}

func (r *MarkDb) CreateMark(ctx context.Context, marks []model.Mark) ([]model.Mark, error) {
	db := r.db.WithContext(ctx)
	err := db.Model(&model.Mark{}).
		Create(&marks).
		Error
	if err != nil {
		return nil, err
	}
	return marks, nil
}

func (r *MarkDb) FindByUserID(ctx context.Context, userID string) ([]model.Mark, error) {
	var marks []model.Mark
	db := r.db.WithContext(ctx)
	err := db.Where("user_id = ?", userID).Find(&marks).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return marks, nil
}

func (r *MarkDb) DeleteMark(ctx context.Context, markID uint) error {
	db := r.db.WithContext(ctx)
	q := db.Model(&model.Mark{})
	err := q.Where("mark_id = ?", markID).Delete(&model.Mark{}).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *MarkDb) FindPostsByUserID(ctx context.Context, userID string) ([]model.Post, error) {
	var marks []model.Mark
	db := r.db.WithContext(ctx)
	err := db.Model(&model.Mark{}).
		Where("user_id = ?", userID).
		Preload("Post").
		Find(&marks).Error
	if err != nil {
		return nil, err
	}
	var posts []model.Post
	for _, mark := range marks {
		posts = append(posts, mark.Post)
	}
	return posts, nil
}
