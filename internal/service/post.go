package service

import (
	"context"
	"work-project/internal/repository"
	"work-project/internal/schema"
)

type Post interface {
	GetListing(ctx context.Context, userId *string) ([]schema.PostResponse, error)
}

type PostService struct {
	postRepo repository.Post
}

func NewPostService(postRepo repository.Post) *PostService {
	return &PostService{postRepo: postRepo}
}

func (s *PostService) GetListing(ctx context.Context, userId *string) ([]schema.PostResponse, error) {
	posts, err := s.postRepo.GetAllForListing(ctx)
	if err != nil {
		return nil, err
	}
	result := make([]schema.PostResponse, len(posts))
	for i, post := range posts {
		data := schema.PostResponse{Post: post}
		result[i] = data
	}

	return result, nil
}
