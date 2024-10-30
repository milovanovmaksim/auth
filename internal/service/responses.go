package service

import (
	"database/sql"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"

	desc "github.com/milovanovmaksim/auth/pkg/auth_v1"
)

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

type CreateUserResponse struct {
	ID int64
}

func (u CreateUserResponse) To() desc.CreateUserResponse {
	return desc.CreateUserResponse{Id: u.ID}
}
