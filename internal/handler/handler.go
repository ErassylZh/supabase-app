package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
	"net/http"
	"runtime"
	"work-project/internal/aggregator"
	"work-project/internal/config"
	v1 "work-project/internal/handler/v1"
	"work-project/internal/middleware"
	"work-project/internal/service"
	"work-project/internal/usecase"
	"work-project/pkg/middlewares"
)

type Handler struct {
	usecases          *usecase.Usecases
	services          *service.Services
	serviceAggregator *aggregator.ServiceAggregatorService
	baseUrl           string
	authMiddleware    middleware.AuthMiddleware
	healthcheckFn     func() error
}

func NewHandlerDelivery(
	usecases *usecase.Usecases,
	services *service.Services,
	serviceAggregator *aggregator.ServiceAggregatorService,
	baseUrl string,
	auth middleware.AuthMiddleware,
	healthcheckFn func() error,
) *Handler {
	return &Handler{
		usecases:          usecases,
		services:          services,
		serviceAggregator: serviceAggregator,
		baseUrl:           baseUrl,
		authMiddleware:    auth,
		healthcheckFn:     healthcheckFn,
	}
}

func (h *Handler) Init(cfg *config.Config) (*gin.Engine, error) {
	app := gin.New()
	app.Use(
		gin.Logger(),
		gin.Recovery(),
		middlewares.Cors(),
		middlewares.Recovery(middleware.GinRecoveryFn),
		middlewares.LoggerMiddleware(),
		func(c *gin.Context) { // Логируем кол-во горутин перед обработкой запроса
			log.Printf("🔄 Перед обработкой запроса: %d горутин", runtime.NumGoroutine())
			c.Next()
			log.Printf("✅ После обработки запроса: %d горутин", runtime.NumGoroutine())
		},
	)

	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	app.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, map[string]string{"message": "pong"})
	})

	app.GET("/goroutine", func(c *gin.Context) {
		c.JSON(http.StatusOK, map[string]interface{}{
			"number": runtime.NumGoroutine(),
			"data": func() string {
				buf := make([]byte, 1<<16)
				n := runtime.Stack(buf, true)
				return fmt.Sprintf("%s", buf[:n])
			}(),
		})
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

	handlerV1 := v1.NewHandler(h.services, h.serviceAggregator, h.usecases, &h.authMiddleware)
	api := baseUrl.Group("/api")
	{
		handlerV1.Init(api)
	}
	ws := baseUrl.Group("/ws")
	{
		handlerV1.InitWs(ws)
	}
}
