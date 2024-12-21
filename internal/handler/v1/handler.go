package v1

import (
	"context"
	"github.com/gin-gonic/gin"
	"work-project/internal/aggregator"
	"work-project/internal/middleware"
	"work-project/internal/model"
	"work-project/internal/service"
	"work-project/internal/usecase"
)

type Handler struct {
	services          *service.Services
	serviceAggregator *aggregator.ServiceAggregatorService
	auth              *middleware.AuthMiddleware
	usecases          *usecase.Usecases
}

func NewHandler(services *service.Services, serviceAggregator *aggregator.ServiceAggregatorService, usecases *usecase.Usecases, auth *middleware.AuthMiddleware) *Handler {
	return &Handler{
		services:          services,
		serviceAggregator: serviceAggregator,
		auth:              auth,
		usecases:          usecases,
	}
}

func (h *Handler) Init(api *gin.RouterGroup) {
	v1 := api.Group("/v1")
	{
		h.initUser(v1)
		h.initReferral(v1)
		h.initBalance(v1)
		h.initNotification(v1)
		h.initUserDeviceToken(v1)
		h.initProduct(v1)
		h.initPost(v1)
		h.initStories(v1)
		h.initMark(v1)
		h.initHashtag(v1)
		h.initCollection(v1)
	}
}

func (h *Handler) InitWs(api *gin.RouterGroup) {
	v1 := api.Group("/v1")
	{
		h.initRatingSocket(v1)
	}
}

func (h *Handler) GetUserFromToken(token string) (model.User, error) {
	return h.services.Auth.GetUserByToken(context.Background(), token)
}
