package service

import (
	"context"
	"fmt"
	"strings"
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
	profileRepo  repository.Profile
	userRepo     repository.User
	badWordsRepo repository.BadWord
}

func NewUserService(userRepo repository.User, profile repository.Profile, badWords repository.BadWord) *UserService {
	return &UserService{
		userRepo:     userRepo,
		profileRepo:  profile,
		badWordsRepo: badWords,
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

	if len(data.Nickname) > 0 {
		badWords, err := s.badWordsRepo.GetAll(ctx)
		if err != nil {
			return err
		}
		nickname := strings.ToLower(data.Nickname)
		for _, badWord := range badWords {
			if strings.Contains(badWord, nickname) {
				return fmt.Errorf("invalid word")
			}
		}

		profile.UserName = data.Nickname
	}

	return s.profileRepo.Update(ctx, profile)
}
