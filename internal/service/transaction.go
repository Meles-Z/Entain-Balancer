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

type transactionService struct {
	repo     repository.TransactionRepository
	userRepo repository.UserRepository
}

func NewTransactionService(
	repo repository.TransactionRepository,
	userRepo repository.UserRepository,
) TransactionService {
	return &transactionService{
		repo:     repo,
		userRepo: userRepo,
	}
}
func (s *transactionService) CreateTransaction(transaction *entities.Transaction) (*entities.Transaction, error) {
	transaction, err := s.repo.CreateTransaction(transaction)
	if err != nil {
		return nil, err
	}
	return transaction, nil
	// Get user by ID
}

func (s *transactionService) UpdateTransaction(tx *entities.Transaction) error {
	// Check if transaction already exists (idempotency)
	exists, err := s.repo.IsTransactionExists(tx.TransactionID)
	if err != nil {
		return err
	}
	if exists {
		return ErrTransactionAlreadyProcessed
	}

	// Get user
	user, err := s.userRepo.GetUserByID(tx.UserID)
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

	// Save transaction and update user balance
	_, err = s.repo.CreateTransaction(tx)
	if err != nil {
		return err
	}
	user.Balance = newBalance
	if err := s.userRepo.UpdateUser(user); err != nil {
		return err
	}

	return nil
}
