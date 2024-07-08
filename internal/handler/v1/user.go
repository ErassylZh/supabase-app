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

func (h *Handler) DeleteUserById(c *gin.Context) error {
	ctx := c.Request.Context()
	userId := c.Param("user_id")
	err := h.services.User.DeleteByID(ctx, userId)
	if err != nil {
		return err
	}
	return schema.Respond(schema.Empty{}, c)
}

func (h *Handler) GetUserByID(c *gin.Context) error {
	ctx := c.Request.Context()
	userId := c.Param("user_id")
	user, err := h.services.User.GetById(ctx, userId)
	if err != nil {
		return err
	}
	return schema.Respond(user, c)
}
