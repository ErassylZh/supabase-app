package model

import (
	"fmt"
	"time"
)

type ContestBook struct {
	CreatedAt     time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt     time.Time `gorm:"column:updated_at" json:"updated_at"`
	ContestBookID uint      `gorm:"primaryKey;column:contest_book_id" json:"contest_book_id"`
	ContestID     uint      `gorm:"column:contest_id" json:"contest_id"`
	PostID        uint      `gorm:"column:post_id" json:"post_id"`
	DayNumber     int       `gorm:"column:day_number" json:"day_number"`
	ContestCoins  int       `gorm:"column:contest_coins" json:"contest_coins"`
	Point         int       `gorm:"column:point" json:"point"`
}

func (h *ContestBook) TableName() string {
	return fmt.Sprintf("public.contest_book")
}
