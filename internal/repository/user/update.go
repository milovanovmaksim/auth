package user

import (
	"context"
	"log"

	"github.com/milovanovmaksim/auth/internal/client/database/postgresql"
	"github.com/milovanovmaksim/auth/internal/repository"
)

func (s *userRepositoryImpl) UpdateUser(ctx context.Context, request repository.UpdateUserRequest) error {
	query := postgresql.Query{Name: "Update user", QueryRow: "UPDATE users SET username = COALESCE($1, username), role = COALESCE($2, role), WHERE id = $3"}

	_, err := s.pgSQL.ExecContext(ctx, query, request.Name, request.Role, request.ID)
	if err != nil {
		log.Printf("failed to update user userRepositoryImpl.UpdateUser || err: %v", err)
		return err
	}

	return nil
}
