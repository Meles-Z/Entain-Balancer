package main

import (
	"log"

	"github.com/meles-z/entainbalancer/internal/config"
	"github.com/meles-z/entainbalancer/internal/domain/transaction"
	"github.com/meles-z/entainbalancer/internal/domain/user"
	"github.com/meles-z/entainbalancer/internal/infrastucture/db"
	"github.com/shopspring/decimal"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Failed to load config:", err)
	}
	cfg.DB.Host = "localhost"

	db, err := db.InitDB(&cfg.DB)
	if err != nil {
		log.Fatal("Failed to initialize database:", err)
	}

	// Create predefined users with IDs 1, 2, and 3 as required
	users := []user.User{
		{
			ID:      1,
			Balance: decimal.RequireFromString("1000.00"),
		},
		{
			ID:      2,
			Balance: decimal.RequireFromString("2000.00"),
		},
		{
			ID:      3,
			Balance: decimal.RequireFromString("3000.00"),
		},
	}

	for _, u := range users {
		// Use FirstOrCreate to avoid duplicates if the app restarts
		if u.Balance.IsNegative() {
			log.Fatalf("transaction balance cannot be negative: %s", u.Balance)
		}
		if err := db.FirstOrCreate(&u, user.User{ID: u.ID}).Error; err != nil {
			log.Fatalf("Failed to create user %d: %v", u.ID, err)
		}
	}

	transactions := []transaction.Transaction{
		{
			TransactionID: "tx1",
			UserID:        1,
			State:         transaction.TransactionStateWin,
			Amount:        "100.00",
			SourceType:    transaction.SourceTypeGame,
		},
		{
			TransactionID: "tx2",
			UserID:        2,
			State:         transaction.TransactionStateLose,
			Amount:        "50.00",
			SourceType:    transaction.SourceTypeServer,
		},
		{
			TransactionID: "tx3",
			UserID:        3,
			State:         transaction.TransactionStateWin,
			Amount:        "200.00",
			SourceType:    transaction.SourceTypePayment,
		},
	}
	for _, tx := range transactions {
		// Use FirstOrCreate to avoid duplicates if the app restarts
		if err := db.FirstOrCreate(&tx, transaction.Transaction{
			TransactionID: tx.TransactionID,
			UserID:        tx.UserID,
			State:         tx.State,
			Amount:        tx.Amount,
			SourceType:    tx.SourceType,
		}).Error; err != nil {
			log.Fatalf("Failed to create transaction %s: %v", tx.TransactionID, err)
		}
	}

	log.Println("Successfully created/verified users with IDs 1, 2, and 3")
}
