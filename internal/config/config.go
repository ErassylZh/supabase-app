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
		SupabaseApiKey string `envconfig:"SUPABASE_API_KEY"`
		SupabaseUrl    string `envconfig:"SUPABASE_URL"`
		Dsn            string `envconfig:"DATABASE_DSN"`
	}
	Security struct {
		Secret string `envconfig:"SECRET"`
	}
	Integration struct {
		PathToFirebaseConfig       string `envconfig:"PATH_TO_FIREBASE_CONFIG" default:"haco-firebase-config.json"`
		AirtableBaseurl            string `envconfig:"AIRTABLE_BASE_URL"`
		AirtableApiKey             string `envconfig:"AIRTABLE_API_KEY"`
		PushNotificationReadPeriod int    `envconfig:"PUSH_NOTIFICATION_READ_PERIOD" default:"300"`
		AirtableSyncPeriod         int    `envconfig:"AIRTABLE_SYNC_PERIOD" default:"300"`
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
