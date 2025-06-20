package dbutils

import (
	"fmt"

	"github.com/meles-z/entainbalancer/internal/entities"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) error {
	err := db.AutoMigrate(
		&entities.User{},
		&entities.Transaction{},
	)
	if err != nil {
		return fmt.Errorf("migration failed: %w", err)
	}
	fmt.Println("✅ AutoMigration completed")
	return nil
}
