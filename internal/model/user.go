package model

import (
	"fmt"
)

type User struct {
	UserID string `json:"id"`
	Email  string `json:"email"`
}

func (u User) TableName() string {
	return fmt.Sprintf("auth.users")
}
