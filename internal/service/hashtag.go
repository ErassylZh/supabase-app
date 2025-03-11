package service

import (
	"context"
	"errors"
	"log"
	"strings"
	"time"
	"work-project/internal/admin"
	"work-project/internal/model"
	"work-project/internal/repository"
)

type Hashtag interface {
	GetAll(ctx context.Context) ([]model.Hashtag, error)
	GetByID(ctx context.Context, id uint) (model.Hashtag, error)
	Delete(ctx context.Context, id uint) error
	Create(ctx context.Context, data admin.CreateHashtag) (model.Hashtag, error)
	Update(ctx context.Context, data admin.UpdateHashtag) (model.Hashtag, error)
	AddToPost(ctx context.Context, data admin.AddHashtag) (model.Hashtag, error)
	DeleteHashtagPost(ctx context.Context, data admin.DeleteHashtagPost) (model.Hashtag, error)
}

type HashtagService struct {
	hashtagRepo repository.Hashtag
	postHashtag repository.PostHashtag
	storage     repository.StorageClient
}

func NewHashtagService(hashtagRepo repository.Hashtag, postHashtag repository.PostHashtag, storage repository.StorageClient) *HashtagService {
	return &HashtagService{hashtagRepo: hashtagRepo, postHashtag: postHashtag, storage: storage}
}

func (s *HashtagService) GetAll(ctx context.Context) ([]model.Hashtag, error) {
	return s.hashtagRepo.GetVisible(ctx)
}

func (s *HashtagService) GetByID(ctx context.Context, id uint) (model.Hashtag, error) {
	return s.hashtagRepo.GetByID(ctx, id)
}

func (s *HashtagService) Delete(ctx context.Context, id uint) error {
	return s.hashtagRepo.Delete(ctx, id)
}

func (s *HashtagService) saveBase64Image(ctx context.Context, base64Str, filename string) (*string, error) {
	if base64Str == "" {
		return nil, nil
	}

	if strings.Contains(base64Str, "base64,") {
		parts := strings.Split(base64Str, "base64,")
		if len(parts) < 2 {
			return nil, errors.New("неправильный формат base64")
		}
		base64Str = parts[1]
	}

	filePath, err := s.storage.CreateImageFromBase64(ctx, string(model.BUCKET_NAME_HASHTAG), time.Now().String()+filename, base64Str)
	if err != nil {
		return nil, err
	}

	return &filePath, nil
}

func (s *HashtagService) Create(ctx context.Context, data admin.CreateHashtag) (model.Hashtag, error) {
	// Обрабатываем изображение, если передано
	var imagePath *string
	if data.Image != nil && data.Image.File != "" {
		newPath, err := s.saveBase64Image(ctx, data.Image.File, data.Image.FileName)
		if err != nil {
			log.Println("Ошибка сохранения image_path:", err)
			return model.Hashtag{}, err
		}
		imagePath = newPath
	}

	// Создаем хэштег
	hashtag := model.Hashtag{
		Name:      data.Name,
		NameRu:    data.NameRu,
		NameKz:    data.NameKz,
		IsVisible: data.IsVisible,
		ImagePath: imagePath,
	}

	createdHashtag, err := s.hashtagRepo.Create(ctx, hashtag)
	if err != nil {
		return model.Hashtag{}, err
	}

	return createdHashtag, nil
}

// Обновление хэштега
func (s *HashtagService) Update(ctx context.Context, data admin.UpdateHashtag) (model.Hashtag, error) {
	hashtag, err := s.hashtagRepo.GetByID(ctx, data.HashtagID)
	if err != nil {
		return model.Hashtag{}, err
	}

	// Обновляем поля
	hashtag.Name = data.Name
	hashtag.NameRu = data.NameRu
	hashtag.NameKz = data.NameKz
	hashtag.IsVisible = data.IsVisible

	// Обрабатываем изображение, если передано
	if data.Image != nil && data.Image.File != "" {
		newPath, err := s.saveBase64Image(ctx, data.Image.File, data.Image.FileName)
		if err != nil {
			log.Println("Ошибка сохранения image_path:", err)
			return model.Hashtag{}, err
		}
		hashtag.ImagePath = newPath
	}

	// Сохраняем обновленный хэштег
	updatedHashtag, err := s.hashtagRepo.Update(ctx, hashtag)
	if err != nil {
		return model.Hashtag{}, err
	}

	return updatedHashtag, nil
}

func (s *HashtagService) AddToPost(ctx context.Context, data admin.AddHashtag) (model.Hashtag, error) {
	_, err := s.postHashtag.Create(ctx, model.PostHashtag{
		PostId:    data.PostID,
		HashtagId: data.HashtagID,
	})
	return model.Hashtag{}, err
}

func (s *HashtagService) DeleteHashtagPost(ctx context.Context, data admin.DeleteHashtagPost) (model.Hashtag, error) {
	err := s.postHashtag.DeleteByPostAndHashtagId(ctx, data.PostID, data.HashtagID)
	return model.Hashtag{}, err
}
