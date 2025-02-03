package model

import (
	"fmt"
	"time"
)

type ContestParticipant struct {
	CreatedAt            time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt            time.Time `gorm:"column:updated_at" json:"updated_at"`
	ContestParticipantID uint      `gorm:"primaryKey;column:contest_participant_id" json:"contest_participant_id"`
	ContestID            uint      `gorm:"column:contest_id" json:"contest_id"`
	UserID               string    `gorm:"column:user_id" json:"user_id"`
	Points               int       `gorm:"column:points" json:"points"`
	ReadTime             int       `gorm:"column:read_time" json:"read_time"`
	Number               int       `gorm:"column:number" json:"number"`
	ContestPrizeId       uint      `gorm:"column:contest_prize_id"  json:"contest_prize_id"`
	PrizeGet             bool      `gorm:"column:prize_get" json:"prize_get"`

	User User `json:"user"`
}

func (h *ContestParticipant) TableName() string {
	return fmt.Sprintf("public.contest_participant")
}
