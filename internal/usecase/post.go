package usecase

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"work-project/internal/model"
	"work-project/internal/schema"
	"work-project/internal/service"
)

type Post interface {
	GetListing(ctx context.Context, userId *string) ([]schema.PostResponse, error)
}

type PostUsecase struct {
	postService service.Post
	markService service.Mark
}

func NewPostUsecase(postService service.Post, markService service.Mark) *PostUsecase {
	return &PostUsecase{postService: postService, markService: markService}
}

func (u *PostUsecase) GetListing(ctx context.Context, userId *string) ([]schema.PostResponse, error) {
	posts, err := u.postService.GetListing(ctx, userId)
	if err != nil {
		return nil, err
	}

	if userId == nil {
		return posts, nil
	}

	userMarks, err := u.markService.FindByUserID(ctx, *userId)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	postIdMark := make(map[uint]model.Mark)
	for _, um := range userMarks {
		postIdMark[um.PostID] = um
	}
	for i := range posts {
		um, exists := postIdMark[posts[i].PostID]
		posts[i].IsMarked = exists
		posts[i].MarkId = &um.MarkID
	}

	return posts, nil
}
