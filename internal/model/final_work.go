package model

import "time"

type FinalWork struct {
	UpdatedAt   time.Time `gorm:"autoUpdateTime;column:updated_at" json:"updated_at"`
	CreatedAt   time.Time `gorm:"autoCreateTime;column:created_at" json:"created_at"`
	FinalWorkID uint      `gorm:"primaryKey;column:final_work_id" json:"id"`
	TeamID      uint      `gorm:"column:team_id" json:"team_id"`
	File        string    `gorm:"column:file" json:"file"`
	Mark        *int      `gorm:"column:mark" json:"mark"`

	Team Team ` json:"team"`
}

func (o *FinalWork) TableName() string {
	return "student_repo.final_work"
}
