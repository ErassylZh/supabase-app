package service

import (
	"context"
	"work-project/internal/model"
	"work-project/internal/repository"
)

type Stories interface {
	GetByUserId(ctx context.Context, userId string) ([]model.Stories, error)
	ReadStory(ctx context.Context, userId string, storyId uint) error
}

type StoriesService struct {
	storiesRepo       repository.Stories
	storyPageRepo     repository.StoryPage
	storyPageUserRepo repository.StoryPageUser
}

func NewStoriesService(storiesRepo repository.Stories, storyPageRepo repository.StoryPage, storyPageUserRepo repository.StoryPageUser) *StoriesService {
	return &StoriesService{storiesRepo: storiesRepo, storyPageRepo: storyPageRepo, storyPageUserRepo: storyPageUserRepo}
}

func (s *StoriesService) GetByUserId(ctx context.Context, userId string) ([]model.Stories, error) {
	if len(userId) == 0 {
		return s.storiesRepo.GetAllActive(ctx)
	}
	return s.storiesRepo.GetAllActiveByUser(ctx, userId)
}

func (s *StoriesService) ReadStory(ctx context.Context, userId string, storyId uint) error {
	return s.storyPageUserRepo.Create(ctx, model.StoryPageUser{
		UserId:      userId,
		StoryPageId: storyId,
	})
}
