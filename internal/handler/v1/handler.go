package v1

import (
	"github.com/gin-gonic/gin"
	"work-project/internal/middleware"
	"work-project/internal/service"
	"work-project/internal/usecase"
)

type Handler struct {
	services *service.Services
	auth     *middleware.AuthMiddleware
	usecases *usecase.Usecases
}

func NewHandler(services *service.Services, usecases *usecase.Usecases, auth *middleware.AuthMiddleware) *Handler {
	return &Handler{
		services: services,
		auth:     auth,
		usecases: usecases,
	}
}

func (h *Handler) Init(api *gin.RouterGroup) {
	v1 := api.Group("/v1")
	{
		h.initUser(v1)
		h.initReferral(v1)
		h.initBalance(v1)
	}
}
