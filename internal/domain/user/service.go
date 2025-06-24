package user

type UserService interface {
	CreateUser(user *User) (*User, error)
	GetUserByID(id uint64) (*User, error)
	UpdateUser(user *User) error
}

type userService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) CreateUser(user *User) (*User, error) {
	createdUser, err := s.repo.CreateUser(user)
	if err != nil {
		return nil, err
	}
	return createdUser, nil
}

func (s *userService) GetUserByID(id uint64) (*User, error) {
	user, err := s.repo.GetUserByID(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *userService) UpdateUser(user *User) error {
	if err := s.repo.UpdateUser(user); err != nil {
		return err
	}
	return nil
}
