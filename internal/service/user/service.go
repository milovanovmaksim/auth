package user

import (
	"github.com/milovanovmaksim/auth/internal/repository"
	"github.com/milovanovmaksim/auth/internal/service"
	"golang.org/x/crypto/bcrypt"
)

type userServiceImpl struct {
	userRepository repository.UserRepository
}

func (s *userServiceImpl) hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func NewUserService(userRepository repository.UserRepository) service.UserService {
	return &userServiceImpl{userRepository}
}
