package repository

import (
	"gorm.io/gorm"
	"work-project/internal/config"
)

type Repositories struct {
	User         User
	Profile      Profile
	ReferralCode ReferralCode
	Referral     Referral
}

func NewRepositories(db *gorm.DB, cfg *config.Config) (*Repositories, error) {
	return &Repositories{
		User:         NewUserDB(db),
		Profile:      NewProfileDB(db),
		ReferralCode: NewReferralCodeDB(db),
		Referral:     NewReferralDB(db),
	}, nil
}
