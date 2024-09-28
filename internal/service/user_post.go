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
	postRepo     repository.Post
}

func NewUserPostService(userPostRepo repository.UserPost, postRepo repository.Post) *UserPostService {
	return &UserPostService{userPostRepo: userPostRepo, postRepo: postRepo}
}

func (s *UserPostService) Create(ctx context.Context, post model.UserPost) (model.UserPost, error) {

	return s.userPostRepo.Create(ctx, post)
}

func (s *UserPostService) AddQuizPoints(ctx context.Context, createUserPost model.UserPost) (model.UserPost, error) {
	up, err := s.userPostRepo.GetByUserAndPost(ctx, createUserPost.UserId, createUserPost.PostId)
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return up, fmt.Errorf("user not readed this post")
	} else if err != nil {
		return model.UserPost{}, err
	}

	posts, err := s.postRepo.GetAllGroupedByPostId(ctx, createUserPost.PostId)
	if err != nil {
		return model.UserPost{}, err
	}
	for _, post := range posts {
		userPost, err := s.userPostRepo.GetByUserAndPost(ctx, createUserPost.UserId, post.PostID)
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return model.UserPost{}, err
		}
		if userPost.QuizPoints != nil || userPost.QuizSapphires != nil {
			return model.UserPost{}, errors.New("you already get coins for this post")
		}
	}

	up.QuizPoints = createUserPost.QuizPoints
	up.QuizSapphires = createUserPost.QuizSapphires
	return s.userPostRepo.Update(ctx, up)
}

func (s *UserPostService) GetByUserAndPost(ctx context.Context, userId string, postId uint) (model.UserPost, error) {
	return s.userPostRepo.GetByUserAndPost(ctx, userId, postId)
}

func (s *UserPostService) GetAllByUser(ctx context.Context, id string) ([]model.UserPost, error) {
	return s.userPostRepo.GetByUserID(ctx, id)
}
