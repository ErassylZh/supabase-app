package v1

import (
	"github.com/gin-gonic/gin"
	"work-project/internal/middleware"
	"work-project/internal/schema"
)

func (h *Handler) initNotification(v1 *gin.RouterGroup) {
	v1.GET(
		"/notification",
		middleware.GinErrorHandle(h.GetNotification),
	)
}

// GetNotification
// WhoAmi godoc
// @Summary получить уведомления
// @Accept json
// @Produce json
// @Success 200 {object} schema.Response[[]model.PushNotification]
// @Failure 400 {object} schema.Response[schema.Empty]
// @Security BearerAuth
// @tags notification
// @Router /api/v1/notification [get]
func (h *Handler) GetNotification(c *gin.Context) error {
	ctx := c.Request.Context()
	tokenDevice := c.Param("token")
	token := c.GetHeader("Authorization")
	_, err := h.services.Auth.VerifyToken(token)
	if err != nil {
		return err
	}
	//todo сделать получение по юзер айди
	notifications, err := h.services.PushNotification.GetAllByUser(ctx, tokenDevice)
	if err != nil {
		return err
	}
	return schema.Respond(notifications, c)
}
