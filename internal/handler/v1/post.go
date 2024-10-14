package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
	"work-project/internal/middleware"
	"work-project/internal/model"
	"work-project/internal/schema"
)

func (h *Handler) initPost(v1 *gin.RouterGroup) {
	v1.GET(
		"/post",
		middleware.GinErrorHandle(h.GetListingPosts),
	)
	v1.POST(
		"/post/read",
		middleware.GinErrorHandle(h.ReadPost),
	)
	v1.POST(
		"/post/save-quiz",
		middleware.GinErrorHandle(h.SaveQuizPoints),
	)
	v1.GET(
		"/post/archive",
		middleware.GinErrorHandle(h.GetArchivePosts),
	)
	v1.GET(
		"/post/check-quiz",
		middleware.GinErrorHandle(h.CheckQuiz),
	)
	v1.GET(
		"/post/filter",
		middleware.GinErrorHandle(h.GetFilterPosts),
	)
}

// GetListingPosts
// WhoAmi godoc
// @Summary список группированных постов
// @Accept json
// @Produce json
// @Success 200 {object} schema.Response[schema.PostResponseByGroup]
// @Failure 400 {object} schema.Response[schema.Empty]
// @Param hashtag_id query string false "hashtag_id"
// @Param collection_id query string false "collection_id"
// @Param language query string true "language"
// @tags post
// @Router /api/v1/post [get]
func (h *Handler) GetListingPosts(c *gin.Context) error {
	ctx := c.Request.Context()
	var userId *string
	token := c.GetHeader("Authorization")
	language := c.Query("language")
	hashtagIDsStr := c.Query("hashtag_id")
	hashtagIds := make([]uint, 0)
	for _, msi := range strings.Split(hashtagIDsStr, ",") {
		if msi == "" {
			continue
		}
		id, _ := strconv.ParseUint(msi, 10, 64)
		hashtagIds = append(hashtagIds, uint(id))
	}
	collectionIDsStr := c.Query("collection_id")
	collectionIds := make([]uint, 0)
	for _, msi := range strings.Split(collectionIDsStr, ",") {
		if msi == "" {
			continue
		}
		id, _ := strconv.ParseUint(msi, 10, 64)
		collectionIds = append(collectionIds, uint(id))
	}

	if token != "" {
		userIdStr, err := h.services.Auth.VerifyToken(token)
		if err != nil {
			return err
		}
		userId = &userIdStr
	}

	posts, err := h.usecases.Post.GetListingWithGroup(ctx, userId, hashtagIds, collectionIds, language)
	if err != nil {
		return err
	}

	return schema.Respond(posts, c)
}

// ReadPost
// WhoAmi godoc
// @Summary прочесть книгу
// @Accept json
// @Produce json
// @Success 200 {object} schema.Response[model.UserPost]
// @Failure 400 {object} schema.Response[schema.Empty]
// @Security BearerAuth
// @Param data body schema.ReadPost true "post"
// @tags post
// @Router /api/v1/post/read [post]
func (h *Handler) ReadPost(c *gin.Context) error {
	ctx := c.Request.Context()
	token := c.GetHeader("Authorization")
	userId, err := h.services.Auth.VerifyToken(token)
	if err != nil {
		return err
	}
	var data model.UserPost
	if err = c.Bind(&data); err != nil {
		return err
	}
	data.UserId = userId

	posts, err := h.usecases.Post.ReadPost(ctx, data)
	if err != nil {
		return err
	}

	return schema.Respond(posts, c)
}

// SaveQuizPoints
// WhoAmi godoc
// @Summary сохраниь что квиз прочитан
// @Accept json
// @Produce json
// @Success 200 {object} schema.Response[model.UserPost]
// @Failure 400 {object} schema.Response[schema.Empty]
// @Security BearerAuth
// @Param data body schema.PassQuizPost true "post"
// @tags post
// @Router /api/v1/post/save-quiz [post]
func (h *Handler) SaveQuizPoints(c *gin.Context) error {
	ctx := c.Request.Context()
	token := c.GetHeader("Authorization")
	userId, err := h.services.Auth.VerifyToken(token)
	if err != nil {
		return err
	}
	var data model.UserPost
	if err := c.Bind(&data); err != nil {
		return err
	}
	data.UserId = userId

	posts, err := h.usecases.Post.SaveQuizPoints(ctx, data)
	if err != nil {
		return err
	}

	return schema.Respond(posts, c)
}

// GetArchivePosts
// WhoAmi godoc
// @Summary архивные посты
// @Accept json
// @Produce json
// @Success 200 {object} schema.Response[[]schema.PostResponse]
// @Failure 400 {object} schema.Response[schema.Empty]
// @Security BearerAuth
// @tags post
// @Router /api/v1/post/archive [get]
func (h *Handler) GetArchivePosts(c *gin.Context) error {
	ctx := c.Request.Context()
	token := c.GetHeader("Authorization")
	userId, err := h.services.Auth.VerifyToken(token)
	if err != nil {
		return err
	}

	language := c.Query("language")
	posts, err := h.usecases.Post.GetArchive(ctx, userId, language)
	if err != nil {
		return err
	}

	return schema.Respond(posts, c)
}

// CheckQuiz
// WhoAmi godoc
// @Summary проверить что квиз прочитан
// @Accept json
// @Produce json
// @Success 200 {object} schema.Response[model.UserPost]
// @Failure 400 {object} schema.Response[schema.Empty]
// @Security BearerAuth
// @Param post_id query int true "post_id"
// @tags post
// @Router /api/v1/post/check-quiz [get]
func (h *Handler) CheckQuiz(c *gin.Context) error {
	ctx := c.Request.Context()
	token := c.GetHeader("Authorization")
	userId, err := h.services.Auth.VerifyToken(token)
	if err != nil {
		return err
	}
	postIdStr := c.Query("post_id")
	postId, err := strconv.ParseUint(postIdStr, 10, 64)
	if err != nil {
		return err
	}

	posts, err := h.usecases.Post.CheckQuiz(ctx, userId, uint(postId))
	if err != nil {
		return err
	}

	return schema.Respond(posts, c)
}

// GetFilterPosts
// WhoAmi godoc
// @Summary список постов с фильтром
// @Accept json
// @Produce json
// @Success 200 {object} schema.Response[[]schema.PostResponse]
// @Failure 400 {object} schema.Response[schema.Empty]
// @Security BearerAuth
// @Param hashtag_id query string false "hashtag_id"
// @Param collection_id query string false "collection_id"
// @Param language query string true "language"
// @Param search query string false "search"
// @Param search post_id int false "post_id"
// @Param post_type query string true "all, post, partner"
// @tags post
// @Router /api/v1/post/filter [get]
func (h *Handler) GetFilterPosts(c *gin.Context) error {
	ctx := c.Request.Context()
	var userId *string
	var postId *uint
	token := c.GetHeader("Authorization")
	language := c.Query("language")
	postIdStr := c.Query("post_id")
	if len(c.Query("post_id")) > 0 {
		tempPostId, err := strconv.ParseUint(postIdStr, 10, 64)
		if err != nil {
			return err
		}
		tempPostIdUint := uint(tempPostId)
		postId = &tempPostIdUint
	}

	hashtagIDsStr := c.Query("hashtag_id")
	hashtagIds := make([]uint, 0)
	for _, msi := range strings.Split(hashtagIDsStr, ",") {
		if msi == "" {
			continue
		}
		id, _ := strconv.ParseUint(msi, 10, 64)
		hashtagIds = append(hashtagIds, uint(id))
	}
	collectionIDsStr := c.Query("collection_id")
	collectionIds := make([]uint, 0)
	for _, msi := range strings.Split(collectionIDsStr, ",") {
		if msi == "" {
			continue
		}
		id, _ := strconv.ParseUint(msi, 10, 64)
		collectionIds = append(collectionIds, uint(id))
	}

	search := c.Query("search")
	postType := c.Query("post_type")
	if postType != "all" && postType != "post" && postType != "partner" {
		return fmt.Errorf("incorrect post_type value")
	}

	if token != "" {
		userIdStr, err := h.services.Auth.VerifyToken(token)
		if err != nil {
			return err
		}
		userId = &userIdStr
	}

	posts, err := h.usecases.Post.GetListing(ctx, userId, hashtagIds, collectionIds, postType, search, language, postId)
	if err != nil {
		return err
	}

	return schema.Respond(posts, c)
}
