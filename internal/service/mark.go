package service

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"work-project/internal/model"
	"work-project/internal/repository"
)

type Mark interface {
	CreateMark(ctx context.Context, mark model.Mark) error
	FindByUserID(ctx context.Context, userID string) ([]model.Mark, error)
	DeleteMark(ctx context.Context, markID uint) error
}

type MarkService struct {
	markRepo repository.Mark
	postRepo repository.Post
}

func NewMarkService(markRepo repository.Mark, postRepo repository.Post) *MarkService {
	return &MarkService{
		markRepo: markRepo,
		postRepo: postRepo,
	}
}
func (s *MarkService) CreateMark(ctx context.Context, mark model.Mark) error {
	if mark.UserID == "" || mark.PostID == 0 {
		return errors.New("invalid mark data: userID or postID is missing")
	}
	_, err := s.markRepo.FindByUserAndPost(ctx, mark.UserID, mark.MarkID)
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		if err := s.markRepo.Create(ctx, mark); err != nil {
			return errors.New("failed to create mark: " + err.Error())
		}
		return nil
	} else if err != nil {
		return err
	}

	return errors.New("this post already marked for current user")
}

func (s *MarkService) FindByUserID(ctx context.Context, userID string) ([]model.Mark, error) {
	marks, err := s.markRepo.FindByUserID(ctx, userID)
	if err != nil {
		return nil, errors.New("failed to find marks: " + err.Error())
	}
	return marks, nil
}

func (s *MarkService) DeleteMark(ctx context.Context, markID uint) error {
	return s.markRepo.DeleteMark(ctx, markID)
}
