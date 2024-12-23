package model

import (
	"fmt"
	"time"
)

type BucketName string

const (
	BUCKET_NAME_PRODUCT     BucketName = "product"
	BUCKET_NAME_POST        BucketName = "post"
	BUCKET_NAME_STORIES     BucketName = "stories"
	BUCKET_NAME_HASHTAG     BucketName = "hashtags"
	BUCKET_NAME_COLLECTION  BucketName = "collection"
	BUCKET_NAME_PRODUCT_TAG BucketName = "product_tag"
)

type Image struct {
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	ImageID   uint      `gorm:"primaryKey;column:image_id" json:"image_id"`
	FileName  string    `gorm:"column:file_name" json:"file_name"`
	ImageUrl  string    `gorm:"column:image_url" json:"url"`
	ProductID *uint     `gorm:"column:product_id" json:"product_id"`
	PostID    *uint     `gorm:"column:post_id" json:"post_id"`
	Type      string    `gorm:"column:type" json:"type"`
}

func (u Image) TableName() string {
	return fmt.Sprintf("public.image")
}
