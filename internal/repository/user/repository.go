package user

import (
	"github.com/milovanovmaksim/auth/internal/pgsql"
	"github.com/milovanovmaksim/auth/internal/repository"
)

type userRepositoryImpl struct {
	pgSQL pgsql.PostgreSQL
}

func NewUserRepositoryImpl(pgSQL pgsql.PostgreSQL) repository.UserRepository {
	return &userRepositoryImpl{pgSQL}
}
