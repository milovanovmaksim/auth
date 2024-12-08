package user

import (
	"context"
	"log"

	repoModel "github.com/milovanovmaksim/auth/internal/repository/user/model"
	serviceModel "github.com/milovanovmaksim/auth/internal/service/user/model"
)

// UpdateUser обновляет данные о пользователе.
func (s *userServiceImpl) UpdateUser(ctx context.Context, request serviceModel.UpdateUserRequest) error {
	err := s.userRepository.UpdateUser(ctx, repoModel.UpdateUserRequest{
		ID:   request.ID,
		Name: request.Name,
		Role: request.Role.String(),
	})
	if err != nil {
		log.Printf("failed to update user userServiceImpl.UpdateUser: %v", err)
		return err
	}

	return nil
}
