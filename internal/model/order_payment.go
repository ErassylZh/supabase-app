package model

import "time"

type OrderPayment struct {
	CreatedAt      time.Time
	UpdatedAt      time.Time
	OrderPaymentId uint
	Coins          int `json:"coins"`
	Sapphires      int `json:"sapphires"`
	PaymentStatus  int `json:"payment_status"`
	PayedAt        *time.Time
}
