package worker

import (
	"context"
	"log"
	"work-project/internal/airtable"
	"work-project/internal/model"
	"work-project/internal/repository"
)

type AirTableSync struct {
	airTable repository.AirTable
	product  repository.Product
	storage  repository.StorageClient
	image    repository.Image
}

func NewAirTableSync(airTable repository.AirTable, product repository.Product, storage repository.StorageClient, image repository.Image) *AirTableSync {
	return &AirTableSync{airTable: airTable, product: product, storage: storage, image: image}
}

func (h *AirTableSync) Run() (err error) {
	ctx := context.Background()
	err = h.syncProducts(ctx)
	if err != nil {
		log.Println("error while sync products", "err", err)
		return err
	}

	log.Println("airtable sync end successful")
	return
}

func (h *AirTableSync) syncProducts(ctx context.Context) error {
	products, err := h.airTable.GetProducts(ctx)
	if err != nil {
		return err
	}
	productsAirtableBySku := make(map[string]airtable.BaseObject[airtable.ProductListResponse])
	for _, product := range products {
		productsAirtableBySku[product.Fields.Sku] = product
	}

	productsDb, err := h.product.GetAll(ctx)
	if err != nil {
		return err
	}
	productsDbBySku := make(map[string]model.Product)
	for _, product := range productsDb {
		productsDbBySku[product.Sku] = product
	}

	newProducts := make([]model.Product, 0)
	updateProducts := make([]model.Product, 0)
	for sku := range productsAirtableBySku {
		if product, exists := productsDbBySku[sku]; exists {
			if product.Point != productsAirtableBySku[sku].Fields.Point ||
				product.Count != productsAirtableBySku[sku].Fields.Count ||
				product.Description != productsAirtableBySku[sku].Fields.Description ||
				product.Title != productsAirtableBySku[sku].Fields.Title ||
				product.Sapphire != productsAirtableBySku[sku].Fields.Sapphire ||
				product.SellType != productsAirtableBySku[sku].Fields.SellType ||
				product.ProductType != productsAirtableBySku[sku].Fields.ProductType ||
				product.Status != productsAirtableBySku[sku].Fields.Status {

				product.Point = productsAirtableBySku[sku].Fields.Point
				product.Sapphire = productsAirtableBySku[sku].Fields.Sapphire
				product.Count = productsAirtableBySku[sku].Fields.Count
				product.Description = productsAirtableBySku[sku].Fields.Description
				product.Title = productsAirtableBySku[sku].Fields.Title
				product.SellType = productsAirtableBySku[sku].Fields.SellType
				product.ProductType = productsAirtableBySku[sku].Fields.ProductType
				product.Status = productsAirtableBySku[sku].Fields.Status
				updateProducts = append(updateProducts, product)
			}
			continue
		}
		newProducts = append(newProducts, model.Product{
			Title:             productsAirtableBySku[sku].Fields.Title,
			Sku:               productsAirtableBySku[sku].Fields.Sku,
			Description:       productsAirtableBySku[sku].Fields.Description,
			Point:             productsAirtableBySku[sku].Fields.Point,
			Sapphire:          productsAirtableBySku[sku].Fields.Sapphire,
			Status:            productsAirtableBySku[sku].Fields.Status,
			Count:             productsAirtableBySku[sku].Fields.Count,
			AirtableProductId: productsAirtableBySku[sku].Id,
			SellType:          productsAirtableBySku[sku].Fields.SellType,
			ProductType:       productsAirtableBySku[sku].Fields.ProductType,
		})
	}
	if len(newProducts) > 0 {
		newProducts, err = h.product.CreateMany(ctx, newProducts)
		if err != nil {
			log.Println(ctx, "error while create new products from airtable ", "err", err)
			return err
		}

		imagesProduct := make([]model.Image, 0)
		for _, np := range newProducts {
			productId := np.ProductID
			for _, img := range productsAirtableBySku[np.Sku].Fields.Image {
				file, err := h.storage.CreateImage(ctx, string(model.BUCKET_NAME_PRODUCT), img.FileName, img.Url)
				if err != nil {
					log.Println(ctx, "some err while create image", "err", err, "pr name", np.Title)
					return err
				}
				log.Println(ctx, "file for "+np.Title+" saved")
				imagesProduct = append(imagesProduct, model.Image{
					ProductID: &productId,
					ImageUrl:  file,
					FileName:  img.FileName,
				})
			}
		}
		_, err = h.image.CreateMany(ctx, imagesProduct)
		if err != nil {
			log.Println(ctx, "error while create images from airtable ", "err", err)
			return err
		}
	}

	if len(updateProducts) > 0 {
		_, err = h.product.UpdateMany(ctx, updateProducts)
		if err != nil {
			log.Println(ctx, "error while update exist products from airtable ", "err", err)
			return err
		}
	}
	return nil
}
