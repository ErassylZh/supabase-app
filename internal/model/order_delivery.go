package model

import (
	"fmt"
	"time"
)

type OrderDelivery struct {
	CreatedAt       time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt       time.Time `gorm:"column:updated_at" json:"updated_at"`
	OrderDeliveryId uint      `gorm:"primaryKey;column:order_delivery_id" json:"order_delivery_id"`
	DeliveryAddress *string   `gorm:"column:delivery_address" json:"delivery_address"`
	DeliveryType    uint      `gorm:"column:delivery_type" json:"delivery_type"`
}

func (d OrderDelivery) TableName() string {
	return fmt.Sprintf("%s.%s", "public", "order_delivery")
}
