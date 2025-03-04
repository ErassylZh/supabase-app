package v1

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"work-project/internal/admin"
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

// CreateCollection
// WhoAmi godoc
// @Accept json
// @Produce json
// @Param data body model.Collection true "CreatePost data"
// @Success 200 {object} schema.Response[model.Collection]
// @Failure 400 {object} schema.Response[schema.Empty]
// @Security BearerAuth
// @tags publication
// @Router /api/v1/collection [post]
func (h *Handler) CreateCollection(c *gin.Context) error {
	ctx := c.Request.Context()

	var data admin.CreateCollection
	err := c.Bind(&data)
	if err != nil {
		return err
	}

	collection, err := h.services.Collection.Create(ctx, data)
	if err != nil {
		return err
	}

	return schema.Respond(collection, c)
}

// GetCollectionByID
// WhoAmi godoc
// @Accept json
// @Produce json
// @Param data body model.Collection true "CreatePost data"
// @Param language query string true "language"
// @Param search query string false "search"
// @Param post_ids query string false "post_ids"
// @Param post_type query string true "all, post, partner"
// @Success 200 {object} schema.Response[[]model.Post]
// @Failure 400 {object} schema.Response[schema.Empty]
// @tags publication
// @Router /api/v1/collection/ [get]
func (h *Handler) GetCollectionByID(c *gin.Context) error {
	ctx := c.Request.Context()

	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		return err
	}

	collection, err := h.services.Collection.GetByID(ctx, uint(id))
	if err != nil {
		return err
	}

	return schema.Respond(collection, c)
}

// UpdateCollection
// WhoAmi godoc
// @Accept json
// @Produce json
// @Param hashtag_id query string false "hashtag_id"
// @Param collection_id query string false "collection_id"
// @Success 200 {object} schema.Response[[]model.Post]
// @Failure 400 {object} schema.Response[schema.Empty]
// @tags publication
// @Router /api/v1/collection/ [get]
func (h *Handler) UpdateCollection(c *gin.Context) error {
	ctx := c.Request.Context()

	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		return err
	}

	var data admin.UpdateCollection
	err = c.Bind(&data)
	if err != nil {
		return err
	}
	data.CollectionID = uint(id)

	collection, err := h.services.Collection.Update(ctx, data)
	if err != nil {
		return err
	}

	return schema.Respond(collection, c)
}

// DeleteCollection
// WhoAmi godoc
// @Accept json
// @Produce json
// @Param collection_id query int false "collection_id"
// @Success 200 {object} schema.Response[schema.Empty]
// @Failure 400 {object} schema.Response[schema.Empty]
// @tags publication
// @Router /api/v1/collection/all [get]
func (h *Handler) DeleteCollection(c *gin.Context) error {
	ctx := c.Request.Context()

	id, err := strconv.Atoi(c.Param("collection_id"))
	if err != nil {
		return err
	}

	err = h.services.Collection.Delete(ctx, uint(id))
	if err != nil {
		return err
	}

	return schema.Respond(schema.Empty{}, c)
}
