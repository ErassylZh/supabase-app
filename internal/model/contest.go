package model

import (
	"fmt"
	"time"
)

type Contest struct {
	CreatedAt time.Time `gorm:"primaryKey;column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
	ContestID uint      `gorm:"column:contest_id" json:"contest_id"`
	StartTime time.Time `gorm:"column:start_time" json:"start_time"`
	EndTime   time.Time `gorm:"column:end_time" json:"end_time"`
	IsActive  bool      `gorm:"column:is_active" json:"is_active"`

	ContestParticipants []ContestParticipant `json:"contest_participants"`
	ContestBooks        []ContestBook        `json:"contest_books"`
}

func (c *Contest) TableName() string {
	return fmt.Sprintf("public.contest")
}

func (c *Contest) UserJoined(userID string) bool {
	for _, cp := range c.ContestParticipants {
		if cp.UserID == userID {
			return true
		}
	}
	return false
}

func (c *Contest) CurrentDayNumber() int {
	now := time.Now()

	if now.Before(c.StartTime) || now.After(c.EndTime) {
		return 0
	}

	duration := now.Sub(c.StartTime)
	dayNumber := int(duration.Hours()/24) + 1 // Добавляем 1, так как день начинается с 1
	return dayNumber
}
