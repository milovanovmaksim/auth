package user

import (
	"context"
	"log"

	"github.com/milovanovmaksim/auth/internal/repository"
)

func (s *userRepositoryImpl) UpdateUser(ctx context.Context, request repository.UpdateUserRequest) error {
	pool := s.pgSQL.GetPool()

	_, err := pool.Exec(ctx, "UPDATE users SET username = COALESCE($1, username), role = COALESCE($2, role), WHERE id = $3",
		request.Name, request.Role, request.ID)
	if err != nil {
		log.Printf("failed to update user userRepositoryImpl.UpdateUser || err: %v", err)
		return err
	}

	return nil
}
