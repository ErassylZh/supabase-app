package model

import "time"

type GrammarCheckerModel struct {
	UpdatedAt             time.Time `gorm:"autoUpdateTime;column:updated_at" json:"updated_at"`
	CreatedAt             time.Time `gorm:"autoCreateTime;column:created_at" json:"created_at"`
	GrammarCheckerModelID uint      `gorm:"primaryKey;column:grammar_checker_id" json:"id"`
	SubmissionMarkID      uint      `gorm:"column:submission_mark_id" json:"submission_mark_id"`
	Text                  string    `gorm:"column:text" json:"text"`
	Status                int       `gorm:"column:status;default:-1" json:"status"`
	Details               *string   `gorm:"column:details" json:"details"`
}

func (o *GrammarCheckerModel) TableName() string {
	return "student_repo.grammar_checker"
}
