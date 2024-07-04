package model

import "time"

type Invitation struct {
	UpdatedAt     time.Time `gorm:"autoUpdateTime;column:updated_at" json:"updated_at"`
	CreatedAt     time.Time `gorm:"autoCreateTime;column:created_at" json:"created_at"`
	InvitationID  uint      `gorm:"primaryKey;column:invitation_id" json:"id"`
	TeamID        uint      `gorm:"column:team_id" json:"team_id"`
	InvitedUserId uint      `gorm:"column:invited_user_id" json:"invited_user_id"`
	AnswerStatus  string    `gorm:"column:answer_status" json:"answer_status"`
	Comment       *string   `gorm:"column:comment" json:"comment"`

	Team        Team `json:"team"`
	InvitedUser User `json:"-"`
}

func (o *Invitation) TableName() string {
	return "student_repo.invitation"
}

type AnswerStatusName string

const (
	INVITATION_ANSWER_STATUS_NAME_ACCEPTED   AnswerStatusName = "accepted"
	INVITATION_ANSWER_STATUS_NAME_DECLINED   AnswerStatusName = "declined"
	INVITATION_ANSWER_STATUS_NAME_IN_PROCESS AnswerStatusName = "in_process"
)
