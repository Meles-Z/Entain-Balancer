package repository

import (
	"gorm.io/gorm"
)

type UnitOfWork interface {
	WithTransaction(fn func(txRepo TransactionRepository, userRepo UserRepository) error) error
}

type unitOfWork struct {
	db *gorm.DB
}

func NewUnitOfWork(db *gorm.DB) UnitOfWork {
	return &unitOfWork{db: db}
}

func (u *unitOfWork) WithTransaction(fn func(txRepo TransactionRepository, userRepo UserRepository) error) error {
	return u.db.Transaction(func(tx *gorm.DB) error {
		txRepo := NewTransactionRepository(tx)
		userRepo := NewUserRepository(tx)
		return fn(txRepo, userRepo)
	})
}
