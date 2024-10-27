package user

import (
	"context"
	"log"
)

func (s *userRepositoryImpl) DeleteUser(ctx context.Context, request int64) error {
	pool := s.pgSQL.GetPool()

	_, err := pool.Exec(ctx, "DELETE FROM USERS WHERE id = $1", request)
	if err != nil {
		log.Printf("failed to delete user userRepositoryImpl.DeleteUser || err: %v", err)
		return err
	}

	return nil
}
