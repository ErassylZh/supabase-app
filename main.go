package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	_ "work-project/docs"
	"work-project/internal/app"
	"work-project/internal/config"
)

// @title           swagger API
// @version         1.0
// @description     Это пример сервера с использованием Swagger в Gin.
// @host            157.90.227.125:8000
// @BasePath        /
//
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {

	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()
	cfg, err := config.GetConfig()
	if err != nil {
		panic(err)
	}

	app.Run(cfg)
}
