package service

import (
	"context"
	"work-project/internal/model"
	"work-project/internal/repository"
)

type Product interface {
	GetListing(ctx context.Context) ([]model.Product, error)
}

type ProductService struct {
	productRepo repository.Product
}

func NewProductService(productRepo repository.Product) *ProductService {
	return &ProductService{productRepo: productRepo}
}

func (s *ProductService) GetListing(ctx context.Context) ([]model.Product, error) {
	return s.productRepo.GetAllListing(ctx)
}
