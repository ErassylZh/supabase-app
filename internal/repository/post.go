package repository

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"work-project/internal/model"
	"work-project/internal/schema"
)

type Post interface {
	CreateMany(ctx context.Context, posts []model.Post) ([]model.Post, error)
	GetAll(ctx context.Context) ([]model.Post, error)
	UpdateMany(ctx context.Context, posts []model.Post) ([]model.Post, error)
	GetAllForListing(ctx context.Context, filter schema.GetListingFilter) ([]model.Post, int64, error)
	GetAllByIds(ctx context.Context, ids []uint) ([]model.Post, error)
	DeleteAllNotInUuid(ctx context.Context, uuids []string) error
	GetAllGroupedByPostId(ctx context.Context, id uint) ([]model.Post, error)
}

type PostDB struct {
	db *gorm.DB
}

func NewPostDB(db *gorm.DB) *PostDB {
	return &PostDB{db: db}
}

func (r *PostDB) CreateMany(ctx context.Context, posts []model.Post) ([]model.Post, error) {
	db := r.db.WithContext(ctx)
	err := db.Model(&model.Post{}).
		Create(&posts).
		Error
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func (r *PostDB) GetAll(ctx context.Context) (posts []model.Post, err error) {
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

func (r *PostDB) UpdateMany(ctx context.Context, posts []model.Post) ([]model.Post, error) {
	// Открываем транзакцию
	err := r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		for _, post := range posts {
			if err := tx.Model(&model.Post{}).Where("post_id = ?", post.PostID).Updates(&post).Error; err != nil {
				return err // Если ошибка — транзакция отменяется
			}
		}
		return nil
	})

	if err != nil {
		return nil, err
	}
	return posts, nil
}

func (r *PostDB) GetAllForListing(ctx context.Context, filter schema.GetListingFilter) (posts []model.Post, total int64, err error) {
	db := r.db.WithContext(ctx)

	query := db.Model(&model.Post{})

	// If hashtagIds are provided, apply the filter
	if len(filter.HashtagIds) > 0 {
		query = query.Joins("JOIN public.post_hashtag ON public.post_hashtag.post_id = public.post.post_id").
			Joins("JOIN public.hashtag ON public.hashtag.hashtag_id = public.post_hashtag.hashtag_id").
			Where("public.hashtag.hashtag_id IN (?)", filter.HashtagIds)
	}
	if len(filter.CollectionIds) > 0 {
		query = query.Joins("JOIN public.post_collection ON public.post_collection.post_id = public.post.post_id").
			Joins("JOIN public.collection ON public.collection.collection_id = public.post_collection.collection_id").
			Where("public.collection.collection_id IN (?)", filter.CollectionIds)
	}
	if filter.Search != nil {
		search := fmt.Sprintf("%%%s%%", *filter.Search)
		query = query.Where("public.post.title ILIKE ? OR public.post.company ILIKE ?", search, search)
	}
	if filter.Language != nil {
		query = query.Where("public.post.language = ?", *filter.Language)
	}
	if len(filter.PostIds) > 0 {
		query = query.Where("public.post.post_id in (?)", filter.PostIds)
	}
	query = query.Where("public.post.status = ?", model.PRODUCT_STATUS_PUBLISH)
	err = query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	if filter.Page != 0 && filter.Size != 0 {
		query.Offset(filter.Size * (filter.Page - 1))
		query.Limit(filter.Size)
	}

	query = query.
		Preload("Images").
		Preload("Hashtags").
		Preload("Collections").
		Group("public.post.post_id").
		Order("post_order")

	err = query.Find(&posts).Error
	if err != nil {
		return nil, 0, err
	}

	return posts, total, nil
}

func (r *PostDB) GetAllByIds(ctx context.Context, ids []uint) (posts []model.Post, err error) {
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

func (r *PostDB) DeleteAllNotInUuid(ctx context.Context, uuids []string) error {
	db := r.db.WithContext(ctx)
	err := db.Model(&model.Post{}).
		Where("code not in (?)", uuids).
		Delete(&model.Post{}).
		Error
	if err != nil {
		return err
	}

	return nil
}

func (r *PostDB) GetAllGroupedByPostId(ctx context.Context, id uint) (posts []model.Post, err error) {
	db := r.db.WithContext(ctx)
	var uuid string

	err = db.Model(&model.Post{}).
		Select("uuid").
		Where("post_id = ?", id).
		Scan(&uuid).
		Error
	if err != nil {
		return []model.Post{}, err
	}

	err = db.Where("uuid IN (?)", uuid).
		Find(&posts).
		Error
	if err != nil {
		return []model.Post{}, err
	}

	return posts, nil
}
