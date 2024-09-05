package v1

import (
	"github.com/gin-gonic/gin"
	"work-project/internal/middleware"
	"work-project/internal/schema"
)

func (h *Handler) initHashtag(v1 *gin.RouterGroup) {
	v1.GET(
		"/hashtag",
		middleware.GinErrorHandle(h.GetAllHashtags),
	)
}

func (h *Handler) GetAllHashtags(c *gin.Context) error {
	ctx := c.Request.Context()
	hashtags, err := h.services.Hashtag.GetAll(ctx)
	if err != nil {
		return err
	}
	return schema.Respond(hashtags, c)
}
