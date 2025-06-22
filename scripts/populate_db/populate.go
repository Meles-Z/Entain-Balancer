package main

import (
	"log"

	"github.com/meles-z/entainbalancer/configs"
	dbutils "github.com/meles-z/entainbalancer/internal/db_utils"
	"github.com/meles-z/entainbalancer/internal/entities"
	"github.com/shopspring/decimal"
)

func main() {
	cfg, err := configs.LoadConfig()
	if err != nil {
		log.Fatal("Failed to load config:", err)
	}
	cfg.DB.Host= "localhost"

	db, err := dbutils.InitDB(&cfg.DB)
	if err != nil {
		log.Fatal("Failed to initialize database:", err)
	}

	// Create predefined users with IDs 1, 2, and 3 as required
	users := []entities.User{
		{
			ID: 1,
			Balance: decimal.RequireFromString("1000.00"),
		},
		{
			ID: 2,
			Balance: decimal.RequireFromString("2000.00"),
		},
		{
			ID: 3,
			Balance: decimal.RequireFromString("3000.00"),
		},
	}

	for _, u := range users {
		// Use FirstOrCreate to avoid duplicates if the app restarts
		if u.Balance.IsNegative() {
			log.Fatalf("User balance cannot be negative: %s", u.Balance)
		}
		if err := db.FirstOrCreate(&u, entities.User{ID: u.ID}).Error; err != nil {
			log.Fatalf("Failed to create user %d: %v", u.ID, err)
		}
	}
	log.Println("Successfully created/verified users with IDs 1, 2, and 3")
}
