package user

import (
	"github.com/milovanovmaksim/auth/internal/client/database"
	"github.com/milovanovmaksim/auth/internal/repository"
)

type userRepositoryImpl struct {
	db database.Client
}

func NewUserRepository(db database.Client) repository.UserRepository {
	return &userRepositoryImpl{db}
}
