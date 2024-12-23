package model

import "fmt"

type ProductProductTag struct {
	ProductId    uint `gorm:"column:product_id" json:"product_id"`
	ProductTagId uint `gorm:"column:product_tag_id" json:"product_tag_id"`
}

func (p ProductProductTag) TableName() string {
	return fmt.Sprintf("public.product_product_tag")
}
