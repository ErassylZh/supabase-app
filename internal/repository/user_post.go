package repository

import (
	"context"
	"gorm.io/gorm"
	"work-project/internal/model"
)

type UserPost interface {
	Create(ctx context.Context, userPost model.UserPost) (model.UserPost, error)
	GetByUserAndPost(ctx context.Context, userId string, postId uint) (model.UserPost, error)
	Update(ctx context.Context, userPost model.UserPost) (model.UserPost, error)
	GetByUserID(ctx context.Context, id string) ([]model.UserPost, error)
}

type UserPostRepository struct {
	db *gorm.DB
}

func NewUserPostRepository(db *gorm.DB) *UserPostRepository {
	return &UserPostRepository{db: db}
}

func (r *UserPostRepository) Create(ctx context.Context, userPost model.UserPost) (model.UserPost, error) {
	db := r.db.WithContext(ctx)
	q := db.Model(&model.UserPost{})
	err := q.Create(&userPost).
		Error
	if err != nil {
		return model.UserPost{}, err
	}
	return userPost, nil
}

func (r *UserPostRepository) GetByUserAndPost(ctx context.Context, userId string, postId uint) (userPost model.UserPost, err error) {
	db := r.db.WithContext(ctx)
	q := db.Model(&model.UserPost{})
	err = q.Where("user_id = ? and post_id = ?", userId, postId).
		First(&userPost).
		Error
	if err != nil {
		return model.UserPost{}, err
	}
	return userPost, nil
}

func (r *UserPostRepository) Update(ctx context.Context, userPost model.UserPost) (model.UserPost, error) {
	db := r.db.WithContext(ctx)
	q := db.Model(&model.UserPost{})
	err := q.Where("user_post_id = ?", userPost.UserPostId).
		Save(&userPost).
		Error
	if err != nil {
		return model.UserPost{}, err
	}
	return userPost, nil
}
func (r *UserPostRepository) GetByUserID(ctx context.Context, id string) (userPosts []model.UserPost, err error) {
	db := r.db.WithContext(ctx)
	q := db.Model(&model.UserPost{})
	err = q.Where("user_id = ? ", id).
		Find(&userPosts).
		Error
	if err != nil {
		return []model.UserPost{}, err
	}
	return userPosts, nil
}
