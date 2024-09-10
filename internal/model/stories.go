package model

import (
	"fmt"
	"time"
)

type Stories struct {
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	StoriesId uint      `gorm:"primaryKey;column:stories_id" json:"stories_id"`
	StartTime time.Time `gorm:"column:start_time" json:"start_time"`
	EndTime   time.Time `gorm:"column:end_time" json:"end_time"`
	Title     string    `gorm:"column:title" json:"title"`
	IconPath  string    `gorm:"column:icon_path" json:"icon_path"`

	StoryPages []StoryPage `json:"story_pages"`
}

func (u *Stories) TableName() string {
	return fmt.Sprintf("public.stories")
}
