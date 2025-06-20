package dbutils

import (
	"fmt"

	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) error {
	err := db.AutoMigrate(
	// &entities.User{},
	// &entities.Transaction{},
	// &entities.Escrow{},
	)
	if err != nil {
		return fmt.Errorf("migration failed: %w", err)
	}
	fmt.Println("✅ AutoMigration completed")
	return nil
}
