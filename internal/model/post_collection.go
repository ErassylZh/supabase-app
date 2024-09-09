package model

import "fmt"

type PostCollection struct {
	PostId       uint `gorm:"column:post_id" json:"post_id"`
	CollectionId uint `gorm:"column:collection_id" json:"collection_id"`
}

func (p PostCollection) TableName() string {
	return fmt.Sprintf("public.post_collection")
}
