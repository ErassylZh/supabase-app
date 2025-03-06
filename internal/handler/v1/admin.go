package v1

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"work-project/internal/admin"
	"work-project/internal/middleware"
	"work-project/internal/schema"
)

func (h *Handler) initAdmin(v1 *gin.RouterGroup) {
	group := v1.Group("/admin")
	group.POST("/post", middleware.GinErrorHandle(h.CreatePost))
	group.GET("/post", middleware.GinErrorHandle(h.GetPost))
	group.PUT("/post", middleware.GinErrorHandle(h.UpdatePost))
	group.DELETE("/post", middleware.GinErrorHandle(h.DeletePost))
	group.POST(
		"/product",
		middleware.GinErrorHandle(h.GetListingProducts),
	)
	group.PUT(
		"/product",
		middleware.GinErrorHandle(h.GetListingProducts),
	)
	group.GET(
		"/product",
		middleware.GinErrorHandle(h.GetListingProducts),
	)
	group.DELETE(
		"/product",
		middleware.GinErrorHandle(h.GetListingProducts),
	)

}

// CreatePost
// WhoAmi godoc
// @Accept json
// @Produce json
// @Param data body admin.CreatePost true "CreatePost data"
// @Success 200 {object} schema.Response[model.Post]
// @Failure 400 {object} schema.Response[schema.Empty]
// @Security BearerAuth
// @tags post
// @Router /api/v1/admin/post [post]
func (h *Handler) CreatePost(c *gin.Context) error {
	ctx := c.Request.Context()

	var data admin.CreatePost
	err := c.Bind(&data)
	if err != nil {
		return err
	}

	post, err := h.services.Post.Create(ctx, data)
	if err != nil {
		return err
	}

	return schema.Respond(post, c)
}

// UpdatePost
// WhoAmi godoc
// @Accept json
// @Produce json
// @Param data body admin.UpdatePost true "UserLogin data"
// @Success 200 {object} schema.Response[model.Post]
// @Failure 400 {object} schema.Response[schema.Empty]
// @Security BearerAuth
// @tags post
// @Router /api/v1/post [put]
func (h *Handler) UpdatePost(c *gin.Context) error {
	ctx := c.Request.Context()

	var data admin.UpdatePost
	err := c.Bind(&data)
	if err != nil {
		return err
	}

	post, err := h.services.Post.Update(ctx, data)
	if err != nil {
		return err
	}

	return schema.Respond(post, c)
}

// DeletePost
// WhoAmi godoc
// @Accept json
// @Produce json
// @Param post_id query int true "id"
// @Success 200 {object} schema.Response[schema.Empty]
// @Failure 400 {object} schema.Response[schema.Empty]
// @Security BearerAuth
// @tags post
// @Router /api/v1/post [delete]
func (h *Handler) DeletePost(c *gin.Context) error {
	ctx := c.Request.Context()

	postId, err := strconv.ParseUint(c.Query("post_id"), 10, 64)
	if err != nil {
		return err
	}

	err = h.services.Post.Delete(ctx, uint(postId))
	if err != nil {
		return err
	}

	return schema.Respond(schema.Empty{}, c)
}

// GetPost
// WhoAmi godoc
// @Accept json
// @Produce json
// @Param post_id query int true "id"
// @Success 200 {object} schema.Response[schema.Empty]
// @Failure 400 {object} schema.Response[schema.Empty]
// @Security BearerAuth
// @tags post
// @Router /api/v1/post [get]
func (h *Handler) GetPost(c *gin.Context) error {
	ctx := c.Request.Context()

	postId, err := strconv.ParseUint(c.Query("post_id"), 10, 64)
	if err != nil {
		return err
	}

	post, err := h.services.Post.GetById(ctx, uint(postId))
	if err != nil {
		return err
	}

	return schema.Respond(post, c)
}

// CreateProduct
// WhoAmi godoc
// @Accept json
// @Produce json
// @Param data body admin.CreateProduct true "CreateProduct data"
// @Success 200 {object} schema.Response[model.Product]
// @Failure 400 {object} schema.Response[schema.Empty]
// @Security BearerAuth
// @tags product
// @Router /api/v1/admin/product [post]
func (h *Handler) CreateProduct(c *gin.Context) error {
	ctx := c.Request.Context()

	var data admin.CreateProduct
	err := c.Bind(&data)
	if err != nil {
		return err
	}

	product, err := h.services.Product.Create(ctx, data)
	if err != nil {
		return err
	}

	return schema.Respond(product, c)
}

// UpdateProduct
// WhoAmi godoc
// @Accept json
// @Produce json
// @Param data body admin.UpdateProduct true "UserLogin data"
// @Success 200 {object} schema.Response[model.Product]
// @Failure 400 {object} schema.Response[schema.Empty]
// @Security BearerAuth
// @tags product
// @Router /api/v1/product [put]
func (h *Handler) UpdateProduct(c *gin.Context) error {
	ctx := c.Request.Context()

	var data admin.UpdateProduct
	err := c.Bind(&data)
	if err != nil {
		return err
	}

	product, err := h.services.Product.Update(ctx, data)
	if err != nil {
		return err
	}

	return schema.Respond(product, c)
}

// DeleteProduct
// WhoAmi godoc
// @Accept json
// @Produce json
// @Param product_id query int true "id"
// @Success 200 {object} schema.Response[schema.Empty]
// @Failure 400 {object} schema.Response[schema.Empty]
// @Security BearerAuth
// @tags product
// @Router /api/v1/product [delete]
func (h *Handler) DeleteProduct(c *gin.Context) error {
	ctx := c.Request.Context()

	productId, err := strconv.ParseUint(c.Query("product_id"), 10, 64)
	if err != nil {
		return err
	}

	err = h.services.Product.Delete(ctx, uint(productId))
	if err != nil {
		return err
	}

	return schema.Respond(schema.Empty{}, c)
}

// GetProduct
// WhoAmi godoc
// @Accept json
// @Produce json
// @Param product_id query int true "id"
// @Success 200 {object} schema.Response[schema.Empty]
// @Failure 400 {object} schema.Response[schema.Empty]
// @Security BearerAuth
// @tags product
// @Router /api/v1/product [get]
func (h *Handler) GetProduct(c *gin.Context) error {
	ctx := c.Request.Context()

	productId, err := strconv.ParseUint(c.Query("product_id"), 10, 64)
	if err != nil {
		return err
	}

	product, err := h.services.Product.GetById(ctx, uint(productId))
	if err != nil {
		return err
	}

	return schema.Respond(product, c)
}
