package model

import (
	"fmt"
	"time"
)

type OrderPayment struct {
	CreatedAt      time.Time  `gorm:"column:created_at" json:"created_at"`
	UpdatedAt      time.Time  `gorm:"column:updated_at" json:"updated_at"`
	OrderPaymentId uint       `gorm:"primaryKey;column:order_payment_id" json:"order_payment_id"`
	Coins          int        `gorm:"column:coins" json:"coins"`
	Sapphires      int        `gorm:"column:sapphires" json:"sapphires"`
	PaymentStatus  int        `gorm:"column:payment_status" json:"payment_status"`
	PayedAt        *time.Time `gorm:"column:payed_at" json:"payed_at"`
}

func (d OrderPayment) TableName() string {
	return fmt.Sprintf("%s.%s", "public", "order_payment")
}
