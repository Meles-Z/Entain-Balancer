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
	GetTransactionByID(id string) (*entities.Transaction, error)
	UpdateTransaction(transaction *entities.Transaction) error
}

type transactionService struct {
	repo repository.TransactionRepository
}

func NewTransactionService(repo repository.TransactionRepository) TransactionService {
	return &transactionService{repo: repo}
}

func (s *transactionService) CreateTransaction(transaction *entities.Transaction) (*entities.Transaction, error) {
	if err := s.repo.CreateTransaction(transaction); err != nil {
		return nil, err
	}
	return transaction, nil
}

func (s *transactionService) GetTransactionByID(id string) (*entities.Transaction, error) {
	transaction, err := s.repo.GetTransactionByID(id)
	if err != nil {
		return nil, err
	}
	return transaction, nil
}

func (s *transactionService) UpdateTransaction(tx *entities.Transaction) error {
	// Check if transaction already exists (idempotency)
	existing, _ := s.repo.GetTransactionByID(tx.TransactionID)
	if existing != nil {
		return ErrTransactionAlreadyProcessed
	}

	// Get user and balance
	user, err := s.repo.GetUserByID(tx.UserID)
	if err != nil {
		return err
	}

	// Convert string amount to decimal
	amount, err := decimal.NewFromString(tx.Amount)
	if err != nil {
		return err
	}

	newBalance := user.Balance
	if tx.State == entities.TransactionStateWin {
		newBalance = newBalance.Add(amount)
	} else if tx.State == entities.TransactionStateLose {
		if user.Balance.LessThan(amount) {
			return ErrInsufficientBalance
		}
		newBalance = newBalance.Sub(amount)
	}

	// Save transaction and update balance sequentially
	if err := s.repo.CreateTransaction(tx); err != nil {
		return err
	}
	if err := s.repo.UpdateUserBalance(tx.UserID, newBalance); err != nil {
		return err
	}

	return nil
}
