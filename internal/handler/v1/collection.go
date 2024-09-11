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
	v1.GET(
		"/recommendation",
		middleware.GinErrorHandle(h.GetAllRecommendations),
	)
}

// GetAllCollections
// WhoAmi godoc
// @Summary получить все коллекций
// @Accept json
// @Produce json
// @Success 200 {object} schema.Response[[]model.Collection]
// @Failure 400 {object} schema.Response[schema.Empty]
// @tags collection
// @Router /api/v1/collection [get]
func (h *Handler) GetAllCollections(c *gin.Context) error {
	ctx := c.Request.Context()
	collections, err := h.services.Collection.GetAllCollection(ctx)
	if err != nil {
		return err
	}
	return schema.Respond(collections, c)
}

// GetAllRecommendations
// WhoAmi godoc
// @Summary получить все рекомендаций
// @Accept json
// @Produce json
// @Success 200 {object} schema.Response[[]model.Collection]
// @Failure 400 {object} schema.Response[schema.Empty]
// @tags collection
// @Router /api/v1/recommendation [get]
func (h *Handler) GetAllRecommendations(c *gin.Context) error {
	ctx := c.Request.Context()
	recommendations, err := h.services.Collection.GetAllRecommendation(ctx)
	if err != nil {
		return err
	}
	return schema.Respond(recommendations, c)
}
