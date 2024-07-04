package model

import (
	"fmt"
	"time"
)

type SubmissionMark struct {
	UpdatedAt        time.Time `gorm:"autoUpdateTime;column:updated_at" json:"updated_at"`
	CreatedAt        time.Time `gorm:"autoCreateTime;column:created_at" json:"created_at"`
	SubmissionMarkID uint      `gorm:"primaryKey;column:submission_mark_id" json:"id"`
	SubmissionID     uint      `gorm:"column:submission_id" json:"submission_id"`
	Feedback         *string   `gorm:"column:feedback" json:"feedback"`
	MarkStatus       string    `gorm:"column:status" json:"status"`
	MarkType         string    `gorm:"column:submission_mark_type" json:"type"`
	TeacherID        *uint     `gorm:"column:teacher_id" json:"teacher_id"`

	Submission Submission
	Teacher    User
}

func (u SubmissionMark) TableName() string {
	return fmt.Sprintf("student_repo.submission_mark")
}

type SubmissionMarkStatus string

const (
	SUBMISSION_MARK_STATUS_NAME_ACCEPTED   AnswerStatusName = "accepted"
	SUBMISSION_MARK_STATUS_NAME_DECLINED   AnswerStatusName = "declined"
	SUBMISSION_MARK_STATUS_NAME_IN_PROCESS AnswerStatusName = "in_process"
)

type SubmissionMarkType string

const (
	SUBMISSION_MARK_TYPE_NAME_ENGLISH          AnswerStatusName = "English"
	SUBMISSION_MARK_TYPE_NAME_ECONOMICS        AnswerStatusName = "Economics"
	SUBMISSION_MARK_TYPE_NAME_STANDARD_CONTROL AnswerStatusName = "Standard Control"
	SUBMISSION_MARK_TYPE_NAME_ANTI_PLAGIARISM  AnswerStatusName = "Anti-Plagiarism"
	SUBMISSION_MARK_TYPE_NAME_FINAL            AnswerStatusName = "Final"
)
