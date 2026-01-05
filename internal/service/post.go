package service

import (
	"context"
	"work-project/internal/model"
	"work-project/internal/repository"
	"work-project/internal/schema"
)

type Post interface {
	GetListing(ctx context.Context, filter schema.GetListingFilter) ([]schema.PostResponse, int64, error)
	GetByIds(ctx context.Context, ids []uint) ([]model.Post, error)
}

type PostService struct {
	postRepo repository.Post
}

func NewPostService(postRepo repository.Post) *PostService {
	return &PostService{postRepo: postRepo}
}

func (s *PostService) GetListing(ctx context.Context, filter schema.GetListingFilter) ([]schema.PostResponse, int64, error) {
	posts, total, err := s.postRepo.GetAllForListing(ctx, filter)
	if err != nil {
		return nil, 0, err
	}
	result := make([]schema.PostResponse, len(posts))
	for i, post := range posts {
		data := schema.PostResponse{Post: post}
		for _, hashtag := range post.Hashtags {
			if hashtag.Name == string(model.HASHTAG_NAME_PARTNER) {
				data.PostType = "partner"
				break
			}
			if hashtag.Name == string(model.HASHTAG_NAME_BESTSELLER) {
				data.PostType = "post"
				break
			}
		}
		result[i] = data
	}

	return result, total, nil
}

func (s *PostService) GetByIds(ctx context.Context, ids []uint) ([]model.Post, error) {
	return s.postRepo.GetAllByIds(ctx, ids)
}
