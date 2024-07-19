package model

import (
	"fmt"
	"time"
)

type Referral struct {
	CreatedAt     time.Time `gorm:"column:created_at" json:"created_at"`
	ReferralID    uint      `gorm:"column:referral_id" json:"referral_id"`
	ReferralCode  string    `gorm:"referral_code" json:"referral_code"`
	UserID        string    `gorm:"column:user_id" json:"user_id"`
	InvitedUserID string    `gorm:"column:invited_user_id" json:"invited_user_id"`
}

func (u Referral) TableName() string {
	return fmt.Sprintf("public.referral")
}
