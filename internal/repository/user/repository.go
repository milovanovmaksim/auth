package user

import (
	"github.com/milovanovmaksim/auth/internal/client/database/postgresql"
	"github.com/milovanovmaksim/auth/internal/repository"
)

type userRepositoryImpl struct {
	pgSQL postgresql.PostgreSQL
}

func NewUserRepository(pgSQL postgresql.PostgreSQL) repository.UserRepository {
	return &userRepositoryImpl{pgSQL}
}
