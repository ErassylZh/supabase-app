package app

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"work-project/internal/aggregator"
	"work-project/internal/config"
	"work-project/internal/handler"
	"work-project/internal/middleware"
	"work-project/internal/repository"
	"work-project/internal/server"
	"work-project/internal/service"
	"work-project/internal/usecase"
	"work-project/internal/worker"
	"work-project/pkg/db/postgresql"
)

func Run(cfg *config.Config) {

	_, db, err := postgresql.NewDB(cfg.Database.Dsn)
	if err != nil {
		fmt.Println(err.Error())
	}

	//connection, err := db.DB()
	//if err != nil {
	//	panic(err)
	//}
	//if err := migrate.Up(connection, "migrations"); err != nil {
	//	fmt.Println("Failed to apply migrations:", err)
	//	panic(err)
	//}
	repositories, err := repository.NewRepositories(db, cfg)
	if err != nil {
		panic("error initialization Repositories " + err.Error())
	}
	services := service.NewServices(service.Deps{
		Repos: repositories,
		Cgf:   cfg,
	})
	usecases := usecase.NewUsecases(usecase.Deps{
		Services:     services,
		Repositories: repositories,
	})

	serviceAggregator := aggregator.NewServiceAggregatorService(*services)

	healthCheckFn := func() error {
		//if err := connection.Ping(); err != nil {
		//	return fmt.Errorf("database is not responding: %w", err)
		//}

		return nil
	}
	//handlers := v1.NewHandler(services)

	authMiddleware := middleware.NewAuthMiddleware(middleware.GinRecoveryFn)
	handlerDelivery := handler.NewHandlerDelivery(usecases, services, serviceAggregator, "", *authMiddleware, healthCheckFn)
	if err != nil {
		fmt.Println("Failed to create handlers:", err)
		panic(err)
	}

	service.NewHub()

	srv, err := server.NewServer(cfg, handlerDelivery)
	if err != nil {
		panic(err.Error())
	}

	go func() {
		if err := srv.Run(); !errors.Is(err, http.ErrServerClosed) {
			fmt.Println("🔥 Server stopped due error", "err", err.Error())
		} else {
			fmt.Println("✅ Server shutdown successfully")
		}
	}()
	handlerWorker := worker.NewHandlerWorker(cfg, services, repositories)
	go func() {
		handlerWorker.Init()
	}()

	fmt.Println(fmt.Sprintf("🚀 Starting server at http://0.0.0.0:%s", cfg.Service.Port))
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit
	_, shutdownCtxCancel := context.WithTimeout(context.Background(), time.Second*30)
	defer shutdownCtxCancel()

}
