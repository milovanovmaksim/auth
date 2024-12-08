package repository

import (
	"context"

	"github.com/milovanovmaksim/auth/internal/repository/user/model"
)

// UserRepository интерфейс, определяющий набор методов CRUD для работы с БД.
type UserRepository interface {
	CreateUser(ctx context.Context, request model.CreateUserRequest) (int64, error)
	GetUser(ctx context.Context, request int64) (*model.GetUserResponse, error)
	DeleteUser(ctx context.Context, request int64) error
	UpdateUser(ctx context.Context, request model.UpdateUserRequest) error
}
