package v1

import (
	"github.com/gin-gonic/gin"
	"work-project/internal/middleware"
	"work-project/internal/schema"
)

func (h *Handler) initPrivacyTerms(v1 *gin.RouterGroup) {
	v1.GET(
		"/privacy-terms",
		middleware.GinErrorHandle(h.GetPrivacyTerms),
	)
}

// GetPrivacyTerms
// WhoAmi godoc
// @Summary получить все коллекций
// @Accept json
// @Produce json
// @Success 200 {object} schema.Response[[]model.PrivacyTerms]
// @Failure 400 {object} schema.Response[schema.Empty]
// @tags privacy-terms
// @Router /api/v1/privacy-terms [get]
func (h *Handler) GetPrivacyTerms(c *gin.Context) error {
	ctx := c.Request.Context()

	terms, err := h.services.PrivacyTerms.GetAll(ctx)
	if err != nil {
		return err
	}
	return schema.Respond(terms, c)
}
