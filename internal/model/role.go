package model

import (
	"fmt"
)

type Role struct {
	RoleID      uint   `gorm:"primaryKey;column:role_id" json:"id"`
	RoleName    string `gorm:"column:role_name;" json:"role_name"`
	Description string `json:"description"`

	//Permissions []Permission `gorm:"many2many:auth.role_permission;foreignKey:RoleID;joinForeignKey:RoleID;joinReferences:RoleID;" json:"permissions"`
}

func (u Role) TableName() string {
	return fmt.Sprintf("auth.role")
}

type RoleCodeName string

const (
	ROLE_NAME_STUDENT    RoleCodeName = "STUDENT"
	ROLE_NAME_SUPERVISOR RoleCodeName = "SUPERVISOR"
	ROLE_NAME_TEACHER    RoleCodeName = "TEACHER"
	ROLE_NAME_TECH_SEC   RoleCodeName = "TECH_SEC"
)
