package service

import (
	"context"
	"work-project/internal/model"
	"work-project/internal/repository"
)

type Hashtag interface {
	GetAll(ctx context.Context) ([]model.Hashtag, error)
}

type HashtagService struct {
	hashtagRepo repository.Hashtag
}

func NewHashtagService(hashtagRepo repository.Hashtag) *HashtagService {
	return &HashtagService{hashtagRepo: hashtagRepo}
}

func (s *HashtagService) GetAll(ctx context.Context) ([]model.Hashtag, error) {
	return s.hashtagRepo.GetAll(ctx)
}
