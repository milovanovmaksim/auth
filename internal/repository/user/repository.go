package user

import (
	"github.com/milovanovmaksim/auth/internal/client/database"
	"github.com/milovanovmaksim/auth/internal/repository"
)

type userRepositoryImpl struct {
	db database.Client
}

// NewUserRepository создает новый объект, удовлетваряющий интерфейсу repository.UserRepository.
func NewUserRepository(db database.Client) repository.UserRepository {
	return &userRepositoryImpl{db}
}
