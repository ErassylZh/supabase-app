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
}

func (h *Contest) TableName() string {
	return fmt.Sprintf("public.contest")
}
