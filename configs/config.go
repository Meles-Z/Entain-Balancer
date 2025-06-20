package configs

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

// Config holds all configuration for the application
type Config struct {
	DB DatabaseConfig
}

// DatabaseConfig contains the database connection parameters
type DatabaseConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
}

// LoadConfig initializes and returns the application configuration
func LoadConfig() (*Config, error) {
	if err := initViper(); err != nil {
		return nil, err
	}

	cfg := &Config{
		DB: DatabaseConfig{
			Host:     viper.GetString("DB_HOST"),
			Port:     viper.GetInt("DB_PORT"),
			User:     viper.GetString("DB_USER"),
			Password: viper.GetString("DB_PASSWORD"),
			DBName:   viper.GetString("DB_NAME"),
		},
	}

	return cfg, nil
}

// initViper sets up the viper configuration reader
func initViper() error {
	viper.SetConfigFile(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Println("No .env file found. Falling back to system environment variables.")
			return nil
		}
		return fmt.Errorf("failed to read config file: %w", err)
	}
	return nil
}
