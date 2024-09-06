package model

import (
	"fmt"
	"time"
)

type StoryPage struct {
	CreatedAt   time.Time `gorm:"column:created_at" json:"created_at"`
	StoryPageId uint      `gorm:"primaryKey;column:story_page_id" json:"story_page_id"`
	ImagePath   string    `gorm:"column:image_path" json:"image_path"`
	StoriesId   uint      `gorm:"column:stories_id" json:"stories_id"`
	Text        string    `gorm:"column:text" json:"text"`
	PageOrder   int       `gorm:"column:page_order" json:"page_order"`
	Uuid        string    `gorm:"column:uuid" json:"uuid"`
	IsReaded    bool      `gorm:"-" json:"is_readed"`
}

func (u *StoryPage) TableName() string {
	return fmt.Sprintf("public.story_page")
}
