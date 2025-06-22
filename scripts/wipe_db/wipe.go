package main

import (
	"log"

	"github.com/meles-z/entainbalancer/configs"
	dbutils "github.com/meles-z/entainbalancer/internal/db_utils"
	"github.com/meles-z/entainbalancer/internal/entities"
)

func main() {

	// Load configuration
	cfg, err := configs.LoadConfig()
	if err != nil {
		log.Fatal("Failed to load configuration:", err)
	}
	cfg.DB.Host = "localhost"
	
	db, err := dbutils.InitDB(&cfg.DB)
	if err != nil {
		log.Fatal("Failed to initialize database:", err)
	}

	// Drop tables
	tables := []any{
		entities.User{},
		entities.Transaction{},
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
