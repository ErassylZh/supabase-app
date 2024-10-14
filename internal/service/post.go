package service

import (
	"context"
	"work-project/internal/model"
	"work-project/internal/repository"
	"work-project/internal/schema"
)

type Post interface {
	GetListing(ctx context.Context, hashtagIds []uint, collectionIds []uint, search string, language string, id *uint) ([]schema.PostResponse, error)
	GetByIds(ctx context.Context, ids []uint) ([]model.Post, error)
}

type PostService struct {
	postRepo repository.Post
}

func NewPostService(postRepo repository.Post) *PostService {
	return &PostService{postRepo: postRepo}
}

func (s *PostService) GetListing(ctx context.Context, hashtagIds []uint, collectionIds []uint, search string, language string, postId *uint) ([]schema.PostResponse, error) {
	posts, err := s.postRepo.GetAllForListing(ctx, hashtagIds, collectionIds, search, language, postId)
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

func (s *PostService) GetByIds(ctx context.Context, ids []uint) ([]model.Post, error) {
	return s.postRepo.GetAllByIds(ctx, ids)
}
