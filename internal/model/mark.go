package model

import (
	"fmt"
	"time"
)

type Mark struct {
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	MarkID    uint      `gorm:"primaryKey;column:mark_id" json:"mark_id"`
	UserID    uint      `gorm:"not null" json:"user_id"`
	PostID    uint      `gorm:"not null" json:"post_id"`
	User      User      `gorm:"foreignKey:UserID;references:UserID" json:"user"`
	Post      Post      `gorm:"foreignKey:PostID;references:PostID" json:"post"`
}

func (m Mark) TableName() string {
	return fmt.Sprintf("public.mark")
}
