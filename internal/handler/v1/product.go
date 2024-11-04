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
	v1.POST(
		"/product/buy",
		middleware.GinErrorHandle(h.BuyProduct),
	)
	v1.GET(
		"/orders",
		middleware.GinErrorHandle(h.GetMyOrders),
	)
}

// GetListingProducts
// WhoAmi godoc
// @Summary список товаров
// @Accept json
// @Produce json
// @Success 200 {object} schema.Response[[]model.Product]
// @Failure 400 {object} schema.Response[schema.Empty]
// @tags product
// @Router /api/v1/product [get]
func (h *Handler) GetListingProducts(c *gin.Context) error {
	ctx := c.Request.Context()
	products, err := h.services.Product.GetListing(ctx)
	if err != nil {
		return err
	}
	return schema.Respond(products, c)
}

// BuyProduct
// WhoAmi godoc
// @Summary покупка товара
// @Accept json
// @Produce json
// @Success 200 {object} schema.Response[schema.Empty]
// @Failure 400 {object} schema.Response[schema.Empty]
// @Security BearerAuth
// @Param data body schema.ProductBuyRequest true "post"
// @tags product
// @Router /api/v1/product/buy [post]
func (h *Handler) BuyProduct(c *gin.Context) error {
	ctx := c.Request.Context()
	token := c.GetHeader("Authorization")
	userId, err := h.services.Auth.VerifyToken(token)
	if err != nil {
		return err
	}
	var data schema.ProductBuyRequest
	if err := c.Bind(&data); err != nil {
		return err
	}
	data.UserId = userId

	err = h.usecases.Product.Buy(ctx, data)
	if err != nil {
		return err
	}

	return schema.Respond(schema.Empty{}, c)
}

// GetMyOrders
// WhoAmi godoc
// @Summary список моих заказов
// @Accept json
// @Produce json
// @Success 200 {object} schema.Response[[]model.Order]
// @Failure 400 {object} schema.Response[schema.Empty]
// @tags product
// @Router /api/v1/orders [get]
func (h *Handler) GetMyOrders(c *gin.Context) error {
	ctx := c.Request.Context()
	token := c.GetHeader("Authorization")
	userId, err := h.services.Auth.VerifyToken(token)
	if err != nil {
		return err
	}

	order, err := h.services.Order.GetOrders(ctx, userId)
	if err != nil {
		return err
	}

	return schema.Respond(order, c)
}
