package db

import (
	"fmt"
	"time"

	"github.com/meles-z/entainbalancer/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// InitDB initializes and returns a GORM database connection
func InitDB(cfg *config.DatabaseConfig) (*gorm.DB, error) {
	if cfg.Host == "" || cfg.Port == 0 || cfg.DBName == "" || cfg.User == "" {
		return nil, fmt.Errorf("incomplete database configuration")
	}

	dsn := fmt.Sprintf(
		"host=%s port=%d dbname=%s user=%s password=%s",
		cfg.Host, cfg.Port, cfg.DBName, cfg.User, cfg.Password,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get DB instance from GORM: %w", err)
	}

	sqlDB.SetMaxIdleConns(20)
	sqlDB.SetMaxOpenConns(50)
	sqlDB.SetConnMaxLifetime(30 * time.Minute)
	return db, nil
}
