package service

import (
	"context"
	"work-project/internal/model"
	"work-project/internal/repository"
)

type Collection interface {
	GetAll(ctx context.Context) ([]model.Collection, error)
}

type CollectionService struct {
	collectionRepo repository.Collection
}

func NewCollectionService(collectionRepo repository.Collection) *CollectionService {
	return &CollectionService{collectionRepo: collectionRepo}
}

func (s *CollectionService) GetAll(ctx context.Context) ([]model.Collection, error) {
	return s.collectionRepo.GetAll(ctx)
}
