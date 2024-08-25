package service

import (
	"context"
	"errors"
	"work-project/internal/model"
	"work-project/internal/repository"
)

type Mark interface {
	CreateMark(ctx context.Context, mark *model.Mark) error
	FindByUserID(ctx context.Context, userID string) ([]model.Mark, error)
	DeleteMark(ctx context.Context, markID uint) error
	FindPostsByUserID(ctx context.Context, userID string) ([]model.Post, error)
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
func (s *MarkService) CreateMark(ctx context.Context, mark *model.Mark) error {
	if mark == nil {
		return errors.New("mark is nil")
	}
	if mark.UserID == "" || mark.PostID == 0 {
		return errors.New("invalid mark data: userID or postID is missing")
	}
	if err := s.markRepo.CreateMark(ctx, mark); err != nil {
		return errors.New("failed to create mark: " + err.Error())
	}

	return nil
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

func (s *MarkService) FindPostsByUserID(ctx context.Context, userID string) ([]model.Post, error) {
	return s.markRepo.FindPostsByUserID(ctx, userID)
}
