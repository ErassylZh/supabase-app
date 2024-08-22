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
	Product          Product
	Post             Post
	Image            Image
	Mark             Mark

	FirebaseMessaging FirebaseMessaging
	Airtable          AirTable
	StorageClient     StorageClient
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
		UserDeviceToken:   NewUserDeviceTokenDB(db),
		Product:           NewProductDb(db),
		Image:             NewImageDb(db),
		Post:              NewPostDb(db),
		FirebaseMessaging: NewFirebaseClient(cfg.Integration.PathToFirebaseConfig),
		Airtable:          airtable,
		StorageClient:     NewStorageClient(cfg.Database.SupabaseUrl, cfg.Database.SupabaseApiKey),
	}, nil
}
