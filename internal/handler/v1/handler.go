package v1

import (
	"github.com/gin-gonic/gin"
	"work-project/internal/middleware"
	"work-project/internal/service"
)

type Handler struct {
	services *service.Services
	auth     *middleware.AuthMiddleware
}

func NewHandler(services *service.Services, auth *middleware.AuthMiddleware) *Handler {
	return &Handler{
		services: services,
		auth:     auth,
	}
}

func (h *Handler) Init(api *gin.RouterGroup) {
	v1 := api.Group("/v1")
	{
		h.initUser(v1)
	}
}
