package main

import (
	"fmt"
	"log"
	"reflect"

	"github.com/meles-z/entainbalancer/internal/config"
	"github.com/meles-z/entainbalancer/internal/domain/transaction"
	"github.com/meles-z/entainbalancer/internal/domain/user"
	"github.com/meles-z/entainbalancer/internal/infrastructure/db"
	"github.com/meles-z/entainbalancer/internal/infrastructure/logger"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Failed to load configuration:", err)
	}
	cfg.DB.Host = "localhost"

	if err := logger.Init(cfg.Auth.AppEnv); err != nil {
		log.Fatalf("Failed to initialize logger: %v", err)
	}
	defer logger.Sync()

	db, err := db.InitDB(&cfg.DB)
	if err != nil {
		logger.Fatal("Failed to initialize database:", "error", err)
	}

	// List of tables to clear
	models := []any{
		&user.User{},
		&transaction.Transaction{},
	}

	for _, model := range models {
		// Get the type name (e.g., "User", "Transaction")
		modelName := reflect.TypeOf(model).Elem().Name()

		if err := db.Exec("TRUNCATE TABLE users, transactions RESTART IDENTITY CASCADE").Error; err != nil {
			logger.Error(fmt.Sprintf("Failed to clear table %s: %v", modelName, err))
		} else {
			logger.Info(fmt.Sprintf("Successfully cleared table %s", modelName))
		}
	}

	log.Println("All specified tables cleared successfully.")
}
