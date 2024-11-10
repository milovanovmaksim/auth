package user

import (
	"context"
	"log"

	"github.com/milovanovmaksim/auth/internal/service"
)

// GetUser возвращает данные о пльзователе по его id.
func (s *userServiceImpl) GetUser(ctx context.Context, request int64) (*service.GetUserResponse, error) {
	user, err := s.userRepository.GetUser(ctx, request)
	if err != nil {
		log.Printf("failed to get user userServiceImpl.GetUser || err: %v", err)
		return nil, err
	}

	res := user.Into()

	return &res, nil
}
