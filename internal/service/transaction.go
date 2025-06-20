package service

import (
	"github.com/meles-z/entainbalancer/internal/entities"
	"github.com/meles-z/entainbalancer/internal/repository"
)

type TransactionService interface {
	CreateTransaction(transaction *entities.Transaction) (*entities.Transaction, error)
	GetTransactionByID(id string) (*entities.Transaction, error)
	GetTransactionsByUserID(userID string) ([]*entities.Transaction, error)
	GetAllTransactions() ([]*entities.Transaction, error)
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
func (s *transactionService) GetTransactionsByUserID(userID string) ([]*entities.Transaction, error) {
	transactions, err := s.repo.GetTransactionsByUserID(userID)
	if err != nil {
		return nil, err
	}
	return transactions, nil
}
func (s *transactionService) GetAllTransactions() ([]*entities.Transaction, error) {
	transactions, err := s.repo.GetAllTransactions()
	if err != nil {
		return nil, err
	}
	return transactions, nil
}

func (s *transactionService) UpdateTransaction(transaction *entities.Transaction) error {
	if err := s.repo.UpdateTransaction(transaction); err != nil {
		return err
	}
	return nil
}
