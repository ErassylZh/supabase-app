package model

import (
	"fmt"
	"time"
)

type ProductStatus string

const (
	PRODUCT_STATUS_PUBLISH ProductStatus = "Publish"
	PRODUCT_STATUS_DRAFT   ProductStatus = "Draft"
)

type ProductType string

const (
	PRODUCT_TYPE_VIRTUAL ProductType = "virtual"
	PRODUCT_TYPE_PHYSIC  ProductType = "physic"
)

type SellType string

const (
	SELL_TYPE_STANDARD SellType = "standard"
	SELL_TYPE_AUCTION  SellType = "auction"
)

// airtable store
type Product struct {
	CreatedAt         time.Time `gorm:"column:created_at" json:"created_at"`
	ProductID         uint      `gorm:"primaryKey;column:product_id" json:"product_id"`
	AirtableProductId string    `gorm:"column:airtable_product_id" json:"airtable_product_id"`
	Title             string    `gorm:"column:title" json:"title"`
	Description       string    `gorm:"column:description" json:"description"`
	Count             int       `gorm:"column:count" json:"count"`
	Point             int       `gorm:"column:point" json:"point"`
	Sapphire          int       `gorm:"column:sapphire" json:"sapphire"`
	Status            string    `gorm:"column:status" json:"status"`
	Sku               string    `gorm:"column:sku" json:"sku"`
	ProductType       string    `gorm:"column:product_type" json:"product_type"`
	SellType          string    `gorm:"column:sell_type" json:"sell_type"`
	Offer             string    `gorm:"column:offer" json:"offer"`
	Discount          string    `gorm:"column:discount" json:"discount"`
	Contacts          string    `gorm:"column:contacts" json:"contacts"`

	Images      []Image      `json:"images"`
	ProductTags []ProductTag `gorm:"many2many:public.product_product_tag;foreignKey:ProductID;joinForeignKey:ProductID;References:ProductTagID;joinReferences:ProductTagID;constraint:OnDelete:CASCADE;" json:"product_tags"`
}

func (u Product) TableName() string {
	return fmt.Sprintf("public.product")
}
