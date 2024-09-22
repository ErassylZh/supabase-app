package model

import (
	"fmt"
	"time"
)

type UserPost struct {
	CreatedAt     time.Time `gorm:"column:created_at" json:"created_at"`
	UserPostId    uint      `gorm:"primaryKey;column:user_post_id" json:"user_post_id"`
	UserId        string    `gorm:"column:user_id" json:"user_id"`
	PostId        uint      `gorm:"column:post_id" json:"post_id"`
	QuizPoints    *int      `gorm:"column:quiz_points" json:"quiz_points"`
	QuizSapphires *int      `gorm:"column:quiz_sapphires" json:"quiz_sapphires"`
}

func (u *UserPost) TableName() string {
	return fmt.Sprintf("%s.%s", "public", "user_post")
}

type ReadPost struct {
	PostId uint `gorm:"column:post_id" json:"post_id"`
}
