package model

import "time"

type AntiPlagiarismModel struct {
	UpdatedAt        time.Time `gorm:"autoUpdateTime;column:updated_at" json:"updated_at"`
	CreatedAt        time.Time `gorm:"autoCreateTime;column:created_at" json:"created_at"`
	AntiPlagiarismID uint      `gorm:"primaryKey;column:anti_plagiarism_id" json:"id"`
	SubmissionMarkID uint      `gorm:"column:submission_mark_id" json:"submission_mark_id"`
	UID              string    `gorm:"column:uid" json:"uid"`
	UniquePercent    *float64  `gorm:"column:unique_percent" json:"unique_percent"`
	Details          *string   `gorm:"column:details" json:"details"`
}

func (o *AntiPlagiarismModel) TableName() string {
	return "student_repo.anti_plagiarism"
}
