package user

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/milovanovmaksim/auth/internal/client/database"
	"github.com/milovanovmaksim/auth/internal/repository/user/model"
)

// UpdateUser обновление данных пользователя.
func (s *userRepositoryImpl) UpdateUser(ctx context.Context, request model.UpdateUserRequest) error {
	username := sql.NullString{Valid: false}

	if request.Name != nil {
		username.String = *request.Name
		username.Valid = true
	}

	queryRow := "UPDATE users SET username = COALESCE($1, username), role = COALESCE($2, role), updated_at = COALESCE($3, updated_at) WHERE id = $4"

	query := database.Query{Name: "Update user", QueryRaw: queryRow}

	_, err := s.db.DB().ExecContext(ctx, query, username, request.Role, time.Now(), request.ID)
	if err != nil {
		log.Printf("failed to update user userRepositoryImpl.UpdateUser: %v", err)
		return err
	}

	return nil
}
