package model

import (
	"fmt"
	"time"
)

type ContestHistory struct {
	CreatedAt        time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt        time.Time `gorm:"column:updated_at" json:"updated_at"`
	ContestHistoryID uint      `gorm:"primaryKey;column:contest_history_id" json:"contest_History_id"`
	ContestID        uint      `gorm:"column:contest_id" json:"contest_id"`
	UserID           string    `gorm:"column:user_id" json:"user_id"`
	Points           int       `gorm:"column:points" json:"points"`
	ReadTime         int       `gorm:"column:read_time" json:"read_time"`
	ContestBookID    uint      `gorm:"column:contest_book_id" json:"contest_book_id"`
}

func (h *ContestHistory) TableName() string {
	return fmt.Sprintf("public.contest_history")
}
