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
		middleware.GinErrorHandle(h.CreateProduct),
	)
	group.PUT(
		"/product",
		middleware.GinErrorHandle(h.UpdateProduct),
	)
	group.GET(
		"/product/id",
		middleware.GinErrorHandle(h.GetProduct),
	)
	group.DELETE(
		"/product",
		middleware.GinErrorHandle(h.DeleteProduct),
	)
	group.POST(
		"/product-tag",
		middleware.GinErrorHandle(h.CreateProductTag),
	)
	group.POST(
		"/product-tag/add",
		middleware.GinErrorHandle(h.AddProductTagToProduct),
	)
	group.DELETE(
		"/product-tag",
		middleware.GinErrorHandle(h.DeleteProductTag),
	)
	group.PUT(
		"/product-tag",
		middleware.GinErrorHandle(h.UpdateProductTag),
	)
	group.DELETE(
		"/product-tag/product",
		middleware.GinErrorHandle(h.DeleteProductTagToProduct),
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
// @Router /api/v1/admin/post [put]
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
// @Router /api/v1/admin/post [delete]
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
// @Router /api/v1/admin/post [get]
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
// @Router /api/v1/admin/product [put]
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
// @Router /api/v1/admin/product [delete]
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
// @Router /api/v1/admin/product/id [get]
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

// CreateProductTag
// WhoAmi godoc
// @Accept json
// @Produce json
// @Param data body admin.CreateProductTag true "CreatePost data"
// @Success 200 {object} schema.Response[model.ProductTag]
// @Failure 400 {object} schema.Response[schema.Empty]
// @Security BearerAuth
// @tags post
// @Router /api/v1/admin/product-tag [post]
func (h *Handler) CreateProductTag(c *gin.Context) error {
	ctx := c.Request.Context()

	var data admin.CreateProductTag
	err := c.Bind(&data)
	if err != nil {
		return err
	}

	productTag, err := h.services.ProductTag.Create(ctx, data)
	if err != nil {
		return err
	}

	return schema.Respond(productTag, c)
}

// UpdateProductTag
// WhoAmi godoc
// @Accept json
// @Produce json
// @Param data body admin.UpdatePost true "UserLogin data"
// @Success 200 {object} schema.Response[model.Post]
// @Failure 400 {object} schema.Response[schema.Empty]
// @Security BearerAuth
// @tags post
// @Router /api/v1/admin/product-tag [put]
func (h *Handler) UpdateProductTag(c *gin.Context) error {
	ctx := c.Request.Context()

	var data admin.UpdateProductTag
	err := c.Bind(&data)
	if err != nil {
		return err
	}

	productTag, err := h.services.ProductTag.Update(ctx, data)
	if err != nil {
		return err
	}

	return schema.Respond(productTag, c)
}

// DeleteProductTag
// WhoAmi godoc
// @Accept json
// @Produce json
// @Param post_id query int true "id"
// @Success 200 {object} schema.Response[schema.Empty]
// @Failure 400 {object} schema.Response[schema.Empty]
// @Security BearerAuth
// @tags post
// @Router /api/v1/admin/product-tag [delete]
func (h *Handler) DeleteProductTag(c *gin.Context) error {
	ctx := c.Request.Context()

	productTagId, err := strconv.ParseUint(c.Query("product_tag_id"), 10, 64)
	if err != nil {
		return err
	}

	err = h.services.ProductTag.Delete(ctx, uint(productTagId))
	if err != nil {
		return err
	}

	return schema.Respond(schema.Empty{}, c)
}

// AddProductTagToProduct
// WhoAmi godoc
// @Summary добавить пост в коллекцию
// @Accept json
// @Produce json
// @Success 200 {object} schema.Response[model.ProductTag]
// @Failure 400 {object} schema.Response[schema.Empty]
// @Param data body admin.AddProductProductTag true "CreatePost data"
// @tags collection
// @Router /api/v1/admin/product-tag/add [post]
func (h *Handler) AddProductTagToProduct(c *gin.Context) error {
	ctx := c.Request.Context()

	var data admin.AddProductProductTag
	if err := c.BindJSON(&data); err != nil {
		return err
	}

	productTag, err := h.services.ProductTag.AddToProduct(ctx, data)
	if err != nil {
		return err
	}
	return schema.Respond(productTag, c)
}

// DeleteProductTagToProduct
// WhoAmi godoc
// @Summary удалить пост из коллекцию
// @Accept json
// @Produce json
// @Success 200 {object} schema.Response[model.ProductTag]
// @Failure 400 {object} schema.Response[schema.Empty]
// @Param data body admin.DeleteProductProductTag true "CreatePost data"
// @tags collection
// @Router /api/v1/collection/delete [delete]
func (h *Handler) DeleteProductTagToProduct(c *gin.Context) error {
	ctx := c.Request.Context()

	var data admin.DeleteProductProductTag
	if err := c.BindJSON(&data); err != nil {
		return err
	}

	productTag, err := h.services.ProductTag.DeleteToProduct(ctx, data)
	if err != nil {
		return err
	}
	return schema.Respond(productTag, c)
}
