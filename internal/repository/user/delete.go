package user

import (
	"context"
	"log"

	"github.com/milovanovmaksim/auth/internal/client/database"
)

// DeleteUser удаляет пользователя из БД.:w
func (s *userRepositoryImpl) DeleteUser(ctx context.Context, request int64) error {
	query := database.Query{Name: "Delete user", QueryRaw: "DELETE FROM USERS WHERE id = $1"}

	_, err := s.db.DB().ExecContext(ctx, query, request)
	if err != nil {
		log.Printf("failed to delete user userRepositoryImpl.DeleteUser || err: %v", err)
		return err
	}

	return nil
}
