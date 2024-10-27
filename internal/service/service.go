package service

import (
	"context"
	"database/sql"
	"time"

	desc "github.com/milovanovmaksim/auth/pkg/auth_v1"
)

type UserService interface {
	CreateUser(ctx context.Context, request CreateUserRequest) (*CreateUserResponse, error)
	GetUser(ctx context.Context, request int64) (*GetUserResponse, error)
	UpdateUser(ctx context.Context, request UpdateUserRequest) error
}

type UpdateUserRequest struct {
	ID   int64
	Name string
	Role desc.Role
}

type GetUserResponse struct {
	ID        int64
	Name      string
	Email     string
	Role      desc.Role
	CreatedAt time.Time
	UpdatedAt sql.NullTime
}

type CreateUserRequest struct {
	Name            string
	Email           string
	Password        string
	PasswordConfirm string
	Role            desc.Role
}

type CreateUserResponse struct {
	ID int64
}
