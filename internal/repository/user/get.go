package user

import (
	"context"
	"log"

	"github.com/milovanovmaksim/auth/internal/client/database"
	repo "github.com/milovanovmaksim/auth/internal/repository"
)

// GetUser возвращает данные пользователя по его id.
func (s *userRepositoryImpl) GetUser(ctx context.Context, request int64) (*repo.GetUserResponse, error) {
	var user repo.GetUserResponse

	query := database.Query{Name: "Get user", QueryRaw: "SELECT id, username, email, role, created_at, updated_at FROM users WHERE id = $1"}

	err := s.db.DB().ScanOneContext(ctx, &user, query, request)
	if err != nil {
		log.Printf("failed to get user from db || err: %v", err)
		return nil, err
	}

	return &user, nil
}
