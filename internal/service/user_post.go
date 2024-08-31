package service

import (
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"work-project/internal/model"
	"work-project/internal/repository"
)

type UserPost interface {
	Create(ctx context.Context, post model.UserPost) (model.UserPost, error)
	AddQuizPoints(ctx context.Context, post model.UserPost) (model.UserPost, error)
	GetByUserAndPost(ctx context.Context, userId string, postId uint) (model.UserPost, error)
	GetAllByUser(ctx context.Context, id string) ([]model.UserPost, error)
}

type UserPostService struct {
	userPostRepo repository.UserPost
}

func NewUserPostService(userPostRepo repository.UserPost) *UserPostService {
	return &UserPostService{userPostRepo: userPostRepo}
}

func (s *UserPostService) Create(ctx context.Context, post model.UserPost) (model.UserPost, error) {
	return s.userPostRepo.Create(ctx, post)
}

func (s *UserPostService) AddQuizPoints(ctx context.Context, post model.UserPost) (model.UserPost, error) {
	up, err := s.userPostRepo.GetByUserAndPost(ctx, post.UserId, post.PostId)
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return up, fmt.Errorf("user not readed this post")
	} else if err != nil {
		return model.UserPost{}, err
	}
	up.QuizPoints = post.QuizPoints
	up.QuizSapphires = post.QuizSapphires
	return s.AddQuizPoints(ctx, up)
}

func (s *UserPostService) GetByUserAndPost(ctx context.Context, userId string, postId uint) (model.UserPost, error) {
	return s.userPostRepo.GetByUserAndPost(ctx, userId, postId)
}

func (s *UserPostService) GetAllByUser(ctx context.Context, id string) ([]model.UserPost, error) {
	return s.userPostRepo.GetByUserID(ctx, id)
}
