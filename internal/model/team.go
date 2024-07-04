package model

import (
	"fmt"
	"time"
)

type Team struct {
	UpdatedAt    time.Time `gorm:"autoUpdateTime;column:updated_at" json:"updated_at"`
	CreatedAt    time.Time `gorm:"autoCreateTime;column:created_at" json:"created_at"`
	TeamID       uint      `gorm:"primaryKey;column:team_id" json:"id"`
	SupervisorID *uint     `gorm:"column:supervisor_id" json:"supervisor_id"`
	ProfileID    uint      `gorm:"column:profile_id" json:"profile_id"`
	Topic        string    `gorm:"column:topic" json:"topic"`

	Profile    Profile `json:"profile"`
	Supervisor *User   `json:"supervisor"`
	TeamUsers  []TeamUsers
}

func (u Team) TableName() string {
	return fmt.Sprintf("student_repo.team")
}
