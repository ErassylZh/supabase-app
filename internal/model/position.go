package model

import (
	"fmt"
	"time"
)

type Position struct {
	UpdatedAt    time.Time `gorm:"autoUpdateTime;column:updated_at" json:"updated_at"`
	CreatedAt    time.Time `gorm:"autoCreateTime;column:created_at" json:"created_at"`
	PositionID   uint      `gorm:"primaryKey;column:position_id" json:"id"`
	PositionName string    `gorm:"column:position_name" json:"name"`
	RoleID       uint      `gorm:"column:role_id" json:"role_id"`
	RoleName     string    `gorm:"-" json:"role"`

	Role Role `json:"-"`
}

func (u Position) TableName() string {
	return fmt.Sprintf("auth.position")
}

type PositionCodeName string

const (
	POSITION_CODE_NAME_ENGLISH          PositionCodeName = "English"
	POSITION_CODE_NAME_EСONOMICS        PositionCodeName = "Economics"
	POSITION_CODE_NAME_ANTI_PLAGIARISM  PositionCodeName = "Anti-Plagiarism"
	POSITION_CODE_NAME_STANDARD_CONTROL PositionCodeName = "Standard Control"
)
