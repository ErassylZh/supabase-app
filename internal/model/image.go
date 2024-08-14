package model

import (
	"fmt"
	"time"
)

type BucketName string

const (
	BUCKET_NAME_PRODUCT BucketName = "product"
)

type Image struct {
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	ImageID   uint      `gorm:"column:image_id" json:"image_id"`
	FileName  string    `gorm:"column:file_name" json:"file_name"`
	ImageUrl  string    `gorm:"column:image_url" json:"url"`
	ProductID *uint     `gorm:"column:product_id" json:"product_id"`
	PostID    *uint     `gorm:"column:post_id" json:"post_id"`
	Type      string    `gorm:"column:type" json:"type"`
}

func (u Image) TableName() string {
	return fmt.Sprintf("public.image")
}
