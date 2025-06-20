package repository

import (
	"github.com/meles-z/entainbalancer/internal/entities"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user *entities.User) error
	GetUserByID(id string) (*entities.User, error)
	GetUserByEmail(email string) (*entities.User, error)
	UpdateUser(user *entities.User) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) CreateUser(user *entities.User) error {
	if err := r.db.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func (r *userRepository) GetUserByID(id string) (*entities.User, error) {
	var user entities.User
	if err := r.db.First(&user, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
func (r *userRepository) GetUserByEmail(email string) (*entities.User, error) {
	var user entities.User
	if err := r.db.First(&user, "email = ?", email).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
func (r *userRepository) UpdateUser(user *entities.User) error {
	if err := r.db.Save(user).Error; err != nil {
		return err
	}
	return nil
}
