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
	SaveQuizPoints(ctx context.Context, data model.UserPost) (model.UserPost, error)
	GetArchive(ctx context.Context, userId string) ([]model.Post, error)
}

type PostUsecase struct {
	postService     service.Post
	markService     service.Mark
	userPostService service.UserPost
	balanceService  service.Balance
}

func NewPostUsecase(services *service.Services) *PostUsecase {
	return &PostUsecase{
		postService:     services.Post,
		markService:     services.Mark,
		userPostService: services.UserPost,
		balanceService:  services.Balance,
	}
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

	userPosts, err := u.userPostService.GetAllByUser(ctx, *userId)
	if err != nil {
		return nil, err
	}
	postIdRead := make(map[uint]bool)
	for _, up := range userPosts {
		postIdRead[up.PostId] = true
	}

	for i := range posts {
		um, exists := postIdMark[posts[i].PostID]
		posts[i].IsMarked = exists
		posts[i].MarkId = &um.MarkID

		_, exists = postIdRead[posts[i].PostID]
		posts[i].IsAlreadyRead = exists
	}

	return posts, nil
}

func (u *PostUsecase) SaveQuizPoints(ctx context.Context, data model.UserPost) (model.UserPost, error) {
	up, err := u.userPostService.AddQuizPoints(ctx, data)
	if err != nil {
		return model.UserPost{}, err
	}
	tr := model.Transaction{
		UserId:            data.UserId,
		TransactionType:   string(model.TRANSACTION_TYPE_INCOME),
		TransactionReason: string(model.TRANSACTION_REASON_POST),
	}
	if up.QuizPoints != nil {
		tr.Coins = *up.QuizPoints
	}
	if up.QuizSapphires != nil {
		tr.Sapphires = *up.QuizSapphires
	}

	_, err = u.balanceService.CreateTransaction(ctx, data.UserId, tr)
	if err != nil {
		return model.UserPost{}, err
	}
	return up, nil
}

func (u *PostUsecase) GetArchive(ctx context.Context, userId string) ([]model.Post, error) {
	userPosts, err := u.userPostService.GetAllByUser(ctx, userId)
	if err != nil {
		return nil, err
	}
	postIds := make([]uint, 0)
	for _, up := range userPosts {
		postIds = append(postIds, up.PostId)
	}
	return u.postService.GetByIds(ctx, postIds)
}
