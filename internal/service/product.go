package service

import (
	"context"
	"work-project/internal/model"
	"work-project/internal/repository"
)

type Product interface {
	GetListing(ctx context.Context, productTagIds []uint) ([]model.Product, error)
	GetProduct(ctx context.Context, id uint) (model.Product, error)
}

type ProductService struct {
	productRepo repository.Product
}

func NewProductService(productRepo repository.Product) *ProductService {
	return &ProductService{productRepo: productRepo}
}

func (s *ProductService) GetListing(ctx context.Context, productTagIds []uint) ([]model.Product, error) {
	return s.productRepo.GetAllListing(ctx, productTagIds)
}

func (s *ProductService) GetProduct(ctx context.Context, id uint) (model.Product, error) {
	return s.productRepo.GetById(ctx, id)
}
