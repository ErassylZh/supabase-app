package model

import (
	"fmt"
	"time"
)

type Order struct {
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
	OrderId   uint      `gorm:"primaryKey;column:order_id" json:"order_id"`
	BuyerId   *string   `gorm:"column:buyer_id" json:"buyer_id"`

	OrderProducts []OrderProduct `json:"order_products"`
}

func (d Order) TableName() string {
	return fmt.Sprintf("%s.%s", "public", "order")
}
