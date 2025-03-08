package v1

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"work-project/internal/admin"
	"work-project/internal/middleware"
	"work-project/internal/schema"
)

func (h *Handler) initStories(v1 *gin.RouterGroup) {
	v1.GET(
		"/stories",
		middleware.GinErrorHandle(h.GetActiveStories),
	)
	v1.POST(
		"/stories",
		middleware.GinErrorHandle(h.ReadStoriesByUser),
	)
	v1.GET(
		"/stories/all",
		middleware.GinErrorHandle(h.GetAllStories),
	)
	v1.GET(
		"/stories/id",
		middleware.GinErrorHandle(h.GetStoriesBuId),
	)
	v1.POST(
		"/stories/create",
		middleware.GinErrorHandle(h.CreateStory),
	)
	v1.DELETE(
		"/stories",
		middleware.GinErrorHandle(h.DeleteStoriesById),
	)
	v1.PUT(
		"/stories",
		middleware.GinErrorHandle(h.UpdateStory),
	)
	v1.POST(
		"/story-page",
		middleware.GinErrorHandle(h.CreateStoryPage),
	)
	v1.DELETE(
		"/story-page",
		middleware.GinErrorHandle(h.CreateStoryPage),
	)
	v1.PUT(
		"/story-page",
		middleware.GinErrorHandle(h.UpdateStoryPage),
	)
}

// ReadStoriesByUser
// WhoAmi godoc
// @Summary пометить сторис прочитанным
// @Accept json
// @Produce json
// @Success 200 {object} schema.Response[schema.Empty]
// @Failure 400 {object} schema.Response[schema.Empty]
// @Param story_page_id query int true "story_page_id"
// @Security BearerAuth
// @tags stories
// @Router /api/v1/stories [post]
func (h *Handler) ReadStoriesByUser(c *gin.Context) error {
	ctx := c.Request.Context()
	token := c.GetHeader("Authorization")
	userId, err := h.services.Auth.VerifyToken(token)
	if err != nil {
		return err
	}

	storiesId, err := strconv.ParseUint(c.Query("story_page_id"), 10, 64)
	if err != nil {
		return err
	}

	err = h.services.Stories.ReadStory(ctx, userId, uint(storiesId))
	if err != nil {
		return err
	}
	return schema.Respond(schema.Empty{}, c)
}

// GetActiveStories
// WhoAmi godoc
// @Summary получить список сторисов
// @Accept json
// @Produce json
// @Success 200 {object} schema.Response[[]model.Stories]
// @Failure 400 {object} schema.Response[schema.Empty]
// @Security BearerAuth
// @tags stories
// @Router /api/v1/stories [get]
func (h *Handler) GetActiveStories(c *gin.Context) error {
	ctx := c.Request.Context()

	token := c.GetHeader("Authorization")
	userId, _ := h.services.Auth.VerifyToken(token)
	response, err := h.services.Stories.GetByUserId(ctx, userId)
	if err != nil {
		return err
	}
	return schema.Respond(response, c)
}

// CreateStory
// WhoAmi godoc
// @Summary создать сторис
// @Accept json
// @Produce json
// @Param data body admin.CreateStories true "CreateProduct data"
// @Success 200 {object} schema.Response[model.Stories]
// @Failure 400 {object} schema.Response[schema.Empty]
// @Security BearerAuth
// @tags stories
// @Router /api/v1/stories/create [post]
func (h *Handler) CreateStory(c *gin.Context) error {
	ctx := c.Request.Context()

	var data admin.CreateStories
	err := c.Bind(&data)
	if err != nil {
		return err
	}

	stories, err := h.services.Stories.Create(ctx, data)
	if err != nil {
		return err
	}

	return schema.Respond(stories, c)
}

// UpdateStory
// WhoAmi godoc
// @Summary обновить сторис
// @Accept json
// @Produce json
// @Param data body admin.UpdateStories true "CreateProduct data"
// @Success 200 {object} schema.Response[model.Stories]
// @Failure 400 {object} schema.Response[schema.Empty]
// @Security BearerAuth
// @tags stories
// @Router /api/v1/stories [put]
func (h *Handler) UpdateStory(c *gin.Context) error {
	ctx := c.Request.Context()

	var data admin.UpdateStories
	err := c.Bind(&data)
	if err != nil {
		return err
	}

	stories, err := h.services.Stories.Update(ctx, data)
	if err != nil {
		return err
	}

	return schema.Respond(stories, c)
}

// GetAllStories
// WhoAmi godoc
// @Summary получить список сторисов
// @Accept json
// @Produce json
// @Success 200 {object} schema.Response[[]model.Stories]
// @Failure 400 {object} schema.Response[schema.Empty]
// @tags stories
// @Router /api/v1/stories/all [get]
func (h *Handler) GetAllStories(c *gin.Context) error {
	ctx := c.Request.Context()

	response, err := h.services.Stories.GetAll(ctx)
	if err != nil {
		return err
	}
	return schema.Respond(response, c)
}

// GetStoriesBuId
// WhoAmi godoc
// @Summary получить список сторисов
// @Accept json
// @Produce json
// @Param stories_id query int true "id"
// @Success 200 {object} schema.Response[model.Stories]
// @Failure 400 {object} schema.Response[schema.Empty]
// @tags stories
// @Router /api/v1/stories/id [get]
func (h *Handler) GetStoriesBuId(c *gin.Context) error {
	ctx := c.Request.Context()

	storiesId, err := strconv.ParseUint(c.Query("stories_id"), 10, 64)
	if err != nil {
		return err
	}
	response, err := h.services.Stories.GetByID(ctx, uint(storiesId))
	if err != nil {
		return err
	}
	return schema.Respond(response, c)
}

// DeleteStoriesById
// WhoAmi godoc
// @Summary удалить сторис
// @Accept json
// @Produce json
// @Param stories_id query int true "id"
// @Success 200 {object} schema.Response[schema.Empty]
// @Failure 400 {object} schema.Response[schema.Empty]
// @tags stories
// @Router /api/v1/stories/id [delete]
func (h *Handler) DeleteStoriesById(c *gin.Context) error {
	ctx := c.Request.Context()

	storiesId, err := strconv.ParseUint(c.Query("stories_id"), 10, 64)
	if err != nil {
		return err
	}
	err = h.services.Stories.DeleteByID(ctx, uint(storiesId))
	if err != nil {
		return err
	}
	return schema.Respond(schema.Empty{}, c)
}

// CreateStoryPage
// WhoAmi godoc
// @Summary добавить страницу в сторис
// @Accept json
// @Produce json
// @Param data body admin.CreateStories true "CreateProduct data"
// @Success 200 {object} schema.Response[model.Product]
// @Failure 400 {object} schema.Response[schema.Empty]
// @Security BearerAuth
// @tags stories
// @Router /api/v1/story-page [post]
func (h *Handler) CreateStoryPage(c *gin.Context) error {
	ctx := c.Request.Context()

	var data admin.CreateStoryPage
	err := c.Bind(&data)
	if err != nil {
		return err
	}

	storyPage, err := h.services.Stories.CreatePage(ctx, data)
	if err != nil {
		return err
	}

	return schema.Respond(storyPage, c)
}

// UpdateStoryPage
// WhoAmi godoc
// @Summary обновить страницу в сторис
// @Accept json
// @Produce json
// @Param data body admin.UpdateStories true "CreateProduct data"
// @Success 200 {object} schema.Response[model.Product]
// @Failure 400 {object} schema.Response[schema.Empty]
// @Security BearerAuth
// @tags stories
// @Router /api/v1/story-page [put]
func (h *Handler) UpdateStoryPage(c *gin.Context) error {
	ctx := c.Request.Context()

	var data admin.UpdateStoryPage
	err := c.Bind(&data)
	if err != nil {
		return err
	}

	stories, err := h.services.Stories.UpdateStoryPage(ctx, data)
	if err != nil {
		return err
	}

	return schema.Respond(stories, c)
}

// DeleteStoryPageById
// WhoAmi godoc
// @Summary удалить страницу в сторис
// @Accept json
// @Produce json
// @Param story_page_id query int true "id"
// @Success 200 {object} schema.Response[schema.Empty]
// @Failure 400 {object} schema.Response[schema.Empty]
// @tags stories
// @Router /api/v1/story-page [delete]
func (h *Handler) DeleteStoryPageById(c *gin.Context) error {
	ctx := c.Request.Context()

	storiesId, err := strconv.ParseUint(c.Query("story_page_id"), 10, 64)
	if err != nil {
		return err
	}
	err = h.services.Stories.DeletePageByID(ctx, uint(storiesId))
	if err != nil {
		return err
	}
	return schema.Respond(schema.Empty{}, c)
}
