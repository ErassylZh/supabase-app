package service

import (
	"context"
	"work-project/internal/model"
	"work-project/internal/repository"
)

type User interface {
	DeleteByID(ctx context.Context, userID string) error
	GetById(ctx context.Context, userID string) (model.User, error)
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
