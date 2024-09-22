package model

import (
	"fmt"
)

type Collection struct {
	CollectionID     uint    `gorm:"primaryKey;column:collection_id" json:"collection_id"`
	Name             string  `gorm:"column:name" json:"name"`
	NameRu           string  `gorm:"column:name_ru" json:"name_ru"`
	NameKz           string  `gorm:"column:name_kz" json:"name_kz"`
	ImagePath        *string `gorm:"column:image_path" json:"image_path"`
	ImagePathKz      *string `gorm:"column:image_path_kz" json:"image_path_kz"`
	ImagePathRu      *string `gorm:"column:image_path_ru" json:"image_path_ru"`
	IsRecommendation bool    `gorm:"column:is_recommendation" json:"is_recommendation"`

	Posts []Post `gorm:"many2many:post_collection;foreignKey:CollectionID;joinForeignKey:CollectionID;References:PostID;joinReferences:PostID;constraint:OnDelete:CASCADE;" json:"posts"`
}

func (h *Collection) TableName() string {
	return fmt.Sprintf("public.collection")
}
