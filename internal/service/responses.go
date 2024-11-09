package service

import (
	"database/sql"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"

	desc "github.com/milovanovmaksim/auth/pkg/auth_v1"
)

// GetUserResponse ответ на запрос о получении информации о пользователе.
type GetUserResponse struct {
	CreatedAt time.Time
	UpdatedAt sql.NullTime
	Name      string
	Email     string
	ID        int64
	Role      desc.Role
}

// Into преобразует в объект типа desc.GetUserResponse.
func (u GetUserResponse) Into() desc.GetUserResponse {
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

// CreateUserResponse ответ на запрос о создании нового пользователя.
type CreateUserResponse struct {
	ID int64
}

// Into преобразует в объект типа desc.CreateUserResponse.
func (u CreateUserResponse) Into() desc.CreateUserResponse {
	return desc.CreateUserResponse{Id: u.ID}
}
