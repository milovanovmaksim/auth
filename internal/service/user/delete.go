package user

import (
	"context"
	"fmt"
)

func (s *userServiceImpl) DeleteUser(ctx context.Context, request int64) error {
	err := s.userRepository.DeleteUser(ctx, request)
	if err != nil {
		fmt.Printf("failed to delete user userServiceImpl || err: %v", err)
		return err
	}

	return nil
}
