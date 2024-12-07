package usecase

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"sort"
	"work-project/internal/model"
	"work-project/internal/schema"
	"work-project/internal/service"
)

type Post interface {
	GetListing(ctx context.Context, userId *string, postType string, filter schema.GetListingFilter) ([]schema.PostResponse, error)
	SaveQuizPoints(ctx context.Context, data model.UserPost) (model.UserPost, error)
	GetArchive(ctx context.Context, userId string, language string) ([]schema.ArchivePost, error)
	CheckQuiz(ctx context.Context, userId string, postId uint) (bool, error)
	GetListingWithGroup(ctx context.Context, userId *string, filter schema.GetListingFilter) (schema.PostResponseByGroup, error)
	ReadPost(ctx context.Context, post schema.ReadPostRequest) (model.UserPost, error)
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

func (u *PostUsecase) ReadPost(ctx context.Context, post schema.ReadPostRequest) (model.UserPost, error) {
	postAlreadyReaded, err := u.userPostService.GetByUserAndPost(ctx, post.UserId, post.PostId)
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return u.userPostService.Create(ctx, model.UserPost{
			PostId:  post.PostId,
			UserId:  post.UserId,
			ReadEnd: post.EndReading,
		})
	}
	if err != nil {
		return model.UserPost{}, err
	}

	if post.EndReading {
		postAlreadyReaded.ReadEnd = true
		return u.userPostService.Update(ctx, postAlreadyReaded)
	}
	return postAlreadyReaded, nil
}

func (u *PostUsecase) GetListing(ctx context.Context, userId *string, postType string, filter schema.GetListingFilter) ([]schema.PostResponse, error) {
	posts, err := u.postService.GetListing(ctx, filter)
	if err != nil {
		return nil, err
	}

	if userId != nil {

		userMarks, err := u.markService.FindByUserID(ctx, *userId)
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}

		postIdMark := make(map[uint]schema.MarkResponse)
		for _, um := range userMarks {
			postIdMark[um.PostID] = um
		}

		userPosts, err := u.userPostService.GetAllByUser(ctx, *userId)
		if err != nil {
			return nil, err
		}
		postIdRead := make(map[uint]bool)
		postIdPassed := make(map[uint]bool)
		for _, up := range userPosts {
			postIdRead[up.PostId] = true
			postIdPassed[up.PostId] = up.QuizPoints != nil || up.QuizSapphires != nil
		}

		for i := range posts {
			um, exists := postIdMark[posts[i].PostID]
			posts[i].IsMarked = exists
			posts[i].MarkId = &um.MarkID

			_, exists = postIdRead[posts[i].PostID]
			posts[i].IsAlreadyRead = exists
			_, exists = postIdPassed[posts[i].PostID]
			posts[i].QuizPassed = exists
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

func (u *PostUsecase) GetArchive(ctx context.Context, userId string, language string) ([]schema.ArchivePost, error) {
	userPosts, err := u.userPostService.GetAllByUser(ctx, userId)
	if err != nil {
		return nil, err
	}
	postIds := make([]uint, len(userPosts))
	for i, up := range userPosts {
		postIds[i] = up.PostId
	}
	posts, err := u.postService.GetByIds(ctx, postIds)
	if err != nil {
		return nil, err
	}

	userPostMap := make(map[uint]model.UserPost)
	for _, up := range userPosts {
		userPostMap[up.PostId] = up
	}

	sort.Slice(posts, func(i, j int) bool {
		return userPostMap[posts[i].PostID].CreatedAt.Before(userPostMap[posts[j].PostID].CreatedAt)
	})

	result := make([]schema.ArchivePost, len(posts))
	for i, post := range posts {
		result[i] = schema.ArchivePost{
			Post:            post,
			EarnedCoins:     userPostMap[post.PostID].QuizPoints,
			EarnedSapphires: userPostMap[post.PostID].QuizSapphires,
			QuizPassed:      userPostMap[post.PostID].QuizSapphires != nil || userPostMap[post.PostID].QuizPoints != nil,
		}
	}

	return result, nil
}

func (u *PostUsecase) CheckQuiz(ctx context.Context, userId string, postId uint) (bool, error) {
	userPost, err := u.userPostService.GetByUserAndPost(ctx, userId, postId)
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return true, nil
	}
	if err != nil {
		return false, err
	}
	return userPost.QuizSapphires == nil && userPost.QuizPoints == nil, nil
}

func (u *PostUsecase) GetListingWithGroup(ctx context.Context, userId *string, filter schema.GetListingFilter) (schema.PostResponseByGroup, error) {
	posts, err := u.postService.GetListing(ctx, filter)
	if err != nil {
		return schema.PostResponseByGroup{}, err
	}
	if userId == nil {
		result := schema.PostResponseByGroup{
			ContinueReading: []schema.PostResponse{},
		}
		for _, post := range posts {
			for _, hashtag := range post.Hashtags {
				if hashtag.Name == string(model.HASHTAG_NAME_BESTSELLER) {
					post.PostType = "post"
					result.Bestsellers = append(result.Bestsellers, post)
				}
				if hashtag.Name == string(model.HASHTAG_NAME_PARTNER) {
					post.PostType = "partner"
					result.Partners = append(result.Partners, post)
				}
			}
		}

		return result, nil
	}
	userPosts, err := u.userPostService.GetAllByUser(ctx, *userId)
	if err != nil {
		return schema.PostResponseByGroup{}, err
	}
	userPostMap := make(map[uint]model.UserPost)
	for _, up := range userPosts {
		userPostMap[up.PostId] = up
	}

	userMarks, err := u.markService.FindByUserID(ctx, *userId)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return schema.PostResponseByGroup{}, err
	}

	postIdMark := make(map[uint]schema.MarkResponse)
	for _, um := range userMarks {
		postIdMark[um.PostID] = um
	}

	for i := range posts {
		um, exists := postIdMark[posts[i].PostID]
		posts[i].IsMarked = exists
		posts[i].MarkId = &um.MarkID

		if up, upExists := userPostMap[posts[i].PostID]; upExists {
			posts[i].IsAlreadyRead = upExists
			posts[i].QuizPassed = up.QuizPoints != nil || up.QuizSapphires != nil
		}
	}

	result := schema.PostResponseByGroup{ContinueReading: []schema.PostResponse{}}
	for _, post := range posts {
		for _, hashtag := range post.Hashtags {
			if hashtag.Name == string(model.HASHTAG_NAME_BESTSELLER) {
				post.PostType = "post"
				result.Bestsellers = append(result.Bestsellers, post)
			}
			if hashtag.Name == string(model.HASHTAG_NAME_PARTNER) {
				post.PostType = "partner"
				result.Partners = append(result.Partners, post)
			}
		}

		if up, exists := userPostMap[post.PostID]; exists && !up.ReadEnd {
			result.ContinueReading = append(result.ContinueReading, post)
		}
	}

	return result, nil
}
