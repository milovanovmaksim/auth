package user

import (
	"context"
	"log"

	"github.com/milovanovmaksim/auth/internal/client/database"
	"github.com/milovanovmaksim/auth/internal/repository/user/model"
)

// GetUser возвращает данные пользователя по его id.
func (s *userRepositoryImpl) GetUser(ctx context.Context, request int64) (*model.GetUserResponse, error) {
	var user model.GetUserResponse

	queryRow := "SELECT id, username, email, role, created_at, updated_at FROM users WHERE id = $1"

	query := database.Query{Name: "Get user", QueryRaw: queryRow}

	err := s.db.DB().ScanOneContext(ctx, &user, query, request)
	if err != nil {
		log.Printf("failed to get user from db: %v", err)
		return nil, err
	}

	return &user, nil
}
