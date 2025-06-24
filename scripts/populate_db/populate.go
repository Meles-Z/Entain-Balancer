package main

import (
	"log"

	"github.com/meles-z/entainbalancer/internal/config"
	"github.com/meles-z/entainbalancer/internal/domain/user"
	"github.com/meles-z/entainbalancer/internal/infrastructure/db"
	"github.com/meles-z/entainbalancer/internal/infrastructure/logger"
	"github.com/shopspring/decimal"
)

func main() {

	cfg, err := config.LoadConfig()
	if err != nil {
		logger.Fatal("Failed to load config", "error", err)
	}

	cfg.DB.Host = "localhost"

	if err := logger.Init(cfg.Auth.AppEnv); err != nil {
		log.Fatalf("Failed to initialize logger: %v", err)
	}
	defer logger.Sync()

	db, err := db.InitDB(&cfg.DB)
	if err != nil {
		logger.Fatal("Failed to initialize database", "error", err)
	}

	users := []user.User{
		{ID: 1, Balance: decimal.RequireFromString("1000.00")},
		{ID: 2, Balance: decimal.RequireFromString("2000.00")},
		{ID: 3, Balance: decimal.RequireFromString("3000.00")},
	}

	for _, u := range users {
		if u.Balance.IsNegative() {
			log.Fatalf("User balance cannot be negative: %s", u.Balance)
		}

		if err := db.FirstOrCreate(&u).Error; err != nil {
			logger.Error("Failed to create user", "userID", u.ID, "error", err)
			continue
		}
	}

	logger.Info("Successfully created/verified users", "userIDs", []int{1, 2, 3})
}
