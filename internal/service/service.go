package service

import (
	"context"
)

// UserService интерфейс, отвечающий за бизнес логику приложения.
type UserService interface {
	CreateUser(ctx context.Context, request CreateUserRequest) (*CreateUserResponse, error)
	GetUser(ctx context.Context, request int64) (*GetUserResponse, error)
	UpdateUser(ctx context.Context, request UpdateUserRequest) error
	DeleteUser(ctx context.Context, request int64) error
}
