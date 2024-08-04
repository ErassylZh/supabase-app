package model

import "time"

type Image struct {
	CreatedAt       time.Time `gorm:"column:created_at" json:"created_at"`
	ImageID         uint      `gorm:"column:image_id" json:"image_id"`
	AirtableImageID *string   `gorm:"column:airtable_image_id" json:"airtable_image_id"`
	FileName        string    `gorm:"column:file_name" json:"file_name"`
	ImageUrl        string    `gorm:"column:image_url" json:"url"`
	ProductID       *string   `gorm:"column:product_id" json:"product_id"`
	PostID          *string   `gorm:"column:post_id" json:"post_id"`
	Type            string    `gorm:"column:type" json:"type"`
}
