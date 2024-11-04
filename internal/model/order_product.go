package model

import (
	"fmt"
	"time"
)

type OrderProduct struct {
	CreatedAt      time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt      time.Time `gorm:"column:updated_at" json:"updated_at"`
	OrderProductID uint      `gorm:"primaryKey;column:order_product_id" json:"order_product_id"`
	Quantity       uint      `gorm:"column:quantity" json:"quantity"`
	Coins          int       `gorm:"column:coins" json:"coins"`
	Sapphires      int       `gorm:"column:sapphires" json:"sapphires"`
	Sku            string    `gorm:"column:sku" json:"sku"`
	OrderID        uint      `gorm:"column:order_id" json:"order_id"`
	ProductID      uint      `gorm:"column:product_id" json:"product_id"`
	ProductType    string    `gorm:"column:product_type" json:"product_type"`

	Product Product `json:"product"`
}

func (p OrderProduct) TableName() string {
	return fmt.Sprintf("%s.%s", "public", "order_product")
}
