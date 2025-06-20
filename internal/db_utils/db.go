package dbutils

import (
	"fmt"
	"log"
	"time"

	"github.com/meles-z/entainbalancer/configs"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// InitDB initializes and returns a GORM database connection
func InitDB(cfg *configs.DatabaseConfig) (*gorm.DB, error) {
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

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	log.Println("✅ Database connected successfully")
	return db, nil
}
