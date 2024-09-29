package service

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"time"
	"work-project/internal/model"
	"work-project/internal/repository"
	"work-project/internal/schema"
)

type Mark interface {
	CreateMark(ctx context.Context, mark schema.CreateMark) error
	FindByUserID(ctx context.Context, userID string) ([]schema.PostResponse, error)
	DeleteMark(ctx context.Context, markID uint) error
}

type MarkService struct {
	markRepo     repository.Mark
	postRepo     repository.Post
	userPostRepo repository.UserPost
}

func NewMarkService(markRepo repository.Mark, postRepo repository.Post, userPostRepo repository.UserPost) *MarkService {
	return &MarkService{
		markRepo:     markRepo,
		postRepo:     postRepo,
		userPostRepo: userPostRepo,
	}
}
func (s *MarkService) CreateMark(ctx context.Context, mark schema.CreateMark) error {
	if mark.UserID == "" || mark.PostId == 0 {
		return errors.New("invalid mark data: userID or postID is missing")
	}
	markModel := model.Mark{
		PostID:    mark.PostId,
		UserID:    mark.UserID,
		CreatedAt: time.Now(),
	}

	_, err := s.markRepo.FindByUserAndPost(ctx, markModel.UserID, markModel.PostID)
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		if err := s.markRepo.Create(ctx, markModel); err != nil {
			return errors.New("failed to create mark: " + err.Error())
		}
		return nil
	} else if err != nil {
		return err
	}

	return errors.New("this post already marked for current user")
}

func (s *MarkService) FindByUserID(ctx context.Context, userID string) ([]schema.PostResponse, error) {
	marks, err := s.markRepo.FindByUserID(ctx, userID)
	if err != nil {
		return nil, errors.New("failed to find marks: " + err.Error())
	}
	result := make([]schema.PostResponse, len(marks))
	for i, mark := range marks {
		_, err := s.userPostRepo.GetByUserAndPost(ctx, mark.UserID, mark.PostID)
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		result[i] = schema.PostResponse{
			Post:          mark.Post,
			IsAlreadyRead: !errors.Is(err, gorm.ErrRecordNotFound),
			MarkId:        &mark.MarkID,
			IsMarked:      true,
		}
	}

	return result, nil
}

func (s *MarkService) DeleteMark(ctx context.Context, markID uint) error {
	return s.markRepo.DeleteMark(ctx, markID)
}
