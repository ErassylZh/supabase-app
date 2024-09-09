package v1

import (
	"github.com/gin-gonic/gin"
	"work-project/internal/middleware"
	"work-project/internal/schema"
)

func (h *Handler) initCollection(v1 *gin.RouterGroup) {
	v1.GET(
		"/collection",
		middleware.GinErrorHandle(h.GetAllCollections),
	)
}

func (h *Handler) GetAllCollections(c *gin.Context) error {
	ctx := c.Request.Context()
	collections, err := h.services.Collection.GetAll(ctx)
	if err != nil {
		return err
	}
	return schema.Respond(collections, c)
}
