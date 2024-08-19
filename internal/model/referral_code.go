package model

import (
	"fmt"
	"time"
)

type ReferralCode struct {
	CreatedAt      time.Time `gorm:"column:created_at" json:"created_at"`
	ReferralCodeID uint      `gorm:"primaryKey;column:referral_code_id" json:"referral_code_id"`
	ReferralCode   string    `gorm:"column:referral_code" json:"referral_code"`
	UserID         string    `gorm:"column:user_id" json:"user_id"`
}

func (u ReferralCode) TableName() string {
	return fmt.Sprintf("public.referral_code")
}
