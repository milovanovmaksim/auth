package repository

import "database/sql"

// UpdateUserRequest запрос на обновление данных пользователя.
type UpdateUserRequest struct {
	Role string
	Name sql.NullString
	ID   int64
}

// CreateUserRequest запрос на создание нового пользователя.
type CreateUserRequest struct {
	Name         string
	Email        string
	HashPassword string
	Role         string
}
