package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/meles-z/entainbalancer/configs"
	dbutils "github.com/meles-z/entainbalancer/internal/db_utils"
)

func Server() {
	cfg, err := configs.LoadConfig()
	if err != nil {
		fmt.Printf("Error to load configuration:%v", err)
		return
	}

	db, err := dbutils.InitDB(&cfg.DB)
	if err != nil {
		fmt.Printf("Error to connect to database:%v", err)
		return
	}
	fmt.Println("Database connection established successfully")
	if err := dbutils.RunMigrations(db); err != nil {
		fmt.Printf("Error running migrations: %v\n", err)
		return
	}

	fmt.Println("Migrations completed successfully")

	log.Println("🚀 Server started at :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("❌ Server failed: %v", err)
	}
}
