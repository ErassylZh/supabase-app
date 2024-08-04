package model

import (
	"fmt"
	"time"
)

// airtable store
type Product struct {
	CreatedAt         time.Time
	ProductID         uint    `gorm:"column:product_id" json:"product_id"`
	AirtableProductId *string `gorm:"primaryKey;column:airtable_product_id" json:"id"`
	Title             string  `gorm:"column:title" json:"title"`
	Description       string  `gorm:"column:description" json:"description"`
	Count             int     `gorm:"column:count" json:"count"`
	Point             int     `gorm:"column:point" json:"point"`
	Status            string  `gorm:"column:status" json:"status"`
}

func (u Product) TableName() string {
	return fmt.Sprintf("public.product")
}
