package model

import (
	"fmt"
)

type ProductTag struct {
	ProductTagID uint    `gorm:"primaryKey;column:product_tag_id" json:"product_tag_id"`
	Name         string  `gorm:"column:name" json:"name"`
	NameRu       string  `gorm:"column:name_ru" json:"name_ru"`
	NameKz       string  `gorm:"column:name_kz" json:"name_kz"`
	ImagePath    *string `gorm:"column:image_path" json:"image_path"`

	//Posts []Post `gorm:"many2many:public.post_hashtag;foreignKey:HashtagID;joinForeignKey:HashtagID;References:PostID;joinReferences:PostID;constraint:OnDelete:CASCADE;" json:"posts"`
}

func (h *ProductTag) TableName() string {
	return fmt.Sprintf("public.product_tag")
}
