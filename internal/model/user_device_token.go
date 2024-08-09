package model

import "fmt"

type UserDeviceToken struct {
	UserDeviceTokenID uint   `gorm:"primaryKey;column:user_device_token_id" json:"user_device_token_id"`
	UserID            string `gorm:"column:user_id" json:"user_id"`
	DeviceToken       string `gorm:"column:device_token" json:"device_token"`
}

func (u UserDeviceToken) TableName() string {
	return fmt.Sprintf("public.user_device_token")
}
