package usecase

import (
	"work-project/internal/repository"
	"work-project/internal/service"
)

type Usecases struct {
	Referral Referral
	Post     Post
}

type Deps struct {
	Services     *service.Services
	Repositories *repository.Repositories
}

func NewUsecases(deps Deps) *Usecases {
	return &Usecases{
		Referral: NewReferralUsecase(deps.Repositories.ReferralCode, deps.Repositories.Referral, deps.Services.PushNotification, deps.Services.Balance, deps.Services.UserDeviceToken),
		Post:     NewPostUsecase(deps.Services),
	}
}
