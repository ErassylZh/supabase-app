package service

import (
	"context"
	"work-project/internal/model"
	"work-project/internal/repository"
)

type Collection interface {
	GetAllCollection(ctx context.Context, language string) ([]model.Collection, error)
	GetAllRecommendation(ctx context.Context, language string) ([]model.Collection, error)
}

type CollectionService struct {
	collectionRepo repository.Collection
}

func NewCollectionService(collectionRepo repository.Collection) *CollectionService {
	return &CollectionService{collectionRepo: collectionRepo}
}

func (s *CollectionService) GetAllCollection(ctx context.Context, language string) ([]model.Collection, error) {
	return s.collectionRepo.GetAllCollection(ctx, language)
}

func (s *CollectionService) GetAllRecommendation(ctx context.Context, language string) ([]model.Collection, error) {
	return s.collectionRepo.GetAllRecommendation(ctx, language)
}
