package model

import (
	"fmt"
)

type Collection struct {
	CollectionID uint    `gorm:"primaryKey;column:collection_id" json:"collection_id"`
	Name         string  `gorm:"column:name" json:"name"`
	NameRu       string  `gorm:"column:name_ru" json:"name_ru"`
	NameKz       string  `gorm:"column:name_kz" json:"name_kz"`
	ImagePath    *string `gorm:"column:image_path" json:"image_path"`
}

func (h *Collection) TableName() string {
	return fmt.Sprintf("public.collection")
}
