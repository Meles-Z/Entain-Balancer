package repository

import (
	"github.com/meles-z/entainbalancer/internal/entities"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type TransactionRepository interface {
	CreateTransaction(transaction *entities.Transaction) error
	GetTransactionByID(id string) (*entities.Transaction, error)
	GetUserByID(userID uint64) (*entities.User, error)
	UpdateUserBalance(userID uint64, newBalance decimal.Decimal) error
}

type transactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) TransactionRepository {
	return &transactionRepository{db: db}
}

func (r *transactionRepository) CreateTransaction(transaction *entities.Transaction) error {
	return r.db.Create(transaction).Error
}

func (r *transactionRepository) GetTransactionByID(id string) (*entities.Transaction, error) {
	var tx entities.Transaction
	if err := r.db.First(&tx, "transaction_id = ?", id).Error; err != nil {
		return nil, err
	}
	return &tx, nil
}

func (r *transactionRepository) GetUserByID(userID uint64) (*entities.User, error) {
	var user entities.User
	if err := r.db.First(&user, "user_id = ?", userID).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *transactionRepository) UpdateUserBalance(userID uint64, newBalance decimal.Decimal) error {
	return r.db.Model(&entities.User{}).
		Where("user_id = ?", userID).
		Update("balance", newBalance).Error
}
