package model

import "fmt"

type TeamUsers struct {
	TeamUsersID uint `gorm:"primaryKey;column:team_user_id" json:"id"`
	TeamID      uint `gorm:"column:team_id" json:"team_id"`
	UserID      uint `gorm:"column:student_id" json:"user_id"`

	Team Team
	User User
}

func (u TeamUsers) TableName() string {
	return fmt.Sprintf("student_repo.team_users")
}
