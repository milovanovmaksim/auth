package user

import (
	"context"
	"log"

	"github.com/milovanovmaksim/auth/internal/service/user/model"
	repoModel "github.com/milovanovmaksim/auth/internal/repository/user/model"
)

// UpdateUser обновляет данные о пользователе.
func (s *userServiceImpl) UpdateUser(ctx context.Context, request model.UpdateUserRequest) error {

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
