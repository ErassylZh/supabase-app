package model

import (
	"fmt"
	"time"
)

type Profile struct {
	ProfileID string    `json:"id"`
	UpdatedAt time.Time `json:"updated_at"`
	FullName  *string   `json:"full_name"`
	AvatarUrl *string   `json:"avatar_url"`
	Phone     *string   `json:"phone"`
	Email     *string   `json:"email"`
}

func (u Profile) TableName() string {
	return fmt.Sprintf("public.profiles")
}
