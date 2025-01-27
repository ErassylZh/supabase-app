package service

import (
	"context"
	"work-project/internal/model"
	"work-project/internal/repository"
	"work-project/internal/schema"
)

type User interface {
	DeleteByID(ctx context.Context, userID string) error
	GetById(ctx context.Context, userID string) (model.User, error)
	Update(ctx context.Context, data schema.UserUpdate) error
}

type UserService struct {
	profileRepo repository.Profile
	userRepo    repository.User
}

func NewUserService(userRepo repository.User, profile repository.Profile) *UserService {
	return &UserService{
		userRepo:    userRepo,
		profileRepo: profile,
	}
}

func (s *UserService) DeleteByID(ctx context.Context, userID string) error {
	err := s.profileRepo.DeleteByID(ctx, userID)
	if err != nil {
		return err
	}
	return s.userRepo.DeleteByID(ctx, userID)
}

func (s *UserService) GetById(ctx context.Context, userID string) (model.User, error) {
	return s.userRepo.GetByID(ctx, userID)
}

func (s *UserService) Update(ctx context.Context, data schema.UserUpdate) error {
	profile, err := s.profileRepo.GetByID(ctx, data.UserID)
	if err != nil {
		return err
	}
	profile.UserName = data.Nickname

	return s.profileRepo.Update(ctx, profile)
}
