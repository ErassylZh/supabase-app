package app

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"io/ioutil"
	"log"
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
	"work-project/internal/schema"
	"work-project/internal/server"
	"work-project/internal/service"
	"work-project/internal/usecase"
	"work-project/internal/worker"
	"work-project/pkg/db/postgresql"
)

const logFilePath = "app.txt"

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

	service.NewHub()

	// Удаление старого лог-файла
	if err := os.Remove(logFilePath); err != nil && !os.IsNotExist(err) {
		log.Fatalf("Ошибка при удалении старого лог-файла: %v", err)
	}

	// Создание лог-файла
	logFile, err := os.Create(logFilePath)
	if err != nil {
		log.Fatalf("Ошибка при создании лог-файла: %v", err)
	}
	defer logFile.Close()

	// Устанавливаем запись логов в файл и консоль
	multiWriter := io.MultiWriter(logFile, os.Stdout)
	log.SetOutput(multiWriter)
	gin.DefaultWriter = multiWriter
	gin.DefaultErrorWriter = logFile

	srv, err := server.NewServer(cfg, handlerDelivery)
	if err != nil {
		panic(err.Error())
	}

	go func() {
		if err := srv.Run(); !errors.Is(err, http.ErrServerClosed) {
			log.Println("🔥 Server stopped due error:", err.Error())
		} else {
			log.Println("✅ Server shutdown successfully")
		}
	}()
	handlerWorker := worker.NewHandlerWorker(cfg, services, repositories)
	go handlerWorker.Init()

	log.Println(fmt.Sprintf("🚀 Starting server at http://0.0.0.0:%s", cfg.Service.Port))

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit

	// Отправка логов по почте
	if !cfg.IsLocal() {
		log.Println("Приложение остановлено, собираем логи...")
		log.Println("Отправка логов на email...")
		sendEmailLogFile(services.EmailSender, cfg.Email.Username, logFilePath)
		log.Println("Логи отправлены по почте")

		// Завершение работы сервера
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()
		if err := srv.Shutdown(ctx); err != nil {
			log.Fatalf("Ошибка при завершении сервера: %v", err)
		}
		log.Println("Сервер успешно завершил работу")
	}
}

func sendEmailLogFile(emailSender service.EmailSender, username, filePath string) {
	logFile, err := os.Open(filePath)
	if err != nil {
		log.Printf("Ошибка при открытии файла: %v", err)
		return
	}
	defer logFile.Close()

	fileData, err := ioutil.ReadAll(logFile)
	if err != nil {
		log.Printf("Ошибка при чтении файла: %v", err)
		return
	}

	// Преобразуем содержимое файла в прикрепляемое вложение
	err = emailSender.Send(context.Background(), schema.Message{
		Subject:  "log от " + time.Now().String(),
		FileData: fileData,
		To:       []string{"erasyl.zholdas@list.ru"},
		From:     username,
		FileName: "log " + time.Now().Format("2006-02-01") + ".txt",
	})

	if err != nil {
		log.Printf("Ошибка при отправке письма: %v", err)
		return
	}

	log.Println("Лог файл успешно отправлен.")
}
