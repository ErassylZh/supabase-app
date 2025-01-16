package model

import (
	"fmt"
	"time"
)

type User struct {
	UserID             string     `json:"id" gorm:"column:id"`
	Email              string     `json:"email" gorm:"column:email"`
	EncryptedPassword  *string    `json:"encrypted_password" gorm:"column:encrypted_password"`
	EmailConfirmedAt   *time.Time `json:"email_confirmed_at" gorm:"column:email_confirmed_at"`
	ConfirmationToken  *string    `json:"confirmation_token" gorm:"column:confirmation_token"`
	ConfirmationSendAt *time.Time `json:"confirmation_send_at" gorm:"column:confirmation_send_at"`

	Profile Profile `json:"profile" gorm:"foreignKey:ID;references:UserID"`
}

func (u User) TableName() string {
	return fmt.Sprintf("auth.users")
}
