package repository

import (
	"gorm.io/gorm"
	"work-project/internal/config"
)

type Repositories struct {
	User             User
	Profile          Profile
	ReferralCode     ReferralCode
	Referral         Referral
	PushNotification PushNotification
	Transaction      Transaction
	Balance          Balance
	UserDeviceToken  UserDeviceToken

	FirebaseMessaging FirebaseMessaging
}

func NewRepositories(db *gorm.DB, cfg *config.Config) (*Repositories, error) {
	return &Repositories{
		User:              NewUserDB(db),
		Profile:           NewProfileDB(db),
		ReferralCode:      NewReferralCodeDB(db),
		Referral:          NewReferralDB(db),
		PushNotification:  NewPushNotificationDB(db),
		Balance:           NewBalanceDB(db),
		Transaction:       NewTransactionDB(db),
		UserDeviceToken:   NewUserDeviceTokenDB(db),
		FirebaseMessaging: NewFirebaseClient(cfg.Integration.PathToFirebaseConfig),
	}, nil
}
