package uow

import (
	"github.com/meles-z/entainbalancer/internal/domain/transaction"
	"github.com/meles-z/entainbalancer/internal/domain/user"
	"gorm.io/gorm"
)

type UnitOfWork interface {
	WithTransaction(fn func(txRepo transaction.TransactionRepository, userRepo user.UserRepository) error) error
}

type unitOfWork struct {
	db *gorm.DB
}

func NewUnitOfWork(db *gorm.DB) UnitOfWork {
	return &unitOfWork{db: db}
}

func (u *unitOfWork) WithTransaction(fn func(txRepo transaction.TransactionRepository, userRepo user.UserRepository) error) error {
	return u.db.Transaction(func(tx *gorm.DB) error {
		txRepo := transaction.NewTransactionRepository(tx)
		userRepo := user.NewUserRepository(tx)
		return fn(txRepo, userRepo)
	})
}
