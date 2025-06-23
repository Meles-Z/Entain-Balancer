package handlers

import (
	"github.com/meles-z/entainbalancer/internal/service"
)

type Handler struct {
	UserService        service.UserService
	TransactionService service.TransactionService
}

func NewHandler(userService service.UserService, transactionService service.TransactionService) *Handler {
	return &Handler{
		UserService:        userService,
		TransactionService: transactionService,
	}
}
