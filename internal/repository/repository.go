package repository

import (
	"gorm.io/gorm"
	"work-project/internal/config"
)

type Repositories struct {
	User    User
	Profile Profile
}

func NewRepositories(db *gorm.DB, cfg *config.Config) (*Repositories, error) {
	return &Repositories{
		User:    NewUserDB(db),
		Profile: NewProfileDB(db),
	}, nil
}
