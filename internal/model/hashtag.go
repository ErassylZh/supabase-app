package model

import (
	"fmt"
)

type Hashtag struct {
	HashtagID uint   `gorm:"primaryKey;column:hashtag_id" json:"hashtag_id"`
	Name      string `gorm:"column:name" json:"name"`
}

func (h *Hashtag) TableName() string {
	return fmt.Sprintf("public.hashtag")
}
