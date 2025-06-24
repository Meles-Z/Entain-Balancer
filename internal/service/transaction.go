package service

import (
	"errors"

	"github.com/meles-z/entainbalancer/internal/entities"
	"github.com/meles-z/entainbalancer/internal/repository"
	"github.com/shopspring/decimal"
)

var (
	ErrTransactionAlreadyProcessed = errors.New("transaction already processed")
	ErrInsufficientBalance         = errors.New("insufficient balance")
)

type TransactionService interface {
	CreateTransaction(transaction *entities.Transaction) (*entities.Transaction, error)
	UpdateTransaction(transaction *entities.Transaction) error
}

type UnitOfWork interface {
	WithTransaction(func(txRepo repository.TransactionRepository, userRepo repository.UserRepository) error) error
}

type transactionService struct {
	uow UnitOfWork
}

func NewTransactionService(uow UnitOfWork) TransactionService {
	return &transactionService{uow: uow}
}

func (s *transactionService) CreateTransaction(transaction *entities.Transaction) (*entities.Transaction, error) {
	var result *entities.Transaction
	err := s.uow.WithTransaction(func(txRepo repository.TransactionRepository, _ repository.UserRepository) error {
		var err error
		result, err = txRepo.CreateTransaction(transaction)
		return err
	})
	return result, err
}

func (s *transactionService) UpdateTransaction(tx *entities.Transaction) error {
	return s.uow.WithTransaction(func(txRepo repository.TransactionRepository, userRepo repository.UserRepository) error {
		// Idempotency check
		exists, err := txRepo.IsTransactionExists(tx.TransactionID)
		if err != nil {
			return err
		}
		if exists {
			return ErrTransactionAlreadyProcessed
		}

		// Get user
		user, err := userRepo.GetUserByID(tx.UserID)
		if err != nil {
			return err
		}

		// Parse amount
		amount, err := decimal.NewFromString(tx.Amount)
		if err != nil {
			return err
		}

		// Balance calculation
		newBalance := user.Balance
		switch tx.State {
		case entities.TransactionStateWin:
			newBalance = newBalance.Add(amount)
		case entities.TransactionStateLose:
			if user.Balance.LessThan(amount) {
				return ErrInsufficientBalance
			}
			newBalance = newBalance.Sub(amount)
		}

		// Save transaction
		if _, err := txRepo.CreateTransaction(tx); err != nil {
			return err
		}

		// Update user
		user.Balance = newBalance
		return userRepo.UpdateUser(user)
	})
}
