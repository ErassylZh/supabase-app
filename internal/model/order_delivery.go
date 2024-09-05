package model

import "time"

type OrderDelivery struct {
	CreatedAt       time.Time
	UpdatedAt       time.Time
	OrderDeliveryId uint    `json:"order_delivery_id"`
	DeliveryAddress *string `json:"delivery_address"`
	DeliveryType    uint    `json:"delivery_type"`
}
