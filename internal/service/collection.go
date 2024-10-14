package service

import (
	"context"
	"work-project/internal/repository"
	"work-project/internal/schema"
)

type Collection interface {
	GetAllCollection(ctx context.Context, language string) ([]schema.CollectionListResponse, error)
	GetAllRecommendation(ctx context.Context, language string) ([]schema.CollectionListResponse, error)
}

type CollectionService struct {
	collectionRepo repository.Collection
	postService    Post
}

func NewCollectionService(collectionRepo repository.Collection, postService Post) *CollectionService {
	return &CollectionService{collectionRepo: collectionRepo, postService: postService}
}

func (s *CollectionService) GetAllCollection(ctx context.Context, language string) ([]schema.CollectionListResponse, error) {
	collections, err := s.collectionRepo.GetAllCollection(ctx, language)
	if err != nil {
		return nil, err
	}

	result := make([]schema.CollectionListResponse, len(collections))
	for i, collection := range collections {
		posts, err := s.postService.GetListing(ctx, []uint{}, []uint{collection.CollectionID}, "", language, nil)
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

func (s *CollectionService) GetAllRecommendation(ctx context.Context, language string) ([]schema.CollectionListResponse, error) {
	collections, err := s.collectionRepo.GetAllRecommendation(ctx, language)
	if err != nil {
		return nil, err
	}

	result := make([]schema.CollectionListResponse, len(collections))
	for i, collection := range collections {
		posts, err := s.postService.GetListing(ctx, []uint{}, []uint{collection.CollectionID}, "", language, nil)
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
