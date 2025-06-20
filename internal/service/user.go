package service

import (
	"github.com/meles-z/entainbalancer/internal/entities"
	"github.com/meles-z/entainbalancer/internal/repository"
)

type UserService interface {
	CreateUser(user *entities.User) (*entities.User, error)
	GetUserByID(id uint64) (*entities.User, error)
	GetUserByEmail(email string) (*entities.User, error)
	UpdateUser(user *entities.User) error
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) CreateUser(user *entities.User) (*entities.User, error) {
	createdUser, err := s.repo.CreateUser(user)
	if err != nil {
		return nil, err
	}
	return createdUser, nil
}

func (s *userService) GetUserByID(id uint64) (*entities.User, error) {
	user, err := s.repo.GetUserByID(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *userService) GetUserByEmail(email string) (*entities.User, error) {
	user, err := s.repo.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *userService) UpdateUser(user *entities.User) error {
	if err := s.repo.UpdateUser(user); err != nil {
		return err
	}
	return nil
}
