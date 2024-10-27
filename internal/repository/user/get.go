package user

import (
	"context"
	"log"

	repo "github.com/milovanovmaksim/auth/internal/repository"
)

func (s *userRepositoryImpl) GetUser(ctx context.Context, request int64) (*repo.GetUserResponse, error) {
	var user repo.GetUserResponse
	pool := s.pgSQL.GetPool()

	row := pool.QueryRow(ctx, "SELECT id, username, email, role, created_at, updated_at FROM users WHERE id = $1", request)

	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Role, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		log.Printf("failed to get user userRepositoryImpl.GetUser || err:  %v", err)
		return nil, err
	}

	return &user, nil
}
