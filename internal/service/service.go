package service

import (
	"context"
	"database/sql"
	"time"

	desc "github.com/milovanovmaksim/auth/pkg/auth_v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type UserService interface {
	CreateUser(ctx context.Context, request CreateUserRequest) (*CreateUserResponse, error)
	GetUser(ctx context.Context, request int64) (*GetUserResponse, error)
	UpdateUser(ctx context.Context, request UpdateUserRequest) error
	DeleteUser(ctx context.Context, request int64) error
}

type To[T any] interface {
	To() T
}

type UpdateUserRequest struct {
	Name string
	ID   int64
	Role desc.Role
}

type GetUserResponse struct {
	CreatedAt time.Time
	UpdatedAt sql.NullTime
	Name      string
	Email     string
	ID        int64
	Role      desc.Role
}

func (u GetUserResponse) To() desc.GetUserResponse {
	return desc.GetUserResponse{
		User: &desc.User{
			Id:        u.ID,
			Name:      u.Name,
			Email:     u.Email,
			Role:      u.Role,
			CreatedAt: timestamppb.New(u.CreatedAt),
			UpdatedAt: timestamppb.New(u.UpdatedAt.Time),
		},
	}
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

func (u CreateUserResponse) To() desc.CreateUserResponse {
	return desc.CreateUserResponse{Id: u.ID}
}
