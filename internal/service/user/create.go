package user

import (
	"context"
	"fmt"
	"log"

	repo "github.com/milovanovmaksim/auth/internal/repository"
	"github.com/milovanovmaksim/auth/internal/service"
)

func (s *userServiceImpl) CreateUser(ctx context.Context, request service.CreateUserRequest) (*service.CreateUserResponse, error) {
	if request.Password != request.PasswordConfirm {
		return nil, fmt.Errorf("password and password_confirm must be the same")
	}

	hashPassword, err := s.hashPassword(request.Password)
	if err != nil {
		log.Printf("failed to get hash for password || err: %v", err)
		return nil, fmt.Errorf("internal error")
	}

	user, err := s.userRepository.CreateUser(ctx, repo.CreateUserRequest{
		Name:         request.Name,
		Email:        request.Email,
		HashPassword: hashPassword,
		Role:         request.Role.String(),
	})
	if err != nil {
		log.Printf("failed to cretae user userServiceImpl.CreateUser || err: %v", err)
		return nil, err
	}

	res := user.Into()

	return &res, nil
}
