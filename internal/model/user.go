package model

import (
	"fmt"
	"time"
)

type User struct {
	UpdatedAt          time.Time  `gorm:"autoUpdateTime;column:updated_at" json:"updated_at"`
	CreatedAt          time.Time  `gorm:"autoCreateTime;column:created_at" json:"created_at"`
	UserID             uint       `gorm:"primaryKey;column:user_id" json:"id"`
	FullName           string     `gorm:"column:full_name;" json:"name"`
	Password           string     `gorm:"column:password" json:"-"`
	Email              string     `gorm:"column:email;unique" json:"email"`
	DeletedAt          *time.Time `gorm:"column:deleted_at" json:"deleted_at"`
	RoleID             uint       `gorm:"column:role_id" json:"role_id"`
	PositionID         *uint      `gorm:"column:position_id" json:"position_id"`
	EducationProgramId *uint      `gorm:"column:education_program_id" json:"education_program_id"`

	Role Role `json:"role"`
	//Positions        []Position       `gorm:"many2many:auth.position_user;foreignKey:UserID;joinForeignKey:UserID;joinReferences:PositionID;" json:"positions"`
	Position         *Position         `json:"position"`
	EducationProgram *EducationProgram `json:"edu_program_group"`
	TeamUsers        []TeamUsers       `json:"-"`
}

func (u User) TableName() string {
	return fmt.Sprintf("auth.user")
}

func (u *User) HasAllRoles(roles ...string) bool {
	for _, role := range roles {
		if u.Role.RoleName == role {
			return true
		}
	}

	return false
}
