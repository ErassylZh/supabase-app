package service

import (
	"context"
	"work-project/internal/model"
	"work-project/internal/repository"
	"work-project/internal/schema"
)

type UserDeviceToken interface {
	GetByUserId(ctx context.Context, userId string) ([]model.UserDeviceToken, error)
	Create(ctx context.Context, data schema.UserDeviceTokenCreateRequest) (model.UserDeviceToken, error)
	DeleteById(ctx context.Context, deviceTokenID uint) error
}

type UserDeviceTokenService struct {
	userDeviceTokenRepo repository.UserDeviceToken
}

func NewUserDeviceTokenService(balanceRepo repository.UserDeviceToken) *UserDeviceTokenService {
	return &UserDeviceTokenService{userDeviceTokenRepo: balanceRepo}
}

func (s *UserDeviceTokenService) GetByUserId(ctx context.Context, userId string) ([]model.UserDeviceToken, error) {
	return s.userDeviceTokenRepo.GetByUserId(ctx, userId)
}

func (s *UserDeviceTokenService) Create(ctx context.Context, data schema.UserDeviceTokenCreateRequest) (model.UserDeviceToken, error) {
	return model.UserDeviceToken{}, s.userDeviceTokenRepo.Create(ctx, model.UserDeviceToken{
		UserID:      data.UserId,
		DeviceToken: data.Token,
	})
}

func (s *UserDeviceTokenService) DeleteById(ctx context.Context, deviceTokenID uint) error {
	return s.userDeviceTokenRepo.Delete(ctx, deviceTokenID)
}
