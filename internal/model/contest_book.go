package model

import (
	"fmt"
	"time"
	"work-project/internal/schema"
)

type ContestBook struct {
	CreatedAt     time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt     time.Time `gorm:"column:updated_at" json:"updated_at"`
	ContestBookID uint      `gorm:"primaryKey;column:contest_book_id" json:"contest_book_id"`
	ContestID     uint      `gorm:"column:contest_id" json:"contest_id"`
	DayNumber     int       `gorm:"column:day_number" json:"day_number"`
	ContestCoins  int       `gorm:"column:contest_coins" json:"contest_coins"`
	Point         int       `gorm:"column:point" json:"point"`
	Title         string    `gorm:"column:title" json:"title"`
	Description   string    `gorm:"column:description" json:"description"`
	Status        string    `gorm:"column:status" json:"status"`
	Body          string    `gorm:"column:body" json:"body"`

	ContestHistory []ContestHistory `json:"contest_history"`
}

func (c *ContestBook) TableName() string {
	return fmt.Sprintf("public.contest_book")
}

type ContestBooks []ContestBook

func (c ContestBooks) GetContestBookSchema(userId string) []schema.ContestBookData {
	result := make([]schema.ContestBookData, 0)
	for _, cb := range c {
		data := schema.ContestBookData{
			ContestBook: cb,
		}
		for _, ch := range cb.ContestHistory {
			if ch.UserID == userId {
				data.Readed = true
				break
			}
		}
		result = append(result, data)
	}
	return result
}
