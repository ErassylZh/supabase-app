package repository

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"work-project/internal/model"
)

type Mark interface {
	Create(ctx context.Context, marks model.Mark) error
	FindByUserID(ctx context.Context, userID string) ([]model.Mark, error)
	DeleteMark(ctx context.Context, markID uint) error
	FindByUserAndPost(ctx context.Context, userId string, postId uint) (model.Mark, error)
}

type MarkDb struct {
	db *gorm.DB
}

func NewMarkDb(db *gorm.DB) *MarkDb {
	return &MarkDb{db: db}
}

func (r *MarkDb) Create(ctx context.Context, marks model.Mark) error {
	db := r.db.WithContext(ctx)
	err := db.Model(&model.Mark{}).
		Create(&marks).
		Error
	if err != nil {
		return err
	}
	return nil
}

func (r *MarkDb) FindByUserID(ctx context.Context, userID string) ([]model.Mark, error) {
	var marks []model.Mark
	db := r.db.WithContext(ctx)
	err := db.Where("user_id = ?", userID).
		Preload("Post").
		Find(&marks).
		Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
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

func (r *MarkDb) FindByUserAndPost(ctx context.Context, userId string, markId uint) (mark model.Mark, err error) {
	db := r.db.WithContext(ctx)
	q := db.Model(&model.Mark{})
	err = q.Where("mark_id = ? and user_id = ?", markId, userId).
		First(&mark).
		Error
	if err != nil {
		return mark, err
	}
	return mark, nil
}
