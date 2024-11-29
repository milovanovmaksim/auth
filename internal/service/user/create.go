package user

import (
	"context"
	"fmt"
	"log"

	repoModel "github.com/milovanovmaksim/auth/internal/repository/user/model"
	"github.com/milovanovmaksim/auth/internal/service"
	serviceModel "github.com/milovanovmaksim/auth/internal/service/user/model"
)

// CreateUser создает новго пользователя.
func (s *userServiceImpl) CreateUser(ctx context.Context, request serviceModel.CreateUserRequest) (int64, error) {
	err := validateInputData(request)
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

func checkPassword(request serviceModel.CreateUserRequest) error {
	if request.Password == "" {
		return service.ValidationError{String: "password is empty"}
	}

	if len(request.Password) <= 8 {
		return service.ValidationError{String: "password must be more than 8 characters"}
	}

	if request.Password != request.PasswordConfirm {
		return service.ValidationError{String: "password and password_confirm must be the same"}
	}

	return nil
}

func checkName(request serviceModel.CreateUserRequest) error {
	if request.Name == "" {
		return service.ValidationError{String: "field 'name' is empty"}
	}
	return nil
}

func checkEmail(request serviceModel.CreateUserRequest) error {
	if request.Email == "" {
		return service.ValidationError{String: "field 'Email' is empty"}
	}

	return nil
}

func validateInputData(request serviceModel.CreateUserRequest) error {
	err := checkName(request)
	if err != nil {
		return err
	}

	err = checkEmail(request)
	if err != nil {
		return err
	}

	err = checkPassword(request)
	if err != nil {
		return err
	}

	return nil
}
