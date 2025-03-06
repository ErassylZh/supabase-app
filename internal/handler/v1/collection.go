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
	v1.POST(
		"/collection",
		middleware.GinErrorHandle(h.CreateCollection),
	)
	v1.GET(
		"/collection/id",
		middleware.GinErrorHandle(h.GetCollectionByID),
	)
	v1.PUT(
		"/collection",
		middleware.GinErrorHandle(h.UpdateCollection),
	)
	v1.DELETE(
		"/collection",
		middleware.GinErrorHandle(h.DeleteCollection),
	)
	v1.POST(
		"/collection/add",
		middleware.GinErrorHandle(h.AddCollectionToPost),
	)
	v1.DELETE(
		"/collection/delete-post",
		middleware.GinErrorHandle(h.DeleteCollectionToPost),
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
// @Param data body admin.CreateCollection true "CreatePost data"
// @Success 200 {object} schema.Response[model.Collection]
// @Failure 400 {object} schema.Response[schema.Empty]
// @Security BearerAuth
// @tags collection
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
// @Param collection_id query int true "collection_id"
// @Success 200 {object} schema.Response[model.Collection]
// @Failure 400 {object} schema.Response[schema.Empty]
// @tags collection
// @Router /api/v1/collection/id [get]
func (h *Handler) GetCollectionByID(c *gin.Context) error {
	ctx := c.Request.Context()

	id, err := strconv.Atoi(c.Query("collection_id"))
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
// @Param data body admin.UpdateCollection true "CreatePost data"
// @Success 200 {object} schema.Response[model.Collection]
// @Failure 400 {object} schema.Response[schema.Empty]
// @tags collection
// @Router /api/v1/collection [put]
func (h *Handler) UpdateCollection(c *gin.Context) error {
	ctx := c.Request.Context()

	var data admin.UpdateCollection
	err := c.Bind(&data)
	if err != nil {
		return err
	}

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
// @tags collection
// @Router /api/v1/collection [delete]
func (h *Handler) DeleteCollection(c *gin.Context) error {
	ctx := c.Request.Context()

	id, err := strconv.Atoi(c.Query("collection_id"))
	if err != nil {
		return err
	}

	err = h.services.Collection.Delete(ctx, uint(id))
	if err != nil {
		return err
	}

	return schema.Respond(schema.Empty{}, c)
}

// AddCollectionToPost
// WhoAmi godoc
// @Summary добавить пост в коллекцию
// @Accept json
// @Produce json
// @Success 200 {object} schema.Response[model.Collection]
// @Failure 400 {object} schema.Response[schema.Empty]
// @Param data body admin.AddCollection true "CreatePost data"
// @tags collection
// @Router /api/v1/collection/add [post]
func (h *Handler) AddCollectionToPost(c *gin.Context) error {
	ctx := c.Request.Context()

	var data admin.AddCollection
	if err := c.BindJSON(&data); err != nil {
		return err
	}

	collections, err := h.services.Collection.AddToPost(ctx, data)
	if err != nil {
		return err
	}
	return schema.Respond(collections, c)
}

// DeleteCollectionToPost
// WhoAmi godoc
// @Summary удалить пост из коллекцию
// @Accept json
// @Produce json
// @Success 200 {object} schema.Response[model.Collection]
// @Failure 400 {object} schema.Response[schema.Empty]
// @Param data body admin.DeleteCollectionPost true "CreatePost data"
// @tags collection
// @Router /api/v1/collection/delete-post [delete]
func (h *Handler) DeleteCollectionToPost(c *gin.Context) error {
	ctx := c.Request.Context()

	var data admin.DeleteCollectionPost
	if err := c.BindJSON(&data); err != nil {
		return err
	}

	collections, err := h.services.Collection.DeleteCollectionPost(ctx, data)
	if err != nil {
		return err
	}
	return schema.Respond(collections, c)
}
