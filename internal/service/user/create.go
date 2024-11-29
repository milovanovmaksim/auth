package user

import (
	"context"
	"fmt"
	"log"

	repoModel "github.com/milovanovmaksim/auth/internal/repository/user/model"
	"github.com/milovanovmaksim/auth/internal/service"
	"github.com/milovanovmaksim/auth/internal/service/user/model"
)

// CreateUser создает новго пользователя.
func (s *userServiceImpl) CreateUser(ctx context.Context, request model.CreateUserRequest) (int64, error) {
	err := checkPassword(request)
	if err != nil {
		log.Printf("failed to create new user || error: %v", err)
		return 0, err
	}

	hashPassword, err := s.hashPassword(request.Password)
	if err != nil {
		log.Printf("failed to get hash for password || err: %v", err)
		return 0, fmt.Errorf("internal error")
	}

	userID, err := s.userRepository.CreateUser(ctx, repoModel.CreateUserRequest{
		Name:         request.Name,
		Email:        request.Email,
		HashPassword: hashPassword,
		Role:         request.Role.String(),
	})
	if err != nil {
		log.Printf("failed to cretae user userServiceImpl.CreateUser || err: %v", err)
		return 0, err
	}

	return userID, nil
}

func checkPassword(request model.CreateUserRequest) error {
	if request.Password == "" {
		return service.NewEmptyPasswordError()
	}

	if len(request.Password) <= 8 {
		return service.NewLengthPasswordError(8)
	}

	if request.Password != request.PasswordConfirm {
		return service.NewConfirmationPasswordError()
	}

	return nil
}
