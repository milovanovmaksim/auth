package repository

import (
	"database/sql"
	"time"

	"github.com/milovanovmaksim/auth/internal/service"
	desc "github.com/milovanovmaksim/auth/pkg/auth_v1"
)

// GetUserResponse ответ на запрос о получении информации о пользователе.
type GetUserResponse struct {
	ID        int64        `db:"id"`
	Name      string       `db:"username"`
	Email     string       `db:"email"`
	Role      desc.Role    `db:"role"`
	CreatedAt time.Time    `db:"created_at"`
	UpdatedAt sql.NullTime `db:"updated_at"`
}

// Into преобоазует объект в service.GetUserResponse.
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

// CreateUserResponse ответ на запрос о создании нового пользователя.
type CreateUserResponse struct {
	ID int64
}

// Into преобразует объект в service.CreateUserResponse.
func (u CreateUserResponse) Into() service.CreateUserResponse {
	return service.CreateUserResponse{ID: u.ID}
}
