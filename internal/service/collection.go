package service

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"log"
	"strings"
	"time"
	"work-project/internal/admin"
	"work-project/internal/model"
	"work-project/internal/repository"
	"work-project/internal/schema"
)

type Collection interface {
	GetAllCollection(ctx context.Context, language string, userId *string, withoutPosts bool) ([]schema.CollectionListResponse, error)
	GetAllRecommendation(ctx context.Context, language string) ([]schema.CollectionListResponse, error)
	Create(ctx context.Context, data admin.CreateCollection) (model.Collection, error)
	GetAll(ctx context.Context) ([]model.Collection, error)
	GetByID(ctx context.Context, id uint) (model.Collection, error)
	Update(ctx context.Context, data admin.UpdateCollection) (model.Collection, error)
	Delete(ctx context.Context, id uint) error
	AddToPost(ctx context.Context, data admin.AddCollection) (model.Collection, error)
	DeleteCollectionPost(ctx context.Context, data admin.DeleteCollectionPost) (model.Collection, error)
}

type CollectionService struct {
	collectionRepo     repository.Collection
	postCollectionRepo repository.PostCollection
	userPostRepo       repository.UserPost
	markRepo           repository.Mark
	storage            repository.Storage
	postService        Post
}

func NewCollectionService(collectionRepo repository.Collection, userPostRepo repository.UserPost, markRepo repository.Mark, storage repository.Storage, postService Post) *CollectionService {
	return &CollectionService{collectionRepo: collectionRepo, userPostRepo: userPostRepo, markRepo: markRepo, storage: storage, postService: postService}
}

func (s *CollectionService) GetAllCollection(ctx context.Context, language string, userId *string, withoutPosts bool) ([]schema.CollectionListResponse, error) {
	collections, err := s.collectionRepo.GetAllCollection(ctx, language, withoutPosts)
	if err != nil {
		return nil, err
	}

	postIdMark := make(map[uint]model.Mark)
	userPostMap := make(map[uint]model.UserPost)
	if userId != nil && !withoutPosts {
		userMarks, err := s.markRepo.FindByUserID(ctx, *userId)
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		for _, um := range userMarks {
			postIdMark[um.PostID] = um
		}

		userPosts, err := s.userPostRepo.GetByUserID(ctx, *userId)
		if err != nil {
			return nil, err
		}
		for _, up := range userPosts {
			userPostMap[up.PostId] = up
		}
	}

	result := make([]schema.CollectionListResponse, len(collections))

	collectionIDs := make([]uint, len(collections))
	for i, collection := range collections {
		collectionIDs[i] = collection.CollectionID
	}

	for i, collection := range collections {
		// Обрабатываем посты в коллекции
		posts := make([]schema.PostResponse, len(collection.Posts))

		for j := range collection.Posts {
			posts[j].Post = collection.Posts[j]
			if um, exists := postIdMark[posts[j].PostID]; exists {
				posts[j].IsMarked = true
				posts[j].MarkId = &um.MarkID
			}

			if up, exists := userPostMap[posts[j].PostID]; exists {
				posts[j].QuizPassed = up.QuizPoints != nil || up.QuizSapphires != nil
				posts[j].IsAlreadyRead = up.ReadEnd
			}
		}

		collectionResponse := schema.CollectionListResponse{
			CollectionID: collection.CollectionID,
			Name:         collection.Name,
			NameRu:       collection.NameRu,
			NameKz:       collection.NameKz,
			ImagePath:    collection.ImagePath,
			ImagePathRu:  collection.ImagePathRu,
			ImagePathKz:  collection.ImagePathKz,
			Posts:        posts,
		}
		result[i] = collectionResponse
	}

	return result, nil
}

func (s *CollectionService) GetAllRecommendation(ctx context.Context, language string) ([]schema.CollectionListResponse, error) {
	collections, err := s.collectionRepo.GetAllRecommendation(ctx, language)
	if err != nil {
		return nil, err
	}

	result := make([]schema.CollectionListResponse, len(collections))
	for i, collection := range collections {
		posts, _, err := s.postService.GetListing(ctx, schema.GetListingFilter{
			CollectionIds: []uint{collection.CollectionID},
			Language:      &language,
		})
		if err != nil {
			return nil, err
		}
		collectionResponse := schema.CollectionListResponse{
			CollectionID: collection.CollectionID,
			Name:         collection.Name,
			NameRu:       collection.NameRu,
			NameKz:       collection.NameKz,
			ImagePath:    collection.ImagePath,
			ImagePathRu:  collection.ImagePathRu,
			ImagePathKz:  collection.ImagePathKz,
			Posts:        posts,
		}
		result[i] = collectionResponse
	}

	return result, nil
}

func (s *CollectionService) Create(ctx context.Context, data admin.CreateCollection) (model.Collection, error) {
	collection := model.Collection{}

	// Обновляем поля
	collection.Name = data.Name
	collection.NameRu = data.NameRu
	collection.NameKz = data.NameKz
	collection.IsRecommendation = data.IsRecommendation

	// Обрабатываем изображения, если они есть
	if data.Image != nil && data.Image.File != "" {
		newPath, err := s.saveBase64Image(ctx, data.Image.File, data.Image.FileName)
		if err != nil {
			log.Println("Ошибка сохранения image_path:", err)
			return model.Collection{}, err
		}
		collection.ImagePath = newPath
	}

	// Обрабатываем изображения, если они есть
	if data.ImageKz != nil && data.ImageKz.File != "" {
		newPath, err := s.saveBase64Image(ctx, data.ImageKz.File, data.ImageKz.FileName)
		if err != nil {
			log.Println("Ошибка сохранения image_path:", err)
			return model.Collection{}, err
		}
		collection.ImagePathKz = newPath
	}

	// Обрабатываем изображения, если они есть
	if data.ImageRu != nil && data.ImageRu.File != "" {
		newPath, err := s.saveBase64Image(ctx, data.ImageRu.File, data.ImageRu.FileName)
		if err != nil {
			log.Println("Ошибка сохранения image_path:", err)
			return model.Collection{}, err
		}
		collection.ImagePathRu = newPath
	}

	// Сохраняем обновленную коллекцию
	updatedCollection, err := s.collectionRepo.Create(ctx, collection)
	if err != nil {
		return model.Collection{}, err
	}

	return updatedCollection, nil
}

// Получение всех коллекций
func (s *CollectionService) GetAll(ctx context.Context) ([]model.Collection, error) {
	return s.collectionRepo.GetAll(ctx)
}

// Получение коллекции по ID
func (s *CollectionService) GetByID(ctx context.Context, id uint) (model.Collection, error) {
	return s.collectionRepo.GetByID(ctx, id)
}

func (s *CollectionService) saveBase64Image(ctx context.Context, base64Str, filename string) (*string, error) {
	if base64Str == "" {
		return nil, nil
	}

	// Проверяем, есть ли `data:image/png;base64,`
	if strings.Contains(base64Str, "base64,") {
		parts := strings.Split(base64Str, "base64,")
		if len(parts) < 2 {
			return nil, errors.New("неправильный формат base64")
		}
		base64Str = parts[1]
	}

	// Сохраняем файл
	filePath, err := s.storage.CreateImageFromBase64(ctx, string(model.BUCKET_NAME_COLLECTION), time.Now().String()+filename, base64Str)
	if err != nil {
		return nil, err
	}

	return &filePath, nil
}

// Обновление коллекции
func (s *CollectionService) Update(ctx context.Context, data admin.UpdateCollection) (model.Collection, error) {
	collection, err := s.collectionRepo.GetByID(ctx, data.CollectionID)
	if err != nil {
		return model.Collection{}, err
	}

	// Обновляем поля
	collection.Name = data.Name
	collection.NameRu = data.NameRu
	collection.NameKz = data.NameKz
	collection.IsRecommendation = data.IsRecommendation

	// Обрабатываем изображения, если они есть
	if data.Image != nil && data.Image.File != "" {
		newPath, err := s.saveBase64Image(ctx, data.Image.File, data.Image.FileName)
		if err != nil {
			log.Println("Ошибка сохранения image_path:", err)
			return model.Collection{}, err
		}
		collection.ImagePath = newPath
	}

	// Обрабатываем изображения, если они есть
	if data.ImageKz != nil && data.ImageKz.File != "" {
		newPath, err := s.saveBase64Image(ctx, data.ImageKz.File, data.ImageKz.FileName)
		if err != nil {
			log.Println("Ошибка сохранения image_path:", err)
			return model.Collection{}, err
		}
		collection.ImagePathKz = newPath
	}

	// Обрабатываем изображения, если они есть
	if data.ImageRu != nil && data.ImageRu.File != "" {
		newPath, err := s.saveBase64Image(ctx, data.ImageRu.File, data.ImageRu.FileName)
		if err != nil {
			log.Println("Ошибка сохранения image_path:", err)
			return model.Collection{}, err
		}
		collection.ImagePathRu = newPath
	}

	// Сохраняем обновленную коллекцию
	updatedCollection, err := s.collectionRepo.Update(ctx, collection)
	if err != nil {
		return model.Collection{}, err
	}

	return updatedCollection, nil
}

// Удаление коллекции
func (s *CollectionService) Delete(ctx context.Context, id uint) error {
	return s.collectionRepo.Delete(ctx, id)
}

func (s *CollectionService) AddToPost(ctx context.Context, data admin.AddCollection) (model.Collection, error) {
	_, err := s.postCollectionRepo.Create(ctx, model.PostCollection{
		PostId:       data.PostID,
		CollectionId: data.CollectionID,
	})
	return model.Collection{}, err
}

func (s *CollectionService) DeleteCollectionPost(ctx context.Context, data admin.DeleteCollectionPost) (model.Collection, error) {
	err := s.postCollectionRepo.DeleteByPostAndCollectionId(ctx, data.PostID, data.CollectionID)
	return model.Collection{}, err
}
