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
	GetListing(ctx context.Context, userId *string, hashtagIds []uint, collectionIds []uint, postType string, search string, language string) ([]schema.PostResponse, error)
	SaveQuizPoints(ctx context.Context, data model.UserPost) (model.UserPost, error)
	GetArchive(ctx context.Context, userId string, language string) ([]model.Post, error)
	CheckQuiz(ctx context.Context, userId string, postId uint) (bool, error)
	GetListingWithGroup(ctx context.Context, userId *string, hashtagIds []uint, collectionIds []uint, language string) (schema.PostResponseByGroup, error)
	ReadPost(ctx context.Context, post model.UserPost) (model.UserPost, error)
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

func (u *PostUsecase) ReadPost(ctx context.Context, post model.UserPost) (model.UserPost, error) {
	postAlreadyReaded, err := u.userPostService.GetByUserAndPost(ctx, post.UserId, post.PostId)
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return u.userPostService.Create(ctx, post)
	}
	if err != nil {
		return model.UserPost{}, err
	}
	return postAlreadyReaded, nil
}

func (u *PostUsecase) GetListing(ctx context.Context, userId *string, hashtagIds []uint, collectionIds []uint, postType string, search string, language string) ([]schema.PostResponse, error) {
	posts, err := u.postService.GetListing(ctx, hashtagIds, collectionIds, search, language)
	if err != nil {
		return nil, err
	}

	if userId != nil {

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
	}
	if postType == "all" {
		return posts, nil
	}

	result := make([]schema.PostResponse, 0)

	for _, post := range posts {
		isPartner := false
		for _, ht := range post.Hashtags {
			if ht.Name == string(model.HASHTAG_NAME_PARTNER) {
				isPartner = true
				break
			}
		}
		if (postType == "partner" && isPartner) || (postType == "post" && !isPartner) {
			result = append(result, post)
		}
	}

	return result, nil
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

func (u *PostUsecase) GetArchive(ctx context.Context, userId string, language string) ([]model.Post, error) {
	userPosts, err := u.userPostService.GetAllByUser(ctx, userId)
	if err != nil {
		return nil, err
	}
	postIds := make([]uint, len(userPosts))
	for i, up := range userPosts {
		postIds[i] = up.PostId
	}
	return u.postService.GetByIds(ctx, postIds)
}

func (u *PostUsecase) CheckQuiz(ctx context.Context, userId string, postId uint) (bool, error) {
	userPost, err := u.userPostService.GetByUserAndPost(ctx, userId, postId)
	if err != nil {
		return false, err
	}
	return userPost.QuizSapphires == nil && userPost.QuizPoints == nil, nil
}

func (u *PostUsecase) GetListingWithGroup(ctx context.Context, userId *string, hashtagIds []uint, collectionIds []uint, language string) (schema.PostResponseByGroup, error) {
	posts, err := u.postService.GetListing(ctx, hashtagIds, collectionIds, "", language)
	if err != nil {
		return schema.PostResponseByGroup{}, err
	}
	if userId == nil {
		result := schema.PostResponseByGroup{}
		for _, post := range posts {
			postAlreadyAdded := false
			for _, hashtag := range post.Hashtags {
				if hashtag.Name == string(model.HASHTAG_NAME_BESTSELLER) {
					postAlreadyAdded = true
					result.Bestsellers = append(result.Bestsellers, post)
				}
				if hashtag.Name == string(model.HASHTAG_NAME_PARTNER) {
					postAlreadyAdded = true
					result.Partners = append(result.Partners, post)
				}
			}
			if !postAlreadyAdded {
				result.Other = append(result.Other, post)
			}
		}

		return result, nil
	}

	userMarks, err := u.markService.FindByUserID(ctx, *userId)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return schema.PostResponseByGroup{}, err
	}

	postIdMark := make(map[uint]model.Mark)
	for _, um := range userMarks {
		postIdMark[um.PostID] = um
	}

	userPosts, err := u.userPostService.GetAllByUser(ctx, *userId)
	if err != nil {
		return schema.PostResponseByGroup{}, err
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

	result := schema.PostResponseByGroup{}
	for _, post := range posts {
		postAlreadyAdded := false
		for _, hashtag := range post.Hashtags {
			if hashtag.Name == string(model.HASHTAG_NAME_BESTSELLER) {
				postAlreadyAdded = true
				result.Bestsellers = append(result.Bestsellers, post)
			}
			if hashtag.Name == string(model.HASHTAG_NAME_PARTNER) {
				postAlreadyAdded = true
				result.Partners = append(result.Partners, post)
			}
		}
		if !postAlreadyAdded {
			result.Other = append(result.Other, post)
		}
	}

	return result, nil
}
