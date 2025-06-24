package main

import (
	"log"

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

	// Drop tables
	tables := []any{
		user.User{},
		transaction.Transaction{},
	}

	for _, table := range tables {
		if err := db.Migrator().DropTable(table); err != nil {
			logger.Error("Error dropping table %T: %v\n", table, err)
		} else {
			logger.Info("Successfully dropped table %T\n", table)
		}
	}

	log.Println("Database wiped successfully.")

}
