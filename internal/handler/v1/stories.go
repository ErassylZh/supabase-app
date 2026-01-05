package v1

import (
	"github.com/gin-gonic/gin"
	"strconv"
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
