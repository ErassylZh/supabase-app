package v1

import (
	"github.com/gin-gonic/gin"
	"work-project/internal/middleware"
	"work-project/internal/schema"
)

func (h *Handler) initProduct(v1 *gin.RouterGroup) {
	v1.GET(
		"/product",
		middleware.GinErrorHandle(h.GetListingProducts),
	)
}

func (h *Handler) GetListingProducts(c *gin.Context) error {
	ctx := c.Request.Context()
	products, err := h.services.Product.GetListing(ctx)
	if err != nil {
		return err
	}
	return schema.Respond(products, c)

}
