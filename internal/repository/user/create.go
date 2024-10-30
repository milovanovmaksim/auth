package user

import (
	"context"
	"log"

	repo "github.com/milovanovmaksim/auth/internal/repository"
)

func (r *userRepositoryImpl) CreateUser(ctx context.Context, request repo.CreateUserRequest) (*repo.CreateUserResponse, error) {
	var response repo.CreateUserResponse

	pool := r.pgSQL.GetPool()

	err := pool.QueryRow(ctx, "INSERT INTO users (username, email, password, role) VALUES($1, $2, $3, $4) returning id",
		request.Name, request.Email, request.HashPassword, request.Role).Scan(&response.ID)
	if err != nil {
		log.Printf("failed to insert user userRepositoryImpl.CreateUser || err: %v", err)
		return nil, err
	}

	return &response, nil
}
