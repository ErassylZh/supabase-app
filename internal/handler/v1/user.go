package v1

import (
	"github.com/gin-gonic/gin"
	"work-project/internal/middleware"
	"work-project/internal/schema"
)

func (h *Handler) initUser(v1 *gin.RouterGroup) {
	v1.DELETE(
		"/user/:user_id",
		middleware.GinErrorHandle(h.DeleteUserById),
	)
	v1.GET(
		"/user/:user_id",
		middleware.GinErrorHandle(h.GetUserByID),
	)
}

// DeleteUserById
// WhoAmi godoc
// @Summary удалить пользователя
// @Accept json
// @Produce json
// @Success 200 {object} schema.Response[schema.Empty]
// @Failure 400 {object} schema.Response[schema.Empty]
// @Security BearerAuth
// @tags user
// @Router /api/v1/user/:user_id [delete]
func (h *Handler) DeleteUserById(c *gin.Context) error {
	ctx := c.Request.Context()
	userId := c.Param("user_id")
	err := h.services.User.DeleteByID(ctx, userId)
	if err != nil {
		return err
	}
	return schema.Respond(schema.Empty{}, c)
}

// GetUserByID
// WhoAmi godoc
// @Summary получить пользователя
// @Accept json
// @Produce json
// @Success 200 {object} schema.Response[model.User]
// @Failure 400 {object} schema.Response[schema.Empty]
// @Security BearerAuth
// @tags user
// @Router /api/v1/user/:user_id [get]
func (h *Handler) GetUserByID(c *gin.Context) error {
	ctx := c.Request.Context()
	userId := c.Param("user_id")
	user, err := h.services.User.GetById(ctx, userId)
	if err != nil {
		return err
	}
	return schema.Respond(user, c)
}
