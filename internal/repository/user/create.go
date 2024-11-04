package user

import (
	"context"
	"log"

	"github.com/milovanovmaksim/auth/internal/client/database"
	repo "github.com/milovanovmaksim/auth/internal/repository"
)

func (r *userRepositoryImpl) CreateUser(ctx context.Context, request repo.CreateUserRequest) (*repo.CreateUserResponse, error) {
	var response repo.CreateUserResponse

	query := database.Query{Name: "Create user", QueryRow: "INSERT INTO users (username, email, password, role) VALUES($1, $2, $3, $4) returning id"}

	err := r.db.DB().ScanOneContext(ctx, response, query, request.Name, request.Email, request.HashPassword, request.Role)
	if err != nil {
		log.Printf("failed to insert user userRepositoryImpl.CreateUser || err: %v", err)
		return nil, err
	}

	return &response, nil
}
