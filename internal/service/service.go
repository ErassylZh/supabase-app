package service

import (
	"work-project/internal/config"
	"work-project/internal/repository"
)

//go:generate go run github.com/golang/mock/mockgen -source=service.go -destination=service_mock.go -package=service

type Services struct {
}

type Deps struct {
	Repos *repository.Repositories
	Cgf   *config.Config
}

func NewServices(deps Deps) *Services {
	return &Services{}
}
