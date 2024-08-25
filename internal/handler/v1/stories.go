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

func (h *Handler) ReadStoriesByUser(c *gin.Context) error {
	ctx := c.Request.Context()
	token := c.GetHeader("Authorization")
	userId, _ := h.services.Auth.VerifyToken(token)

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

func (h *Handler) GetActiveStories(c *gin.Context) error {
	ctx := c.Request.Context()

	token := c.GetHeader("Authorization")
	userId, err := h.services.Auth.VerifyToken(token)
	if err != nil {
		return err
	}
	response, err := h.services.Stories.GetByUserId(ctx, userId)
	if err != nil {
		return err
	}
	return schema.Respond(response, c)
}
