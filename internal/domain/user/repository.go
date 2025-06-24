package user

import (
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user *User) (*User, error)
	GetUserByID(id uint64) (*User, error)
	UpdateUser(user *User) error
	WithTrx(fn func(userRepo UserRepository) error) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) CreateUser(user *User) (*User, error) {
	if err := r.db.Create(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userRepository) GetUserByID(id uint64) (*User, error) {
	var user User
	if err := r.db.First(&user, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) UpdateUser(user *User) error {
	if err := r.db.Model(&User{}).Where("id = ?", user.ID).Update("balance", user.Balance).Error; err != nil {
		return err
	}
	return nil
}

func (r *userRepository) WithTrx(fn func(userRepo UserRepository) error) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		userTrx := NewUserRepository(tx)
		return fn(userTrx)
	})
}
