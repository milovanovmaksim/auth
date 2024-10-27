package user

import (
	"context"
	"fmt"
)

func (s *userRepositoryImpl) DeleteUser(ctx context.Context, request int64) error {
	pool := s.pgSQL.GetPool()

	_, err := pool.Exec(ctx, "DELETE FROM USERS WHERE id = $1", request)
	if err != nil {
		fmt.Printf("failed to delete user userRepositoryImpl.DeleteUser || err: %v", err)
		return err
	}

	return nil
}
