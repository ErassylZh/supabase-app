package model

import (
	"fmt"
	"time"
)

type PushNotification struct {
	CreatedAt          time.Time `gorm:"column:created_at" json:"created_at"`
	PushNotificationID uint      `gorm:"column:push_notification_id" json:"push_notification_id"`
	Title              string    `gorm:"column:title" json:"title"`
	Text               string    `gorm:"column:text" json:"text"`
	Token              *string   `gorm:"column:token" json:"token"`
	Topic              *string   `gorm:"column:topic" json:"topic"`
	Condition          *string   `gorm:"column:condition" json:"condition"`
}

func (u PushNotification) TableName() string {
	return fmt.Sprintf("public.push_notification")
}
