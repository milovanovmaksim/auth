package model

import (
	"database/sql"
	"time"
)

// UpdateUserRequest запрос на обновление данных пользователя.
type UpdateUserRequest struct {
	Role string
	Name *string
	ID   int64
}

// CreateUserRequest запрос на создание нового пользователя.
type CreateUserRequest struct {
	Name         string
	Email        string
	HashPassword string
	Role         string
}

// GetUserResponse ответ на запрос о получении информации о пользователе.
type GetUserResponse struct {
	CreatedAt time.Time    `db:"created_at"`
	UpdatedAt sql.NullTime `db:"updated_at"`
	Name      string       `db:"username"`
	Email     string       `db:"email"`
	ID        int64        `db:"id"`
	Role      string       `db:"role"`
}
