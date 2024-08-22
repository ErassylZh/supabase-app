package repository

import (
	"context"
	"gorm.io/gorm"
	"work-project/internal/model"
)

type Mark interface {
	Create(ctx context.Context, mark *model.Mark) error
	FindByUserID(ctx context.Context, userID uint) ([]model.Mark, error)
	Delete(ctx context.Context, markID uint) error
	FindPostsByUserID(ctx context.Context, userID uint) ([]model.Post, error)
}

type MarkDb struct {
	db *gorm.DB
}

func NewMarkDb(db *gorm.DB) *MarkDb {
	return &MarkDb{db: db}
}

func (r *MarkDb) Create(ctx context.Context, marks []model.Mark) ([]model.Mark, error) {
	db := r.db.WithContext(ctx)
	err := db.Model(&model.Mark{}).
		Create(&marks).
		Error
	if err != nil {
		return nil, err
	}
	return marks, nil
}

func (r *MarkDb) FindByUserID(ctx context.Context, userID uint) ([]model.Mark, error) {
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

func (r *MarkDb) Delete(ctx context.Context, markID uint) error {
	db := r.db.WithContext(ctx)
	q := db.Model(&model.Mark{})
	err := q.Where("mark_id = ?", markID).Delete(&model.Mark{}).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *MarkDb) FindPostsByUserID(ctx context.Context, userID uint) ([]model.Post, error) {
	var posts []model.Post
	db := r.db.WithContext(ctx)

	err := db.Table("marks").
		Select("posts.*").
		Joins("JOIN posts ON marks.post_id = posts.id").
		Where("marks.user_id = ?", userID).
		Scan(&posts).Error

	if err != nil {
		return nil, err
	}
	return posts, nil
}
