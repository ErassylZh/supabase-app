package service

import (
	"context"
	"work-project/internal/admin"
	"work-project/internal/model"
	"work-project/internal/repository"
)

type ProductTag interface {
	GetAll(ctx context.Context) ([]model.ProductTag, error)
	Create(ctx context.Context, data admin.CreateProductTag) (model.ProductTag, error)
	Update(ctx context.Context, data admin.UpdateProductTag) (model.ProductTag, error)
	Delete(ctx context.Context, productTagID uint) error
	AddToProduct(ctx context.Context, data admin.AddProductProductTag) (model.ProductTag, error)
	DeleteToProduct(ctx context.Context, data admin.DeleteProductProductTag) (model.ProductTag, error)
}

type ProductTagService struct {
	productTagRepo        repository.ProductTag
	productProductTagRepo repository.ProductProductTagDB
}

func NewProductTagService(hashtagRepo repository.ProductTag) *ProductTagService {
	return &ProductTagService{productTagRepo: hashtagRepo}
}

func (s *ProductTagService) GetAll(ctx context.Context) ([]model.ProductTag, error) {
	return s.productTagRepo.GetAll(ctx)
}

func (s *ProductTagService) Create(ctx context.Context, data admin.CreateProductTag) (model.ProductTag, error) {
	return s.productTagRepo.Create(ctx, model.ProductTag{
		Name:   data.Name,
		NameKz: data.NameKz,
		NameRu: data.NameEn,
	})
}

func (s *ProductTagService) Delete(ctx context.Context, productTagID uint) error {
	return s.productTagRepo.Delete(ctx, productTagID)
}

func (s *ProductTagService) Update(ctx context.Context, data admin.UpdateProductTag) (model.ProductTag, error) {
	productTag, err := s.productTagRepo.GetByID(ctx, data.ProductTagID)
	if err != nil {
		return model.ProductTag{}, err
	}
	if data.Name != nil {
		productTag.Name = *data.Name
	}
	if data.NameEn != nil {
		productTag.NameRu = *data.NameEn
	}
	if data.NameKz != nil {
		productTag.NameKz = *data.NameKz
	}
	return s.productTagRepo.Update(ctx, productTag)
}

func (s *ProductTagService) AddToProduct(ctx context.Context, data admin.AddProductProductTag) (model.ProductTag, error) {
	_, err := s.productProductTagRepo.Create(ctx, model.ProductProductTag{
		ProductTagID: data.ProductTagID,
		ProductID:    data.ProductID,
	})

	return model.ProductTag{}, err
}

func (s *ProductTagService) DeleteToProduct(ctx context.Context, data admin.DeleteProductProductTag) (model.ProductTag, error) {
	err := s.productProductTagRepo.DeleteByProductAndTagId(ctx, data.ProductID, data.ProductTagID)
	return model.ProductTag{}, err
}
