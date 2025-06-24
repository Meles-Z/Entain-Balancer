package config

import (
	"github.com/meles-z/entainbalancer/internal/infrastucture/logger"
	"github.com/spf13/viper"
)

// Config holds all configuration for the application
type Config struct {
	DB   DatabaseConfig
	Auth Auth
}

// DatabaseConfig contains the database connection parameters
type DatabaseConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
}

type Auth struct {
	Appenv string
}

// LoadConfig initializes and returns the application configuration
func LoadConfig() (*Config, error) {
	viper.AddConfigPath(".")
	viper.SetConfigFile(".env")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			logger.Error("Error reading config file:", "error", err)
			return nil, err
		}
		logger.Warn("No .env file found, using system env vars", "error", err)
	}

	cfg := &Config{
		DB: DatabaseConfig{
			Host:     viper.GetString("DB_HOST"),
			Port:     viper.GetInt("DB_PORT"),
			User:     viper.GetString("DB_USER"),
			Password: viper.GetString("DB_PASSWORD"),
			DBName:   viper.GetString("DB_NAME"),
		},
		Auth: Auth{
			Appenv: viper.GetString("APP_ENV"),
		},
	}

	return cfg, nil
}
