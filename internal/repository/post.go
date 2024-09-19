package repository

import (
	"context"
	"gorm.io/gorm"
	"work-project/internal/model"
)

type Post interface {
	CreateMany(ctx context.Context, posts []model.Post) ([]model.Post, error)
	GetAll(ctx context.Context) ([]model.Post, error)
	UpdateMany(ctx context.Context, posts []model.Post) ([]model.Post, error)
	GetAllForListing(ctx context.Context, hashtagIds []uint, collectionIds []uint) ([]model.Post, error)
	GetAllByIds(ctx context.Context, ids []uint) ([]model.Post, error)
	DeleteAllNotInUuid(ctx context.Context, uuids []string) error
}

type PostDb struct {
	db *gorm.DB
}

func NewPostDb(db *gorm.DB) *PostDb {
	return &PostDb{db: db}
}

func (r *PostDb) CreateMany(ctx context.Context, posts []model.Post) ([]model.Post, error) {
	db := r.db.WithContext(ctx)
	err := db.Model(&model.Post{}).
		Create(&posts).
		Error
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func (r *PostDb) GetAll(ctx context.Context) (posts []model.Post, err error) {
	db := r.db.WithContext(ctx)
	err = db.Model(&model.Post{}).
		Preload("Hashtags").
		Preload("Collections").
		Find(&posts).
		Error
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func (r *PostDb) UpdateMany(ctx context.Context, posts []model.Post) ([]model.Post, error) {
	db := r.db.WithContext(ctx)

	for _, post := range posts {
		if err := db.Model(&model.Post{}).Where("post_id = ?", post.PostID).Updates(&post).Error; err != nil {
			return nil, err
		}
	}
	return posts, nil
}
func (r *PostDb) GetAllForListing(ctx context.Context, hashtagIds []uint, collectionIds []uint) (posts []model.Post, err error) {
	db := r.db.WithContext(ctx)

	query := db.Model(&model.Post{})

	// If hashtagIds are provided, apply the filter
	if len(hashtagIds) > 0 {
		query = query.Joins("JOIN public.post_hashtag ON public.post_hashtag.post_id = public.post.post_id").
			Joins("JOIN public.hashtag ON public.hashtag.hashtag_id = public.post_hashtag.hashtag_id").
			Where("public.hashtag.hashtag_id IN (?)", hashtagIds)
	}
	if len(collectionIds) > 0 {
		query = query.Joins("JOIN public.post_collection ON public.post_collection.post_id = public.post.post_id").
			Joins("JOIN public.collection ON public.collection.collection_id = public.post_collection.collection_id").
			Where("public.collection.collection_id IN (?)", collectionIds)
	}
	query = query.
		Where("public.post.status = ?", model.PRODUCT_STATUS_PUBLISH).
		Preload("Images").
		Preload("Hashtags").
		Preload("Collections").
		Group("public.post.post_id")

	err = query.Find(&posts).Error
	if err != nil {
		return nil, err
	}

	return posts, nil
}

func (r *PostDb) GetAllByIds(ctx context.Context, ids []uint) (posts []model.Post, err error) {
	db := r.db.WithContext(ctx)
	err = db.Model(&model.Post{}).
		Where("status = ? and post_id in (?)", model.PRODUCT_STATUS_PUBLISH, ids).
		Preload("Images").
		Preload("Hashtags").
		Preload("Collections").
		Find(&posts).
		Error
	if err != nil {
		return nil, err
	}

	return posts, nil
}

func (r *PostDb) DeleteAllNotInUuid(ctx context.Context, uuids []string) error {
	db := r.db.WithContext(ctx)
	err := db.Model(&model.Post{}).
		Where("uuid not in (?)", uuids).
		Delete(&model.Post{}).
		Error
	if err != nil {
		return err
	}

	return nil
}
