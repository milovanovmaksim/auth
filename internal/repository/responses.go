package repository

import (
	"database/sql"
	"time"

	"github.com/milovanovmaksim/auth/internal/service"
	desc "github.com/milovanovmaksim/auth/pkg/auth_v1"
)

type GetUserResponse struct {
	ID        int64        `db:"id"`
	Name      string       `db:"username"`
	Email     string       `db:"email"`
	Role      desc.Role    `db:"role"`
	CreatedAt time.Time    `db:"created_at"`
	UpdatedAt sql.NullTime `db:"updated_at"`
}

func (u GetUserResponse) Into() service.GetUserResponse {
	return service.GetUserResponse{
		ID:        u.ID,
		Name:      u.Name,
		Email:     u.Email,
		Role:      u.Role,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}

type CreateUserResponse struct {
	ID int64
}

func (u CreateUserResponse) Into() service.CreateUserResponse {
	return service.CreateUserResponse{ID: u.ID}
}
