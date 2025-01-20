package service

import (
	"work-project/internal/config"
	"work-project/internal/repository"
)

//go:generate go run github.com/golang/mock/mockgen -source=service.go -destination=service_mock.go -package=service

type Services struct {
	Auth             Auth
	User             User
	PushNotification PushNotification
	Balance          Balance
	UserDeviceToken  UserDeviceToken
	Product          Product
	Post             Post
	Stories          Stories
	Mark             Mark
	Hashtag          Hashtag
	Collection       Collection
	UserPost         UserPost
	Order            Order
	EmailSender      EmailSender
	ProductTag       ProductTag
	Contest          Contest
}

type Deps struct {
	Repos *repository.Repositories
	Cgf   *config.Config
}

func NewServices(deps Deps) *Services {
	postService := NewPostService(deps.Repos.Post)
	markService := NewMarkService(deps.Repos.Mark, deps.Repos.Post, deps.Repos.UserPost)
	return &Services{
		User:             NewUserService(deps.Repos.User, deps.Repos.Profile),
		Auth:             NewAuthService(deps.Repos.User, deps.Cgf.Security.Secret),
		PushNotification: NewPushNotificationService(deps.Repos.PushNotification, deps.Repos.FirebaseMessaging),
		Balance:          NewBalanceService(deps.Repos.Balance, deps.Repos.Transaction),
		UserDeviceToken:  NewUserDeviceTokenService(deps.Repos.UserDeviceToken),
		Product:          NewProductService(deps.Repos.Product),
		Post:             NewPostService(deps.Repos.Post),
		Stories:          NewStoriesService(deps.Repos.Stories, deps.Repos.StoryPage, deps.Repos.StoryPageUser),
		Mark:             markService,
		Hashtag:          NewHashtagService(deps.Repos.Hashtag),
		UserPost:         NewUserPostService(deps.Repos.UserPost, deps.Repos.Post),
		Collection:       NewCollectionService(deps.Repos.Collection, deps.Repos.UserPost, postService, deps.Repos.Mark),
		EmailSender:      NewEmailSenderService(deps.Cgf.Email.Username, deps.Cgf.Email.Password, deps.Cgf.Email.Addr),
		Order:            NewOrderService(deps.Repos.Order, deps.Repos.OrderProduct),
		ProductTag:       NewProductTagService(deps.Repos.ProductTag),
		Contest:          NewContestService(deps.Repos.Contest, deps.Repos.ContestParticipant, deps.Repos.ContestBook),
	}
}
