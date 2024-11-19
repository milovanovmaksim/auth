package user

import (
	"context"
	"log"
	"time"

	"github.com/milovanovmaksim/auth/internal/client/database"
	"github.com/milovanovmaksim/auth/internal/repository"
)

// UpdateUser обновление данных пользователя.
func (s *userRepositoryImpl) UpdateUser(ctx context.Context, request repository.UpdateUserRequest) error {
	query := database.Query{Name: "Update user", QueryRaw: "UPDATE users SET username = COALESCE($1, username), role = COALESCE($2, role), updated_at = COALESCE($3, updated_at) WHERE id = $4"}

	_, err := s.db.DB().ExecContext(ctx, query, request.Name, request.Role, time.Now(), request.ID)
	if err != nil {
		log.Printf("failed to update user userRepositoryImpl.UpdateUser || err: %v", err)
		return err
	}

	return nil
}
