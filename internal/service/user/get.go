package user

import (
	"context"
	"log"

	"github.com/milovanovmaksim/auth/internal/service"
)

func (s *userServiceImpl) GetUser(ctx context.Context, request int64) (*service.GetUserResponse, error) {
	user, err := s.userRepository.GetUser(ctx, request)
	if err != nil {
		log.Printf("failed to get user userServiceImpl.GetUser || err: %v", err)
		return nil, err
	}

	return &service.GetUserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Role:      user.Role,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil
}
