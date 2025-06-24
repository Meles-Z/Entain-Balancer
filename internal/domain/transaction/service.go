package transaction

import (
	"errors"

	"github.com/meles-z/entainbalancer/internal/domain/user"
	"github.com/shopspring/decimal"
)

var (
	ErrTransactionAlreadyProcessed = errors.New("transaction already processed")
	ErrInsufficientBalance         = errors.New("insufficient balance")
)

type TransactionService interface {
	CreateTransaction(transaction *Transaction) (*Transaction, error)
	UpdateTransaction(transaction *Transaction) error
}

type UnitOfWork interface {
	WithTransaction(func(txRepo TransactionRepository, userRepo user.UserRepository) error) error
}

type transactionService struct {
	uow UnitOfWork
}

func NewTransactionService(uow UnitOfWork) TransactionService {
	return &transactionService{uow: uow}
}

func (s *transactionService) CreateTransaction(transaction *Transaction) (*Transaction, error) {
	var result *Transaction
	err := s.uow.WithTransaction(func(txRepo TransactionRepository, _ user.UserRepository) error {
		var err error
		result, err = txRepo.CreateTransaction(transaction)
		return err
	})
	return result, err
}

func (s *transactionService) UpdateTransaction(tx *Transaction) error {
	return s.uow.WithTransaction(func(txRepo TransactionRepository, userRepo user.UserRepository) error {
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
		case TransactionStateWin:
			newBalance = newBalance.Add(amount)
		case TransactionStateLose:
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
