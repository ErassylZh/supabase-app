package v1

import (
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
		middleware.GinErrorHandle(h.GetListingPosts),
	)
	v1.GET(
		"/post/check-quiz",
		middleware.GinErrorHandle(h.SaveQuizPoints),
	)
	v1.GET(
		"/post/filter",
		middleware.GinErrorHandle(h.GetFilterPosts),
	)
}

func (h *Handler) GetListingPosts(c *gin.Context) error {
	ctx := c.Request.Context()
	var userId *string
	token := c.GetHeader("Authorization")
	hashtagIDsStr := c.Query("hashtag_id")
	hashtagIds := make([]uint, 0)
	for _, msi := range strings.Split(hashtagIDsStr, ",") {
		id, _ := strconv.ParseUint(msi, 10, 64)
		hashtagIds = append(hashtagIds, uint(id))
	}

	if token != "" {
		userIdStr, err := h.services.Auth.VerifyToken(token)
		if err != nil {
			return err
		}
		userId = &userIdStr
	}

	posts, err := h.usecases.Post.GetListing(ctx, userId, hashtagIds)
	if err != nil {
		return err
	}

	return schema.Respond(posts, c)
}

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

	posts, err := h.services.UserPost.Create(ctx, data)
	if err != nil {
		return err
	}

	return schema.Respond(posts, c)
}

func (h *Handler) SaveQuizPoints(c *gin.Context) error {
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

	posts, err := h.usecases.Post.SaveQuizPoints(ctx, data)
	if err != nil {
		return err
	}

	return schema.Respond(posts, c)
}

func (h *Handler) GetArchivePosts(c *gin.Context) error {
	ctx := c.Request.Context()
	token := c.GetHeader("Authorization")
	userId, err := h.services.Auth.VerifyToken(token)
	if err != nil {
		return err
	}

	posts, err := h.usecases.Post.GetArchive(ctx, userId)
	if err != nil {
		return err
	}

	return schema.Respond(posts, c)
}

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

func (h *Handler) GetFilterPosts(c *gin.Context) error {
	ctx := c.Request.Context()
	var userId *string
	token := c.GetHeader("Authorization")
	hashtagIDsStr := c.Query("hashtag_id")
	hashtagIds := make([]uint, 0)
	for _, msi := range strings.Split(hashtagIDsStr, ",") {
		id, _ := strconv.ParseUint(msi, 10, 64)
		hashtagIds = append(hashtagIds, uint(id))
	}

	if token != "" {
		userIdStr, err := h.services.Auth.VerifyToken(token)
		if err != nil {
			return err
		}
		userId = &userIdStr
	}

	posts, err := h.usecases.Post.GetListing(ctx, userId, hashtagIds)
	if err != nil {
		return err
	}

	return schema.Respond(posts, c)
}
