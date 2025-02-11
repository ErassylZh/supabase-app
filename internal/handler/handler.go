package handler

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
	"work-project/internal/config"
	v1 "work-project/internal/handler/v1"
	"work-project/internal/middleware"
	"work-project/internal/service"
	"work-project/internal/usecase"
	"work-project/pkg/middlewares"
)

type Handler struct {
	usecases       *usecase.Usecases
	services       *service.Services
	baseUrl        string
	authMiddleware middleware.AuthMiddleware
	healthcheckFn  func() error
}

func NewHandlerDelivery(
	usecases *usecase.Usecases,
	services *service.Services,
	baseUrl string,
	auth middleware.AuthMiddleware,
	healthcheckFn func() error,
) *Handler {
	return &Handler{
		usecases:       usecases,
		services:       services,
		baseUrl:        baseUrl,
		authMiddleware: auth,
		healthcheckFn:  healthcheckFn,
	}
}

func (h *Handler) Init(cfg *config.Config) (*gin.Engine, error) {
	app := gin.Default()
	app.Use(
		middlewares.Cors(),
		middlewares.Recovery(middleware.GinRecoveryFn),
		//h.authMiddleware.SetCurrentUser(),
	)
	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	app.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, map[string]string{"message": "pong"})
	})
	app.GET("/readiness", func(c *gin.Context) {
		if err := h.healthcheckFn(); err != nil {
			c.JSON(http.StatusServiceUnavailable, map[string]string{"message": err.Error()})
			c.Error(err)
		} else {
			c.JSON(http.StatusOK, map[string]string{"message": "ok"})
		}
	})
	app.GET("/liveness", func(c *gin.Context) {
		if err := h.healthcheckFn(); err != nil {
			c.JSON(http.StatusServiceUnavailable, map[string]string{"message": err.Error()})
			c.Error(err)
		} else {
			c.JSON(http.StatusOK, map[string]string{"message": "ok"})
		}
	})

	h.initAPI(app)

	return app, nil
}

func (h *Handler) initAPI(router *gin.Engine) {
	baseUrl := router.Group(h.baseUrl)

	handlerV1 := v1.NewHandler(h.services, h.usecases, &h.authMiddleware)
	api := baseUrl.Group("/api")
	{
		handlerV1.Init(api)
	}
}
