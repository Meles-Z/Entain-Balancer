package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/meles-z/entainbalancer/configs"
	dbutils "github.com/meles-z/entainbalancer/internal/db_utils"
	"github.com/meles-z/entainbalancer/internal/handlers"
	"github.com/meles-z/entainbalancer/internal/repository"
	"github.com/meles-z/entainbalancer/internal/service"
)

func Server() {
	cfg, err := configs.LoadConfig()
	if err != nil {
		fmt.Printf("Error to load configuration: %v\n", err)
		return
	}

	db, err := dbutils.InitDB(&cfg.DB)
	if err != nil {
		fmt.Printf("Error to connect to database: %v\n", err)
		return
	}
	fmt.Println("Database connection established successfully")

	if err := dbutils.RunMigrations(db); err != nil {
		fmt.Printf("Error running migrations: %v\n", err)
		return
	}
	fmt.Println("Migrations completed successfully")

	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)

	uow := repository.NewUnitOfWork(db)
	txService := service.NewTransactionService(uow)

	handler := handlers.NewHandler(userService, txService)

	router := Route(handler)

	log.Println("🚀 Server listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
