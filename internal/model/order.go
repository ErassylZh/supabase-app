package model

import "time"

type Order struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	OrderId   uint  `json:"order_id"`
	BuyerId   *uint `json:"buyer_id"`
}
