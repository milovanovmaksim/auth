package user

import (
	"context"
	"errors"
	"fmt"
	"log"

	repo "github.com/milovanovmaksim/auth/internal/repository"
	"github.com/milovanovmaksim/auth/internal/service"
)

// CreateUser создает новго пользователя.
func (s *userServiceImpl) CreateUser(ctx context.Context, request service.CreateUserRequest) (*service.CreateUserResponse, error) {
	err := checkPassword(request)
	if err != nil {
		log.Printf("failed to create new user || error: %v", err)
		return nil, err
	}

	hashPassword, err := s.hashPassword(request.Password)
	if err != nil {
		log.Printf("failed to get hash for password || err: %v", err)
		return nil, fmt.Errorf("internal error")
	}

	user, err := s.userRepository.CreateUser(ctx, repo.NewCreateUserRequest(request.Name,
		request.Email,
		hashPassword,
		request.Role.String(),
	))
	if err != nil {
		log.Printf("failed to cretae user userServiceImpl.CreateUser || err: %v", err)
		return nil, err
	}

	res := user.Into()

	return &res, nil
}

func checkPassword(request service.CreateUserRequest) error {
	if request.Password == "" {
		return errors.New("password is empty")
	}

	if request.Password != request.PasswordConfirm {
		return errors.New("password and password_confirm must be the same")
	}

	if len(request.Password) <= 8 {
		return errors.New("password must be more then 8 characters")
	}

	return nil
}
