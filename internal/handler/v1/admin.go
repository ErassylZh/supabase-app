package v1

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"work-project/internal/admin"
	"work-project/internal/middleware"
	"work-project/internal/schema"
)

func (h *Handler) initAdmin(v1 *gin.RouterGroup) {
	group := v1.Group("/admin")
	group.POST("/post", middleware.GinErrorHandle(h.CreatePost))
	group.GET("/post/all", middleware.GinErrorHandle(h.GetPostList))
	group.GET("/post", middleware.GinErrorHandle(h.GetPostList))
	group.PUT("/post", middleware.GinErrorHandle(h.UpdatePost))
	group.DELETE("/post", middleware.GinErrorHandle(h.DeletePost))

}

// CreatePost
// WhoAmi godoc
// @Accept json
// @Produce json
// @Param data body admin.CreatePost true "CreatePost data"
// @Success 200 {object} schema.Response[model.Post]
// @Failure 400 {object} schema.Response[schema.Empty]
// @Security BearerAuth
// @tags publication
// @Router /api/v1/admin/post [post]
func (h *Handler) CreatePost(c *gin.Context) error {
	ctx := c.Request.Context()

	var data admin.CreatePost
	err := c.Bind(&data)
	if err != nil {
		return err
	}

	post, err := h.services.Post.Create(ctx, data)
	if err != nil {
		return err
	}

	return schema.Respond(post, c)
}

// UpdatePost
// WhoAmi godoc
// @Accept json
// @Produce json
// @Param data body schema.UpdatePost true "UserLogin data"
// @Success 200 {object} schema.Response[model.Post]
// @Failure 400 {object} schema.Response[schema.Empty]
// @Security BearerAuth
// @tags publication
// @Router /api/v1/publication [put]
func (h *Handler) UpdatePost(c *gin.Context) error {
	ctx := c.Request.Context()

	var data admin.UpdatePost
	err := c.Bind(&data)
	if err != nil {
		return err
	}

	publication, err := h.services.Post.Update(ctx, data)
	if err != nil {
		return err
	}

	return schema.Respond(publication, c)
}

// DeletePost
// WhoAmi godoc
// @Accept json
// @Produce json
// @Param post_id query int true "id"
// @Success 200 {object} schema.Response[schema.Empty]
// @Failure 400 {object} schema.Response[schema.Empty]
// @Security BearerAuth
// @tags publication
// @Router /api/v1/publication [delete]
func (h *Handler) DeletePost(c *gin.Context) error {
	ctx := c.Request.Context()

	postId, err := strconv.ParseUint(c.Query("post_id"), 10, 64)
	if err != nil {
		return err
	}

	err = h.services.Post.Delete(ctx, uint(postId))
	if err != nil {
		return err
	}

	return schema.Respond(schema.Empty{}, c)
}

// GetPostList
// WhoAmi godoc
// @Accept json
// @Produce json
// @Param hashtag_id query string false "hashtag_id"
// @Param collection_id query string false "collection_id"
// @Param language query string true "language"
// @Param search query string false "search"
// @Param post_ids query string false "post_ids"
// @Param post_type query string true "all, post, partner"
// @Success 200 {object} schema.Response[[]model.Post]
// @Failure 400 {object} schema.Response[schema.Empty]
// @tags publication
// @Router /api/v1/publication/all [get]
func (h *Handler) GetPostList(c *gin.Context) error {
	ctx := c.Request.Context()

	publications, err := h.services.Post.GetListing(ctx)
	if err != nil {
		return err
	}

	return schema.Respond(publications, c)
}
