package service

import (
	"context"
	"log"
	"time"
	"work-project/internal/admin"
	"work-project/internal/model"
	"work-project/internal/repository"
)

type Stories interface {
	GetByUserId(ctx context.Context, userId string) ([]model.Stories, error)
	ReadStory(ctx context.Context, userId string, storyId uint) error
	Create(ctx context.Context, data admin.CreateStories) (model.Stories, error)
	Update(ctx context.Context, data admin.UpdateStories) (model.Stories, error)
	GetAll(ctx context.Context) ([]model.Stories, error)
	GetByID(ctx context.Context, id uint) (model.Stories, error)
	DeleteByID(ctx context.Context, id uint) error
	CreatePage(ctx context.Context, data admin.CreateStoryPage) (model.StoryPage, error)
	UpdateStoryPage(ctx context.Context, data admin.UpdateStoryPage) (model.StoryPage, error)
	DeletePageByID(ctx context.Context, id uint) error
}

type StoriesService struct {
	storiesRepo       repository.Stories
	storyPageRepo     repository.StoryPage
	storyPageUserRepo repository.StoryPageUser
	storage           repository.Storage
}

func NewStoriesService(storiesRepo repository.Stories, storyPageRepo repository.StoryPage, storyPageUserRepo repository.StoryPageUser, storage repository.Storage) *StoriesService {
	return &StoriesService{storiesRepo: storiesRepo, storyPageRepo: storyPageRepo, storyPageUserRepo: storyPageUserRepo, storage: storage}
}

func (s *StoriesService) GetByUserId(ctx context.Context, userId string) ([]model.Stories, error) {
	if len(userId) == 0 {
		return s.storiesRepo.GetAllActive(ctx)
	}
	stories, err := s.storiesRepo.GetAllActive(ctx)
	if err != nil {
		return nil, err
	}
	storyIds := make([]uint, len(stories))
	for i, story := range stories {
		storyIds[i] = story.StoriesId
	}
	userPages, err := s.storyPageUserRepo.GetAllByStoryIdUserID(ctx, storyIds, userId)
	if err != nil {
		return nil, err
	}
	for i := range stories {
		for j := range stories[i].StoryPages {
			for _, up := range userPages {
				if up.StoryPageId == stories[i].StoryPages[j].StoryPageId {
					stories[i].StoryPages[j].IsReaded = true
					break
				}
			}
		}
	}
	return stories, nil
}

func (s *StoriesService) ReadStory(ctx context.Context, userId string, storyId uint) error {
	return s.storyPageUserRepo.Create(ctx, model.StoryPageUser{
		UserId:      userId,
		StoryPageId: storyId,
	})
}

func (s *StoriesService) Create(ctx context.Context, data admin.CreateStories) (model.Stories, error) {
	if data.Icon != nil {
		file, err := s.storage.CreateImageFromBase64(ctx, string(model.BUCKET_NAME_STORIES), time.Now().String()+data.Icon.FileName, data.Icon.File)
		if err != nil {
			log.Println(ctx, "some err while create image", "err", err, "story name", data.Title)
			return model.Stories{}, err
		}

		return s.storiesRepo.Create(ctx, model.Stories{
			Title:     data.Title,
			StartTime: data.StartTime,
			EndTime:   data.EndTime,
			IconPath:  file,
		})
	}

	return s.storiesRepo.Create(ctx, model.Stories{
		Title:     data.Title,
		StartTime: data.StartTime,
		EndTime:   data.EndTime,
	})
}

func (s *StoriesService) Update(ctx context.Context, data admin.UpdateStories) (model.Stories, error) {
	story, err := s.storiesRepo.GetByID(ctx, data.StoriesID)
	if err != nil {
		return model.Stories{}, err
	}

	if data.EndTime != nil {
		story.EndTime = *data.EndTime
	}
	if data.StartTime != nil {
		story.StartTime = *data.StartTime
	}
	if data.Title != nil {
		story.Title = *data.Title
	}

	if data.Icon != nil {
		file, err := s.storage.CreateImageFromBase64(ctx, string(model.BUCKET_NAME_STORIES), time.Now().String()+data.Icon.FileName, data.Icon.File)
		if err != nil {
			log.Println(ctx, "some err while create image", "err", err, "story name", data.Title)
			return model.Stories{}, err
		}

		story.IconPath = file
	}

	return s.storiesRepo.Update(ctx, story)
}

func (s *StoriesService) GetAll(ctx context.Context) ([]model.Stories, error) {
	return s.storiesRepo.GetAll(ctx)
}

func (s *StoriesService) GetByID(ctx context.Context, id uint) (model.Stories, error) {
	return s.storiesRepo.GetByID(ctx, id)
}

func (s *StoriesService) DeleteByID(ctx context.Context, id uint) error {
	return s.storiesRepo.DeleteByID(ctx, id)
}

func (s *StoriesService) CreatePage(ctx context.Context, data admin.CreateStoryPage) (model.StoryPage, error) {
	if data.Icon != nil {
		file, err := s.storage.CreateImageFromBase64(ctx, string(model.BUCKET_NAME_STORIES), time.Now().String()+data.Icon.FileName, data.Icon.File)
		if err != nil {
			log.Println(ctx, "some err while create image", "err", err, "story name", data.Text)
			return model.StoryPage{}, err
		}

		return s.storyPageRepo.Create(ctx, model.StoryPage{
			StoriesId: data.StoriesId,
			Text:      data.Text,
			PageOrder: data.PageOrder,
			ImagePath: file,
		})
	}

	return s.storyPageRepo.Create(ctx, model.StoryPage{
		StoriesId: data.StoriesId,
		Text:      data.Text,
		PageOrder: data.PageOrder,
	})
}

func (s *StoriesService) UpdateStoryPage(ctx context.Context, data admin.UpdateStoryPage) (model.StoryPage, error) {
	storyPage, err := s.storyPageRepo.GetByID(ctx, data.StoryPageID)
	if err != nil {
		return model.StoryPage{}, err
	}

	if data.Text != nil {
		storyPage.Text = *data.Text
	}
	if data.PageOrder != nil {
		storyPage.PageOrder = *data.PageOrder
	}
	if data.Icon != nil {
		file, err := s.storage.CreateImageFromBase64(ctx, string(model.BUCKET_NAME_STORIES), time.Now().String()+data.Icon.FileName, data.Icon.File)
		if err != nil {
			log.Println(ctx, "some err while create image", "err", err, "story name", data.Text)
			return model.StoryPage{}, err
		}

		storyPage.ImagePath = file
	}

	return s.storyPageRepo.Update(ctx, storyPage)
}

func (s *StoriesService) DeletePageByID(ctx context.Context, id uint) error {
	return s.storyPageRepo.Delete(ctx, id)
}
