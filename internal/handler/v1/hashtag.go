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

// Получение хэштега по ID
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

// Обновление хэштега
func (h *Handler) UpdateHashtag(c *gin.Context) error {
	ctx := c.Request.Context()

	id, err := strconv.Atoi(c.Query("hashtag_id"))
	if err != nil {
		return err
	}

	var data admin.UpdateHashtag
	err = c.Bind(&data)
	if err != nil {
		return err
	}
	data.HashtagID = uint(id)

	hashtag, err := h.services.Hashtag.Update(ctx, data)
	if err != nil {
		return err
	}

	return schema.Respond(hashtag, c)
}

// Удаление хэштега
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
// @Param language query string true "language"
// @tags collection
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
// @Success 200 {object} schema.Response[model.Collection]
// @Failure 400 {object} schema.Response[schema.Empty]
// @Param language query string true "language"
// @tags collection
// @Router /api/v1/hashtag/delete-post [delete]
func (h *Handler) DeleteHashtagToPost(c *gin.Context) error {
	ctx := c.Request.Context()

	var data admin.DeleteHashtagPost
	if err := c.BindJSON(&data); err != nil {
		return err
	}

	collections, err := h.services.Hashtag.DeleteCollectionPost(ctx, data)
	if err != nil {
		return err
	}
	return schema.Respond(collections, c)
}
