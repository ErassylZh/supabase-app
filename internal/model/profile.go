package model

import (
	"fmt"
	"time"
)

type Profile struct {
	ID        string    `json:"id"`
	UpdatedAt time.Time `json:"updated_at"`
	FullName  *string   `json:"full_name"`
	AvatarUrl *string   `json:"avatar_url"`
	Phone     *string   `json:"phone"`
	Email     *string   `json:"email"`
	Nickname  *string   `json:"user_name"`
}

func (u Profile) TableName() string {
	return fmt.Sprintf("public.profiles")
}
