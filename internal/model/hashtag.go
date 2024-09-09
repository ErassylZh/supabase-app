package model

import (
	"fmt"
)

type HashtagName string

const (
	HASHTAG_NAME_BESTSELLER HashtagName = "bestseller"
	HASHTAG_NAME_PARTNER    HashtagName = "partner"
)

type Hashtag struct {
	HashtagID uint    `gorm:"primaryKey;column:hashtag_id" json:"hashtag_id"`
	Name      string  `gorm:"column:name" json:"name"`
	NameRu    string  `gorm:"column:name_ru" json:"name_ru"`
	NameKz    string  `gorm:"column:name_kz" json:"name_kz"`
	ImagePath *string `gorm:"column:image_path" json:"image_path"`
}

func (h *Hashtag) TableName() string {
	return fmt.Sprintf("public.hashtag")
}
