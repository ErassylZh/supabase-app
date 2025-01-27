package service

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"work-project/internal/model"
	"work-project/internal/repository"
	"work-project/internal/schema"
)

type Collection interface {
	GetAllCollection(ctx context.Context, language string, userId *string, withoutPosts bool) ([]schema.CollectionListResponse, error)
	GetAllRecommendation(ctx context.Context, language string) ([]schema.CollectionListResponse, error)
}

type CollectionService struct {
	collectionRepo repository.Collection
	userPostRepo   repository.UserPost
	markRepo       repository.Mark
	postService    Post
}

func NewCollectionService(collectionRepo repository.Collection, userPostRepo repository.UserPost, postService Post, markRepo repository.Mark) *CollectionService {
	return &CollectionService{collectionRepo: collectionRepo, postService: postService, markRepo: markRepo, userPostRepo: userPostRepo}
}

func (s *CollectionService) GetAllCollection(ctx context.Context, language string, userId *string, withoutPosts bool) ([]schema.CollectionListResponse, error) {
	collections, err := s.collectionRepo.GetAllCollection(ctx, language, withoutPosts)
	if err != nil {
		return nil, err
	}

	postIdMark := make(map[uint]model.Mark)
	userPostMap := make(map[uint]model.UserPost)
	if userId != nil {
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
		posts, err := s.postService.GetListing(ctx, schema.GetListingFilter{
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
