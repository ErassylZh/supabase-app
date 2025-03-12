package service

import (
	"context"
	"log"
	"time"
	"work-project/internal/admin"
	"work-project/internal/model"
	"work-project/internal/repository"
)

type Product interface {
	GetListing(ctx context.Context, productTagIds []uint) ([]model.Product, error)
	GetProduct(ctx context.Context, id uint) (model.Product, error)
	Create(ctx context.Context, data admin.CreateProduct) (model.Product, error)
	Update(ctx context.Context, data admin.UpdateProduct) (model.Product, error)
	Delete(ctx context.Context, id uint) error
	GetById(ctx context.Context, id uint) (model.Product, error)
}

type ProductService struct {
	productRepo repository.Product
	image       repository.Image
	storage     repository.Storage
}

func NewProductService(productRepo repository.Product, image repository.Image, storage repository.Storage) *ProductService {
	return &ProductService{productRepo: productRepo, image: image, storage: storage}
}

func (s *ProductService) GetListing(ctx context.Context, productTagIds []uint) ([]model.Product, error) {
	return s.productRepo.GetAllListing(ctx, productTagIds)
}

func (s *ProductService) GetProduct(ctx context.Context, id uint) (model.Product, error) {
	return s.productRepo.GetById(ctx, id)
}

func (s *ProductService) Create(ctx context.Context, data admin.CreateProduct) (model.Product, error) {
	product := model.Product{
		Title:         data.Title,
		TitleEn:       data.TitleEn,
		TitleKz:       data.TitleKz,
		Description:   data.Description,
		DescriptionKz: data.DescriptionKz,
		DescriptionEn: data.DescriptionEn,
		Status:        string(model.PRODUCT_STATUS_DRAFT),
		Count:         data.Count,
		Point:         data.Point,
		Sapphire:      data.Sapphire,
		Sku:           data.Sku,
		ProductType:   data.ProductType,
		SellType:      data.SellType,
		Offer:         data.Offer,
		OfferEn:       data.OfferEn,
		OfferKz:       data.OfferKz,
		Discount:      data.Discount,
		DiscountKz:    data.DiscountKz,
		DiscountEn:    data.DiscountEn,
		Contacts:      data.Contacts,
		ContactsKz:    data.ContactsKz,
		ContactsEn:    data.ContactsEn,
	}
	product, err := s.productRepo.Create(ctx, product)
	if err != nil {
		return model.Product{}, err
	}

	images := make([]model.Image, 0)

	if data.Logo != nil {
		file, err := s.storage.CreateImageFromBase64(ctx, string(model.BUCKET_NAME_PRODUCT), time.Now().String()+data.Logo.FileName, data.Logo.File)
		if err != nil {
			log.Println(ctx, "some err while create image", "err", err, "product name", product.Title)
			return model.Product{}, err
		}
		images = append(images, model.Image{
			ProductID: &product.ProductID,
			ImageUrl:  file,
			Type:      string(model.POST_IMAGE_TYPE_LOGO),
		})
	}

	if data.Logo != nil {
		file, err := s.storage.CreateImageFromBase64(ctx, string(model.BUCKET_NAME_PRODUCT), time.Now().String()+data.Logo.FileName, data.Logo.File)
		if err != nil {
			log.Println(ctx, "some err while create image", "err", err, "product name", product.Title)
			return model.Product{}, err
		}
		images = append(images, model.Image{
			ProductID: &product.ProductID,
			ImageUrl:  file,
			Type:      string(model.POST_IMAGE_TYPE_IMAGE),
		})
	}

	if len(images) > 0 {
		_, err = s.image.CreateMany(ctx, images)
		return model.Product{}, err
	}

	return product, nil
}

func (s *ProductService) Update(ctx context.Context, data admin.UpdateProduct) (model.Product, error) {
	product, err := s.productRepo.GetById(ctx, data.ProductID)
	if err != nil {
		return model.Product{}, err
	}
	images := make([]model.Image, 0)

	if data.Logo != nil {
		file, err := s.storage.CreateImageFromBase64(ctx, string(model.BUCKET_NAME_PRODUCT), time.Now().String()+data.Logo.FileName, data.Logo.File)
		if err != nil {
			log.Println(ctx, "some err while create image", "err", err, "product name", product.Title)
			return model.Product{}, err
		}
		images = append(images, model.Image{
			ProductID: &product.ProductID,
			ImageUrl:  file,
			Type:      string(model.POST_IMAGE_TYPE_LOGO),
		})
	}

	if data.Logo != nil {
		file, err := s.storage.CreateImageFromBase64(ctx, string(model.BUCKET_NAME_PRODUCT), time.Now().String()+data.Logo.FileName, data.Logo.File)
		if err != nil {
			log.Println(ctx, "some err while create image", "err", err, "product name", product.Title)
			return model.Product{}, err
		}
		images = append(images, model.Image{
			ProductID: &product.ProductID,
			ImageUrl:  file,
			Type:      string(model.POST_IMAGE_TYPE_IMAGE),
		})
	}

	if len(images) > 0 {
		_, err = s.image.CreateMany(ctx, images)
		return model.Product{}, err
	}

	return s.productRepo.Update(ctx, product)
}

func (s *ProductService) Delete(ctx context.Context, id uint) error {
	return s.productRepo.DeleteById(ctx, id)
}

func (s *ProductService) GetById(ctx context.Context, id uint) (model.Product, error) {
	return s.productRepo.GetById(ctx, id)
}
