package user

import (
	serviceModel "github.com/milovanovmaksim/auth/internal/service/user/model"
	appError "github.com/milovanovmaksim/auth/internal/error"
)

func checkPassword(request serviceModel.CreateUserRequest) error {
	if request.Password == "" {
		return appError.ValidationError{String: "password is empty"}
	}

	if len(request.Password) <= 8 {
		return appError.ValidationError{String: "password must be more than 8 characters"}
	}

	if request.Password != request.PasswordConfirm {
		return appError.ValidationError{String: "password and password_confirm must be the same"}
	}

	return nil
}

func checkName(request serviceModel.CreateUserRequest) error {
	if request.Name == "" {
		return appError.ValidationError{String: "field 'name' is empty"}
	}
	return nil
}

func checkEmail(request serviceModel.CreateUserRequest) error {
	if request.Email == "" {
		return appError.ValidationError{String: "field 'Email' is empty"}
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
