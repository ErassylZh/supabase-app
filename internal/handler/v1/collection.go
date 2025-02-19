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
// @Param language query string true "language"
// @tags collection
// @Router /api/v1/collection [get]
func (h *Handler) GetAllCollections(c *gin.Context) error {
	ctx := c.Request.Context()
	token := c.GetHeader("Authorization")
	var userIdp *string
	if len(token) > 0 {
		userId, err := h.services.Auth.VerifyToken(token)
		if err != nil {
			return err
		}
		userIdp = &userId
	}

	language := c.Query("language")
	collections, err := h.services.Collection.GetAllCollection(ctx, language, userIdp, true)
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
// @Param language query string true "language"
// @tags collection
// @Router /api/v1/recommendation [get]
func (h *Handler) GetAllRecommendations(c *gin.Context) error {
	ctx := c.Request.Context()
	language := c.Query("language")
	recommendations, err := h.services.Collection.GetAllRecommendation(ctx, language)
	if err != nil {
		return err
	}
	return schema.Respond(recommendations, c)
}
