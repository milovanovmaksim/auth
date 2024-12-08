package model

import (
	"time"

	desc "github.com/milovanovmaksim/auth/pkg/auth_v1"
)

// UpdateUserRequest запрос на обновление данных о пользователе.
type UpdateUserRequest struct {
	Name *string
	ID   int64
	Role desc.Role
}

// CreateUserRequest запрос на создание нового пользователя.
type CreateUserRequest struct {
	Name            string
	Email           string
	Password        string
	PasswordConfirm string
	Role            desc.Role
}

// GetUserResponse ответ на запрос о получении информации о пользователе.
type GetUserResponse struct {
	CreatedAt time.Time
	UpdatedAt *time.Time
	Name      string
	Email     string
	ID        int64
	Role      string
}
