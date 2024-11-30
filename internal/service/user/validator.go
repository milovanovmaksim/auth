package user

import (
	"github.com/milovanovmaksim/auth/internal/service"
	serviceModel "github.com/milovanovmaksim/auth/internal/service/user/model"
)

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

// ValidateInputData валидирует входные данные.
func ValidateInputData(request serviceModel.CreateUserRequest) error {
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
