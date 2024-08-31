package model

import (
	"fmt"
	"time"
)

type TransactionType string

const (
	TRANSACTION_TYPE_INCOME TransactionType = "income"
	TRANSACTION_TYPE_SPEND  TransactionType = "spend"
)

type CurrencyType string

const (
	CURRENCY_TYPE_COIN     CurrencyType = "coin"
	CURRENCY_TYPE_SAPPHIRE CurrencyType = "sapphire"
)

type TransactionReason string

const (
	TRANSACTION_REASON_REFERRAL TransactionReason = "referral"
	TRANSACTION_REASON_STORE    TransactionReason = "store"
	TRANSACTION_REASON_POST     TransactionReason = "post"
)

type Transaction struct {
	CreatedAt         time.Time `gorm:"column:created_at" json:"created_at"`
	TransactionId     uint      `gorm:"primaryKey;column:transaction_id" json:"transaction_id"`
	UserId            string    `gorm:"column:user_id" json:"user_id"`
	TransactionType   string    `gorm:"column:transaction_type" json:"transaction_type"`
	Coins             int       `gorm:"column:coins" json:"coins"`
	Sapphires         int       `gorm:"column:sapphires" json:"sapphires"`
	TransactionReason string    `gorm:"column:reason" json:"reason"`
}

func (u Transaction) TableName() string {
	return fmt.Sprintf("public.transaction")
}
