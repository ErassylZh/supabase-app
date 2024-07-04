package repository

import (
	"work-project/internal/config"
)

type Repositories struct {
}

func NewRepositories(cfg *config.Config) (*Repositories, error) {

	return &Repositories{}, nil
}
