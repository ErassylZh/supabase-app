package model

import "time"

type Profile struct {
	UpdatedAt   time.Time `gorm:"autoUpdateTime;column:updated_at" json:"updated_at"`
	CreatedAt   time.Time `gorm:"autoCreateTime;column:created_at" json:"created_at"`
	ProfileID   uint      `gorm:"primaryKey;column:profile_id" json:"id"`
	ProfileName string    `gorm:"column:profile_name" json:"name"`
}

func (o *Profile) TableName() string {
	return "student_repo.profile"
}
