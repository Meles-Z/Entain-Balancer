package db

import (
	"fmt"

	"github.com/meles-z/entainbalancer/internal/domain/transaction"
	"github.com/meles-z/entainbalancer/internal/domain/user"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) error {
	err := db.AutoMigrate(
		&user.User{},
		&transaction.Transaction{},
	)
	if err != nil {
		return fmt.Errorf("migration failed: %w", err)
	}

	return nil
}
