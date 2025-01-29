package main

import (
	_ "work-project/docs"
	"work-project/internal/app"
	"work-project/internal/config"
)

// @title           swagger API
// @version         1.0
// @description     Это пример сервера с использованием Swagger в Gin.
// @host            localhost:8000
// @BasePath        /
//
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {

	cfg, err := config.GetConfig()
	if err != nil {
		panic(err)
	}

	app.Run(cfg)
}
