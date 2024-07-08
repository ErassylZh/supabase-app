package model

import (
	"fmt"
)

type Profile struct {
	ProfileID   string `json:"profile_id"`
	ProfileName string `json:"profile_name"`
}

func (u Profile) TableName() string {
	return fmt.Sprintf("public.profile")
}
