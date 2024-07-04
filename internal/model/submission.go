package model

import (
	"fmt"
	"time"
)

type Submission struct {
	UpdatedAt    time.Time `gorm:"autoUpdateTime;column:updated_at" json:"updated_at"`
	CreatedAt    time.Time `gorm:"autoCreateTime;column:created_at" json:"created_at"`
	SubmissionID uint      `gorm:"primaryKey;column:submission_id" json:"id"`
	TeamID       uint      `gorm:"column:team_id" json:"team_id"`
	File         string    `gorm:"column:file" json:"file"`

	Team            Team             `json:"team"`
	SubmissionMarks []SubmissionMark `json:"marks"`
}

func (u Submission) TableName() string {
	return fmt.Sprintf("student_repo.submission")
}

type SubmissionAnswerStatus string

const (
	SUBMISSION_STATUS_NAME_ACCEPTED   AnswerStatusName = "accepted"
	SUBMISSION_STATUS_NAME_DECLINED   AnswerStatusName = "declined"
	SUBMISSION_STATUS_NAME_IN_PROCESS AnswerStatusName = "in_process"
)
