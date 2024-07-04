package config

import (
	"fmt"
	"sync"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type (
	Config struct {
		Service     *Service
		Database    *Database
		Security    *Security
		Integration *Integration
	}

	Service struct {
		Port string `envconfig:"PORT" default:"8000"`
	}

	Database struct {
		ApiKey      string `envconfig:"API_KEY"`
		SupabaseUrl string `envconfig:"SUPABASE_URL"`
	}
	Security struct {
	}
	Integration struct {
	}
)

var (
	once   sync.Once
	config *Config
)

func GetConfig(envfiles ...string) (*Config, error) {
	var err error
	once.Do(func() {
		_ = godotenv.Load(envfiles...)

		var c Config
		err = envconfig.Process("", &c)
		if err != nil {

			err = fmt.Errorf("error parse config from env variables: %w\n", err)
			return
		}

		config = &c
	})

	return config, err
}
