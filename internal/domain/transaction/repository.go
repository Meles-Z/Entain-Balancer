package transaction

import (
	"gorm.io/gorm"
)

type TransactionRepository interface {
	CreateTransaction(transaction *Transaction) (*Transaction, error)
	IsTransactionExists(id string) (bool, error)
	WithTrx(fn func(txRepo TransactionRepository) error) error
}

type transactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) TransactionRepository {
	return &transactionRepository{db: db}
}

func (r *transactionRepository) CreateTransaction(transaction *Transaction) (*Transaction, error) {
	if err := r.db.Create(&transaction).Error; err != nil {
		return nil, err
	}
	return transaction, nil
}

func (r *transactionRepository) IsTransactionExists(id string) (bool, error) {
	var count int64
	if err := r.db.Model(&Transaction{}).
		Where("transaction_id = ?", id).
		Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}

func (r *transactionRepository) WithTrx(fn func(txRepo TransactionRepository) error) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		txRepo := NewTransactionRepository(tx)
		return fn(txRepo)
	})
}
