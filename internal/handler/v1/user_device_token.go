package v1

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"work-project/internal/middleware"
	"work-project/internal/schema"
)

func (h *Handler) initUserDeviceToken(v1 *gin.RouterGroup) {
	v1.GET(
		"/user-device-token",
		middleware.GinErrorHandle(h.GetDeviceTokensOfUser),
	)
	v1.POST(
		"/user-device-token",
		middleware.GinErrorHandle(h.AddUserDeviceToken),
	)
	v1.DELETE(
		"/user-device-token",
		middleware.GinErrorHandle(h.DeleteDeviceToken),
	)
}

func (h *Handler) AddUserDeviceToken(c *gin.Context) error {
	ctx := c.Request.Context()
	token := c.GetHeader("Authorization")
	userId, err := h.services.Auth.VerifyToken(token)
	if err != nil {
		return err
	}
	var data schema.UserDeviceTokenCreateRequest
	if err = c.Bind(&data); err != nil {
		return err
	}
	data.UserId = userId

	result, err := h.services.UserDeviceToken.Create(ctx, data)
	if err != nil {
		return err
	}
	return schema.Respond(result, c)
}

func (h *Handler) GetDeviceTokensOfUser(c *gin.Context) error {
	ctx := c.Request.Context()
	token := c.GetHeader("Authorization")
	userId, err := h.services.Auth.VerifyToken(token)
	if err != nil {
		return err
	}
	result, err := h.services.UserDeviceToken.GetByUserId(ctx, userId)
	if err != nil {
		return err
	}
	return schema.Respond(result, c)
}

func (h *Handler) DeleteDeviceToken(c *gin.Context) error {
	ctx := c.Request.Context()
	token := c.GetHeader("Authorization")
	_, err := h.services.Auth.VerifyToken(token)
	if err != nil {
		return err
	}
	deviceTokenIdStr := c.Param("deviceTokenId")
	deviceTokenId, err := strconv.Atoi(deviceTokenIdStr)
	if err != nil {
		return err
	}

	err = h.services.UserDeviceToken.DeleteById(ctx, uint(deviceTokenId))
	if err != nil {
		return err
	}
	return schema.Respond(schema.Empty{}, c)
}
