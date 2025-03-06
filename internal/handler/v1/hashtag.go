package v1

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"work-project/internal/admin"
	"work-project/internal/middleware"
	"work-project/internal/schema"
)

func (h *Handler) initHashtag(v1 *gin.RouterGroup) {
	v1.GET(
		"/hashtag",
		middleware.GinErrorHandle(h.GetAllHashtags),
	)
	v1.POST(
		"/hashtag",
		middleware.GinErrorHandle(h.CreateHashtag),
	)
	v1.GET(
		"/hashtag/id",
		middleware.GinErrorHandle(h.GetHashtagByID),
	)
	v1.PUT(
		"/hashtag",
		middleware.GinErrorHandle(h.UpdateHashtag),
	)
	v1.DELETE(
		"/hashtag",
		middleware.GinErrorHandle(h.DeleteHashtag),
	)
	v1.POST(
		"/hashtag/add",
		middleware.GinErrorHandle(h.AddHashtagToPost),
	)
	v1.DELETE(
		"/hashtag/delete-post",
		middleware.GinErrorHandle(h.DeleteHashtagToPost),
	)
}

// GetAllHashtags
// WhoAmi godoc
// @Summary получить все коллекций
// @Accept json
// @Produce json
// @Success 200 {object} schema.Response[[]model.Hashtag]
// @Failure 400 {object} schema.Response[schema.Empty]
// @tags hashtag
// @Router /api/v1/hashtag [get]
func (h *Handler) GetAllHashtags(c *gin.Context) error {
	ctx := c.Request.Context()
	hashtags, err := h.services.Hashtag.GetAll(ctx)
	if err != nil {
		return err
	}
	return schema.Respond(hashtags, c)
}

// CreateHashtag
// WhoAmi godoc
// @Accept json
// @Produce json
// @Param data body admin.CreateHashtag true "CreatePost data"
// @Success 200 {object} schema.Response[model.Hashtag]
// @Failure 400 {object} schema.Response[schema.Empty]
// @Security BearerAuth
// @tags publication
// @Router /api/v1/hashtag [post]
func (h *Handler) CreateHashtag(c *gin.Context) error {
	ctx := c.Request.Context()

	var data admin.CreateHashtag
	err := c.Bind(&data)
	if err != nil {
		return err
	}

	hashtag, err := h.services.Hashtag.Create(ctx, data)
	if err != nil {
		return err
	}

	return schema.Respond(hashtag, c)
}

// GetHashtagByID
// WhoAmi godoc
// @Accept json
// @Produce json
// @Param hashtag_id query int true "hashtag_id"
// @Success 200 {object} schema.Response[model.Hashtag]
// @Failure 400 {object} schema.Response[schema.Empty]
// @tags collection
// @Router /api/v1/hashtag/id [get]
func (h *Handler) GetHashtagByID(c *gin.Context) error {
	ctx := c.Request.Context()

	id, err := strconv.Atoi(c.Query("hashtag_id"))
	if err != nil {
		return err
	}

	hashtag, err := h.services.Hashtag.GetByID(ctx, uint(id))
	if err != nil {
		return err
	}

	return schema.Respond(hashtag, c)
}

// UpdateHashtag
// WhoAmi godoc
// @Accept json
// @Produce json
// @Param data body admin.UpdateHashtag true "CreatePost data"
// @Success 200 {object} schema.Response[model.Collection]
// @Failure 400 {object} schema.Response[schema.Empty]
// @tags hashtag
// @Router /api/v1/hashtag [put]
func (h *Handler) UpdateHashtag(c *gin.Context) error {
	ctx := c.Request.Context()

	var data admin.UpdateHashtag
	err := c.Bind(&data)
	if err != nil {
		return err
	}

	hashtag, err := h.services.Hashtag.Update(ctx, data)
	if err != nil {
		return err
	}

	return schema.Respond(hashtag, c)
}

// DeleteHashtag
// WhoAmi godoc
// @Accept json
// @Produce json
// @Param hashtag_id query int false "hashtag_id"
// @Success 200 {object} schema.Response[schema.Empty]
// @Failure 400 {object} schema.Response[schema.Empty]
// @tags hashtag
// @Router /api/v1/hashtag [delete]
func (h *Handler) DeleteHashtag(c *gin.Context) error {
	ctx := c.Request.Context()

	id, err := strconv.Atoi(c.Query("hashtag_id"))
	if err != nil {
		return err
	}

	err = h.services.Hashtag.Delete(ctx, uint(id))
	if err != nil {
		return err
	}

	return schema.Respond(schema.Empty{}, c)
}

// AddHashtagToPost
// WhoAmi godoc
// @Summary добавить пост в коллекцию
// @Accept json
// @Produce json
// @Success 200 {object} schema.Response[model.Hashtag]
// @Failure 400 {object} schema.Response[schema.Empty]
// @Param data body admin.AddHashtag true "CreatePost data"
// @tags hashtag
// @Router /api/v1/hashtag/add [post]
func (h *Handler) AddHashtagToPost(c *gin.Context) error {
	ctx := c.Request.Context()

	var data admin.AddHashtag
	if err := c.BindJSON(&data); err != nil {
		return err
	}

	hashtag, err := h.services.Hashtag.AddToPost(ctx, data)
	if err != nil {
		return err
	}
	return schema.Respond(hashtag, c)
}

// DeleteHashtagToPost
// WhoAmi godoc
// @Summary удалить пост из коллекцию
// @Accept json
// @Produce json
// @Success 200 {object} schema.Response[model.Hashtag]
// @Failure 400 {object} schema.Response[schema.Empty]
// @Param data body admin.DeleteHashtagPost true "CreatePost data"
// @tags hashtag
// @Router /api/v1/hashtag/delete-post [delete]
func (h *Handler) DeleteHashtagToPost(c *gin.Context) error {
	ctx := c.Request.Context()

	var data admin.DeleteHashtagPost
	if err := c.BindJSON(&data); err != nil {
		return err
	}

	hashtags, err := h.services.Hashtag.DeleteHashtagPost(ctx, data)
	if err != nil {
		return err
	}
	return schema.Respond(hashtags, c)
}
