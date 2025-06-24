package main

import (
	"log"

	"github.com/meles-z/entainbalancer/internal/config"
	"github.com/meles-z/entainbalancer/internal/domain/transaction"
	"github.com/meles-z/entainbalancer/internal/domain/user"
	"github.com/meles-z/entainbalancer/internal/infrastucture/db"
)

func main() {

	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Failed to load configuration:", err)
	}
	cfg.DB.Host = "localhost"

	db, err := db.InitDB(&cfg.DB)
	if err != nil {
		log.Fatal("Failed to initialize database:", err)
	}

	// Drop tables
	tables := []any{
		user.User{},
		transaction.Transaction{},
	}

	for _, table := range tables {
		if err := db.Migrator().DropTable(table); err != nil {
			log.Printf("Error dropping table %T: %v\n", table, err)
		} else {
			log.Printf("Successfully dropped table %T\n", table)
		}
	}

	log.Println("Database wiped successfully.")

}
