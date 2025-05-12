package services

import (
	"goravel/app/models"
	"goravel/app/repositories/interfaces"
	serviceinterfaces "goravel/app/services/interfaces"
)

type UserService struct {
	userRepository interfaces.UserRepositoryInterface
}

// Ensure UserService implements UserServiceInterface
var _ serviceinterfaces.UserServiceInterface = (*UserService)(nil)

func NewUserService(userRepo interfaces.UserRepositoryInterface) *UserService {
	return &UserService{
		userRepository: userRepo,
	}
}

func (s *UserService) GetAllUsers() ([]models.User, error) {
	return s.userRepository.All()
}

func (s *UserService) GetUser(id int) (models.User, error) {
	return s.userRepository.Find(id)
}

func (s *UserService) CreateUser(user models.User) (models.User, error) {
	return s.userRepository.Create(user)
}

func (s *UserService) UpdateUser(id int, user models.User) (models.User, error) {
	return s.userRepository.Update(id, user)
}

func (s *UserService) DeleteUser(id int) error {
	return s.userRepository.Delete(id)
}
