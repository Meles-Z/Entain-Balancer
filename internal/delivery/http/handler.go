package http

import (
	"github.com/meles-z/entainbalancer/internal/domain/transaction"
	"github.com/meles-z/entainbalancer/internal/domain/user"
)

type Handler struct {
	UserService       user.UserService
	TransactionService transaction.TransactionService
}

func NewHandler(userService user.UserService, transactionService transaction.TransactionService) *Handler {
	return &Handler{
		UserService:        userService,
		TransactionService: transactionService,
	}
}
