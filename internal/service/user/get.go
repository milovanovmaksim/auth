package user

import (
	"context"
	"log"

	"github.com/milovanovmaksim/auth/internal/repository/user/converter"
	"github.com/milovanovmaksim/auth/internal/service/user/model"
)

// GetUser возвращает данные о пльзователе по его id.
func (s *userServiceImpl) GetUser(ctx context.Context, request int64) (*model.GetUserResponse, error) {
	user, err := s.userRepository.GetUser(ctx, request)
	if err != nil {
		log.Printf("failed to get user userServiceImpl.GetUser: %v", err)
		return nil, err
	}

	res := converter.ToServiceFromGetUserResponse(*user)

	return &res, nil
}
