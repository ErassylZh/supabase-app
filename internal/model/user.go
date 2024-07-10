package model

import (
	"fmt"
	"time"
)

type User struct {
	UserID             string     `json:"id"`
	Email              string     `json:"email"`
	EncryptedPassword  *string    `json:"encrypted_password"`
	EmailConfirmedAt   *time.Time `json:"email_confirmed_at"`
	ConfirmationToken  *string    `json:"confirmation_token"`
	ConfirmationSendAt *time.Time `json:"confirmation_send_at"`
}

func (u User) TableName() string {
	return fmt.Sprintf("auth.users")
}
