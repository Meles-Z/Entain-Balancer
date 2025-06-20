package repository

import (
	"github.com/meles-z/entainbalancer/internal/entities"
	"gorm.io/gorm"
)

type Transaction interface {
	CreateTransaction(transaction *entities.Transaction) error
	GetTransactionByID(id string) (*entities.Transaction, error)
	GetTransactionsByUserID(userID string) ([]*entities.Transaction, error)
	GetAllTransactions() ([]*entities.Transaction, error)
	UpdateTransaction(transaction *entities.Transaction) error
}
type transactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) Transaction {
	return &transactionRepository{db: db}
}

func (r *transactionRepository) CreateTransaction(transaction *entities.Transaction) error {
	if err := r.db.Create(transaction).Error; err != nil {
		return err
	}
	return nil
}
func (r *transactionRepository) GetTransactionByID(id string) (*entities.Transaction, error) {
	var transaction entities.Transaction
	if err := r.db.First(&transaction, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &transaction, nil
}
func (r *transactionRepository) GetTransactionsByUserID(userID string) ([]*entities.Transaction, error) {
	var transactions []*entities.Transaction
	if err := r.db.Where("user_id = ?", userID).Find(&transactions).Error; err != nil {
		return nil, err
	}
	return transactions, nil
}
func (r *transactionRepository) GetAllTransactions() ([]*entities.Transaction, error) {
	var transactions []*entities.Transaction
	if err := r.db.Find(&transactions).Error; err != nil {
		return nil, err
	}
	return transactions, nil
}
func (r *transactionRepository) UpdateTransaction(transaction *entities.Transaction) error {
	if err := r.db.Save(transaction).Error; err != nil {
		return err
	}
	return nil
}
