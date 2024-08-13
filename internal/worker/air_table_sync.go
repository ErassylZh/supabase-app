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
}

func NewAirTableSync(airTable repository.AirTable, product repository.Product) *AirTableSync {
	return &AirTableSync{airTable: airTable, product: product}
}

func (h *AirTableSync) Run() (err error) {
	ctx := context.Background()

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
		_, err = h.product.CreateMany(ctx, newProducts)
		if err != nil {
			log.Println(ctx, "error while create new products from airtable ", "err", err)
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
	log.Println("airtable sync end successful")
	return
}
