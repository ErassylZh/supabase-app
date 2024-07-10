package repository

import (
	"context"
	"gorm.io/gorm"
	"work-project/internal/model"
)

type User interface {
	GetByID(ctx context.Context, userId string) (model.User, error)
	DeleteByID(ctx context.Context, userId string) error
}

type UserDB struct {
	db *gorm.DB
}

func NewUserDB(db *gorm.DB) *UserDB {
	return &UserDB{db: db}
}

func (r *UserDB) GetByID(ctx context.Context, userId string) (user model.User, err error) {
	db := r.db.WithContext(ctx)
	q := db.Model(&model.User{})
	err = q.Where("id = ?", userId).
		First(&user).
		Error
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}

func (r *UserDB) DeleteByID(ctx context.Context, userId string) (err error) {
	db := r.db.WithContext(ctx)
	q := db.Model(&model.User{})
	err = q.Where("id = ?", userId).
		Delete(&model.User{}).
		Error
	if err != nil {
		return err
	}
	return nil
}
