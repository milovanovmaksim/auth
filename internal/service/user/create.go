package user

import (
	"context"
	"fmt"
	"log"

	repoModel "github.com/milovanovmaksim/auth/internal/repository/user/model"
	serviceModel "github.com/milovanovmaksim/auth/internal/service/user/model"
)

// CreateUser создает новго пользователя.
func (s *userServiceImpl) CreateUser(ctx context.Context, request serviceModel.CreateUserRequest) (int64, error) {
	err := ValidateInputData(request)
	if err != nil {
		log.Printf("failed to validate input data: %v", err)
		return 0, err
	}

	hashPassword, err := s.hashPassword(request.Password)
	if err != nil {
		log.Printf("failed to get hash for password: %v", err)
		return 0, fmt.Errorf("internal error")
	}

	userID, err := s.userRepository.CreateUser(ctx, repoModel.CreateUserRequest{
		Name:         request.Name,
		Email:        request.Email,
		HashPassword: hashPassword,
		Role:         request.Role.String(),
	})
	if err != nil {
		log.Printf("failed to cretae user userServiceImpl.CreateUser: %v", err)
		return 0, err
	}

	return userID, nil
}
