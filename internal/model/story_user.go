package model

import (
	"fmt"
	"time"
)

type StoryPageUser struct {
	CreatedAt       time.Time `gorm:"column:created_at" json:"created_at"`
	StoryPageUserId uint      `gorm:"primaryKey;column:story_page_user_id" json:"story_page_user_id"`
	StoryPageId     uint      `gorm:"column:story_page_id" json:"story_page_id"`
	UserId          string    `gorm:"column:user_id" json:"user_id"`
}

func (u *StoryPageUser) TableName() string {
	return fmt.Sprintf("public.story_page_user")
}
