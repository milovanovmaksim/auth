package user

import (
	"github.com/milovanovmaksim/auth/internal/client/database"
	"github.com/milovanovmaksim/auth/internal/repository"
	"github.com/milovanovmaksim/auth/internal/service"
	"golang.org/x/crypto/bcrypt"
)

type userServiceImpl struct {
	userRepository repository.UserRepository
	txManager      database.TxManager
}

func (s *userServiceImpl) hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// NewUserService создает новый объект, удовлетворяющий интерфейсу service.UserService.
func NewUserService(userRepository repository.UserRepository, txManager database.TxManager) service.UserService {
	return &userServiceImpl{userRepository, txManager}
}
