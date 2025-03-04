package server

import (
	"context"
	"fmt"
	"net/http"
	"time"
	"work-project/internal/config"
	"work-project/internal/handler"
)

type Server struct {
	server *http.Server
}

func NewServer(cfg *config.Config, handler *handler.Handler) (*Server, error) {
	httpHandler, err := handler.Init(cfg)
	if err != nil {
		return nil, err
	}

	return &Server{
		server: &http.Server{
			Addr:           fmt.Sprintf(":%s", cfg.Service.Port),
			Handler:        httpHandler,
			ReadTimeout:    10 * time.Second,
			WriteTimeout:   10 * time.Second,
			IdleTimeout:    30 * time.Second, // Закрывает неиспользуемые соединения
			MaxHeaderBytes: 1 << 20,          // 1MB
		},
	}, nil
}

func (s *Server) Run() error {
	return s.server.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}
