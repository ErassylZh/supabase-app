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
	Hashtag          Hashtag
	PostHashtag      PostHashtag
	Stories          Stories
	StoryPage        StoryPage
	StoryPageUser    StoryPageUser
	Mark             Mark
	UserPost         UserPost
	Collection       Collection
	PostCollection   PostCollection

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
		PostHashtag:       NewPostHashtagDb(db),
		Hashtag:           NewHashtagDB(db),
		Stories:           NewStoriesDB(db),
		StoryPage:         NewStoryPageDB(db),
		StoryPageUser:     NewStoryPageUserDB(db),
		Mark:              NewMarkDb(db),
		UserPost:          NewUserPostRepository(db),
		Collection:        NewCollectionDB(db),
		PostCollection:    NewPostCollectionDb(db),
		FirebaseMessaging: NewFirebaseClient(cfg.Integration.PathToFirebaseConfig),
		Airtable:          airtable,
		StorageClient:     NewStorageClient(cfg.Database.SupabaseUrl, cfg.Database.SupabaseApiKey),
	}, nil
}
