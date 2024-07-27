package model

import "fmt"

type Balance struct {
	BalanceId uint   `gorm:"primaryKey;column:balance_id" json:"balance_id"`
	UserId    string `gorm:"column:user_id" json:"user_id"`
	Coins     int    `gorm:"column:coins" json:"coins"`
	Sapphires int    `gorm:"column:sapphires" json:"sapphires"`
}

func (u Balance) TableName() string {
	return fmt.Sprintf("public.balance")
}
