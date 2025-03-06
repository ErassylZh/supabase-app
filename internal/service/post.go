package service

import (
	"context"
	"github.com/google/uuid"
	"log"
	"work-project/internal/admin"
	"work-project/internal/model"
	"work-project/internal/repository"
	"work-project/internal/schema"
)

type Post interface {
	GetListing(ctx context.Context, filter schema.GetListingFilter) ([]schema.PostResponse, int64, error)
	GetByIds(ctx context.Context, ids []uint) ([]model.Post, error)
	Create(ctx context.Context, data admin.CreatePost) (model.Post, error)
	Update(ctx context.Context, data admin.UpdatePost) (model.Post, error)
	Delete(ctx context.Context, id uint) error
	GetById(ctx context.Context, id uint) (model.Post, error)
}

type PostService struct {
	postRepo repository.Post
	image    repository.Image
	storage  repository.Storage
}

func NewPostService(postRepo repository.Post) *PostService {
	return &PostService{postRepo: postRepo}
}

func (s *PostService) GetListing(ctx context.Context, filter schema.GetListingFilter) ([]model.Post, int64, error) {
	return s.postRepo.GetAllForListing(ctx, filter)
}

func (s *PostService) GetById(ctx context.Context, id uint) (model.Post, error) {
	return s.postRepo.GetById(ctx, id)
}

func (s *PostService) GetByIds(ctx context.Context, ids []uint) ([]model.Post, error) {
	return s.postRepo.GetAllByIds(ctx, ids)
}

func (s *PostService) Create(ctx context.Context, data admin.CreatePost) (model.Post, error) {
	post := model.Post{
		Body:             &data.Body,
		Title:            data.Title,
		Company:          data.Company,
		Language:         data.Language,
		Description:      data.Description,
		Status:           string(model.POST_STATUS_DRAFT),
		ReadTime:         data.ReadTime,
		Point:            data.Point,
		Sapphire:         data.Sapphire,
		QuizTime:         data.QuizTime,
		RatingStatus:     data.RatingStatus,
		ShortDescription: data.ShortDescription,
		Uuid:             uuid.New().String(),
	}
	post, err := s.postRepo.Create(ctx, post)
	if err != nil {
		return model.Post{}, err
	}
	images := make([]model.Image, 0)

	if data.Logo != nil {
		file, err := s.storage.CreateImageFromBase64(ctx, string(model.BUCKET_NAME_PRODUCT), data.Logo.FileName, data.Logo.File)
		if err != nil {
			log.Println(ctx, "some err while create image", "err", err, "post name", post.Title)
			return model.Post{}, err
		}
		images = append(images, model.Image{
			PostID:   &post.PostID,
			ImageUrl: file,
			Type:     string(model.POST_IMAGE_TYPE_LOGO),
		})
	}

	if data.Image != nil {
		file, err := s.storage.CreateImageFromBase64(ctx, string(model.BUCKET_NAME_PRODUCT), data.Image.FileName, data.Image.File)
		if err != nil {
			log.Println(ctx, "some err while create image", "err", err, "post name", post.Title)
			return model.Post{}, err
		}
		images = append(images, model.Image{
			PostID:   &post.PostID,
			ImageUrl: file,
			Type:     string(model.POST_IMAGE_TYPE_IMAGE),
		})
	}

	if len(images) > 0 {
		_, err = s.image.CreateMany(ctx, images)
		return model.Post{}, err
	}

	return post, nil
}

func (s *PostService) Update(ctx context.Context, data admin.UpdatePost) (model.Post, error) {
	post, err := s.postRepo.GetById(ctx, data.PostID)
	if err != nil {
		return model.Post{}, err
	}
	images := make([]model.Image, 0)

	if data.Logo != nil {
		file, err := s.storage.CreateImageFromBase64(ctx, string(model.BUCKET_NAME_PRODUCT), data.Logo.FileName, data.Logo.File)
		if err != nil {
			log.Println(ctx, "some err while create image", "err", err, "post name", post.Title)
			return model.Post{}, err
		}
		images = append(images, model.Image{
			PostID:   &post.PostID,
			ImageUrl: file,
			Type:     string(model.POST_IMAGE_TYPE_LOGO),
		})
	}

	if data.Image != nil {
		file, err := s.storage.CreateImageFromBase64(ctx, string(model.BUCKET_NAME_PRODUCT), data.Image.FileName, data.Image.File)
		if err != nil {
			log.Println(ctx, "some err while create image", "err", err, "post name", post.Title)
			return model.Post{}, err
		}
		images = append(images, model.Image{
			PostID:   &post.PostID,
			ImageUrl: file,
			Type:     string(model.POST_IMAGE_TYPE_IMAGE),
		})
	}

	if len(images) > 0 {
		_, err = s.image.CreateMany(ctx, images)
		return model.Post{}, err
	}

	return s.postRepo.Update(ctx, post)
}

func (s *PostService) Delete(ctx context.Context, id uint) error {
	return s.postRepo.DeleteById(ctx, id)
}
