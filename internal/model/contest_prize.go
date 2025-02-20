package model

import (
	"fmt"
	"time"
)

type ContestPrize struct {
	CreatedAt      time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt      time.Time `gorm:"column:updated_at" json:"updated_at"`
	ContestPrizeID uint      `gorm:"primaryKey;column:contest_prize_id" json:"contest_prize_id"`
	ContestID      uint      `gorm:"column:contest_id" json:"contest_id"`
	Number         int       `gorm:"column:number" json:"number"`
	PrizeName      string    `gorm:"column:prize_name" json:"prize_name"`
	PhotoPath      *string   `gorm:"column:photo_path" json:"photo_path"`
}

func (c *ContestPrize) TableName() string {
	return fmt.Sprintf("public.contest_prize")
}
