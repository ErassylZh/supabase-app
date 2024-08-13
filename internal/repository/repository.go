package repository

import (
	"gorm.io/gorm"
	"work-project/internal/config"
	"work-project/internal/repository/integration"
)

type Repositories struct {
	User             User
	Profile          Profile
	ReferralCode     ReferralCode
	Referral         Referral
	PushNotification PushNotification
	Transaction      Transaction
	Balance          Balance
	Product          Product

	FirebaseMessaging integration.FirebaseMessaging
	Airtable          AirTable
	StorageClient     integration.StorageClient
}

func NewRepositories(db *gorm.DB, cfg *config.Config) (*Repositories, error) {
	airtable, err := NewAirTableClient(cfg.Integration.AirtableBaseurl, cfg.Integration.AirtableApiKey)
	if err != nil {
		return nil, err
	}
	return &Repositories{
		User:              NewUserDB(db),
		Profile:           NewProfileDB(db),
		ReferralCode:      NewReferralCodeDB(db),
		Referral:          NewReferralDB(db),
		PushNotification:  NewPushNotificationDB(db),
		Balance:           NewBalanceDB(db),
		Transaction:       NewTransactionDB(db),
		Product:           NewProductDb(db),
		FirebaseMessaging: integration.NewFirebaseClient(cfg.Integration.PathToFirebaseConfig),
		Airtable:          airtable,
		StorageClient:     integration.NewStorageClient(cfg.Database.SupabaseUrl, cfg.Database.SupabaseApiKey),
	}, nil
}
