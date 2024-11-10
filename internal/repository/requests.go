package repository

import "database/sql"

// UpdateUserRequest запрос на обновление данных пользователя.
type UpdateUserRequest struct {
	ID   int64
	Name sql.NullString
	Role string
}

// CreateUserRequest запрос на создание нового пользователя.
type CreateUserRequest struct {
	Name         string
	Email        string
	HashPassword string
	Role         string
}
