package main

import (
	"work-project/internal/app"
	"work-project/internal/config"
)

func main() {

	cfg, err := config.GetConfig()
	if err != nil {
		panic(err)
	}

	app.Run(cfg)
}
