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

	User User `json:"user"`
}

func (h *ContestParticipant) TableName() string {
	return fmt.Sprintf("public.contest_participant")
}
