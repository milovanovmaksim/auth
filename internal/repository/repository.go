package repository

import (
	"context"
)

type UserRepository interface {
	CreateUser(ctx context.Context, request CreateUserRequest) (*CreateUserResponse, error)
	GetUser(ctx context.Context, request int64) (*GetUserResponse, error)
	DeleteUser(ctx context.Context, request int64) error
	UpdateUser(ctx context.Context, request UpdateUserRequest) error
}
