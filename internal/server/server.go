package server

import (
	"log"
	"net/http"

	"github.com/meles-z/entainbalancer/internal/config"
	httpHandler "github.com/meles-z/entainbalancer/internal/delivery/http"
	"github.com/meles-z/entainbalancer/internal/domain/transaction"
	uow "github.com/meles-z/entainbalancer/internal/domain/unit_of_work"
	"github.com/meles-z/entainbalancer/internal/domain/user"
	"github.com/meles-z/entainbalancer/internal/infrastructure/db"
	"github.com/meles-z/entainbalancer/internal/infrastructure/logger"
)

func Server() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	if err := logger.Init(cfg.Auth.AppEnv); err != nil {
		log.Fatalf("Failed to initialize logger: %v", err)
	}
	defer logger.Sync()

	log.Println("Local development environment started.")

	dbconn, err := db.InitDB(&cfg.DB)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	log.Println("Database connection established successfully")

	if err := db.RunMigrations(dbconn); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}
	log.Println("Migrations completed successfully")

	userRepo := user.NewUserRepository(dbconn)
	userService := user.NewUserService(userRepo)

	uow := uow.NewUnitOfWork(dbconn)
	txService := transaction.NewTransactionService(uow)

	handler := httpHandler.NewHandler(userService, txService)
	router := Route(handler)

	log.Println("🚀 Server listening on :8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
