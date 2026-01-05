package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
	"work-project/internal/middleware"
	"work-project/internal/schema"
)

func (h *Handler) initPostV2(v2 *gin.RouterGroup) {
	v2.GET(
		"/post",
		middleware.GinErrorHandle(h.GetListingPostsV2),
	)
	v2.GET(
		"/post/filter",
		middleware.GinErrorHandle(h.GetFilterPostsV2),
	)
}

// GetListingPostsV2
// WhoAmi godoc
// @Summary список группированных постов
// @Accept json
// @Produce json
// @Success 200 {object} schema.Response[schema.PostResponseByGroup]
// @Failure 400 {object} schema.Response[schema.Empty]
// @Param hashtag_id query int false "hashtag_id"
// @Param collection_id query int false "collection_id"
// @Param language query string true "language"
// @Param page query int true "page"
// @Param size query int true "size"
// @tags post
// @Router /api/v2/post [get]
func (h *Handler) GetListingPostsV2(c *gin.Context) error {
	ctx := c.Request.Context()

	var pagination schema.Pagination
	if err := c.BindQuery(&pagination); err != nil {
		return err
	}
	pagination.Validate()

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

	posts, total, err := h.usecases.Post.GetListingWithGroup(ctx, userId, schema.GetListingFilter{
		HashtagIds:    hashtagIds,
		CollectionIds: collectionIds,
		Language:      &language,
		Pagination:    pagination,
	})
	if err != nil {
		return err
	}

	return schema.RespondPaginate(posts, total, pagination.Page, c)
}

// GetFilterPostsV2
// WhoAmi godoc
// @Summary список постов с фильтром
// @Accept json
// @Produce json
// @Success 200 {object} schema.Response[[]schema.PostResponse]
// @Failure 400 {object} schema.Response[schema.Empty]
// @Param hashtag_id query string false "hashtag_id"
// @Param collection_id query string false "collection_id"
// @Param language query string true "language"
// @Param search query string false "search"
// @Param post_ids query string false "post_ids"
// @Param post_type query string true "all, post, partner"
// @Param page query int true "page"
// @Param size query int true "size"
// @tags post
// @Router /api/v2/post/filter [get]
func (h *Handler) GetFilterPostsV2(c *gin.Context) error {
	ctx := c.Request.Context()
	var userId *string
	var postIds []uint

	var pagination schema.Pagination
	if err := c.BindQuery(&pagination); err != nil {
		return err
	}
	pagination.Validate()

	token := c.GetHeader("Authorization")
	language := c.Query("language")
	postIdsStr := c.Query("post_ids")
	for _, msi := range strings.Split(postIdsStr, ",") {
		if msi == "" {
			continue
		}
		id, _ := strconv.ParseUint(msi, 10, 64)
		postIds = append(postIds, uint(id))
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

	posts, total, err := h.usecases.Post.GetListing(ctx, userId, postType, schema.GetListingFilter{
		Search:        &search,
		HashtagIds:    hashtagIds,
		CollectionIds: collectionIds,
		Language:      &language,
		PostIds:       postIds,
		Pagination:    pagination,
	})
	if err != nil {
		return err
	}

	return schema.RespondPaginate(posts, total, pagination.Page, c)
}
