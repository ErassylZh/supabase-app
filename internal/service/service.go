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
}

type Deps struct {
	Repos *repository.Repositories
	Cgf   *config.Config
}

func NewServices(deps Deps) *Services {
	return &Services{
		User:             NewUserService(deps.Repos.User, deps.Repos.Profile),
		Auth:             NewAuthService(deps.Cgf.Security.Secret),
		PushNotification: NewPushNotificationService(deps.Repos.PushNotification, deps.Repos.FirebaseMessaging),
		Balance:          NewBalanceService(deps.Repos.Balance, deps.Repos.Transaction),
		UserDeviceToken:  NewUserDeviceTokenService(deps.Repos.UserDeviceToken),
		Product:          NewProductService(deps.Repos.Product),
		Post:             NewPostService(deps.Repos.Post),
		Stories:          NewStoriesService(deps.Repos.Stories, deps.Repos.StoryPage, deps.Repos.StoryPageUser),
		Mark:             NewMarkService(deps.Repos.Mark, deps.Repos.Post),
	}
}
