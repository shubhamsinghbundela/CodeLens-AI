package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Environment string

const (
	Development   Environment = "development"
	Staging       Environment = "staging"
	PreProduction Environment = "pre-production"
	Production    Environment = "production"
)

func (u Environment) IsValid() bool {
	switch u {
	case Development, Staging, PreProduction, Production:
		return true
	default:
		return false
	}
}

const (
	envFilePath = ".env"
)

type Config struct {
	AppEnv   Environment
	Port     string
	LogLevel string
}

func Load() (*Config, error) {
	err := godotenv.Load(envFilePath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, fmt.Errorf("%s file not found", envFilePath)
		}
		return nil, fmt.Errorf("error loading .env file: %v", err)
	}

	config := &Config{
		Port:     getEnv("PORT"),
		LogLevel: getEnv("LOG_LEVEL"),
	}

	appEnv := Environment(getEnv("APP_ENV"))

	if !appEnv.IsValid() {
		return nil, fmt.Errorf("invalid APP_ENV: %s", appEnv)
	}
	config.AppEnv = appEnv

	return config, nil
}

func getEnv(key string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	panic("required environment variable " + key + " is not set")
}
