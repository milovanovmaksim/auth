package service

import (
	"context"

	"github.com/milovanovmaksim/auth/internal/service/user/model"
)

// UserService интерфейс, отвечающий за бизнес логику приложения.
type UserService interface {
	CreateUser(ctx context.Context, request model.CreateUserRequest) (int64, error)
	GetUser(ctx context.Context, request int64) (*model.GetUserResponse, error)
	UpdateUser(ctx context.Context, request model.UpdateUserRequest) error
	DeleteUser(ctx context.Context, request int64) error
}
