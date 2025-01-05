package service

import (
	"context"
	"work-project/internal/model"
	"work-project/internal/repository"
)

type ProductTag interface {
	GetAll(ctx context.Context) ([]model.ProductTag, error)
}

type ProductTagService struct {
	hashtagRepo repository.ProductTag
}

func NewProductTagService(hashtagRepo repository.ProductTag) *ProductTagService {
	return &ProductTagService{hashtagRepo: hashtagRepo}
}

func (s *ProductTagService) GetAll(ctx context.Context) ([]model.ProductTag, error) {
	return s.hashtagRepo.GetAll(ctx)
}
