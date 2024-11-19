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
	Name         sql.NullString
	Email        sql.NullString
	HashPassword sql.NullString
	Role         sql.NullString
}

// NewCreateUserRequest создает новый объект.
func NewCreateUserRequest(name string, email string, hashPassword string, role string) CreateUserRequest {
	Name := sql.NullString{String: "", Valid: false}
	Email := sql.NullString{String: "", Valid: false}
	HashPassword := sql.NullString{String: "", Valid: false}
	Role := sql.NullString{String: "", Valid: false}

	if name != "" {
		Name.String = name
		Name.Valid = true
	}

	if email != "" {
		Email.String = email
		Email.Valid = true
	}

	if hashPassword != "" {
		HashPassword.String = hashPassword
		HashPassword.Valid = true
	}

	if role != "" {
		Role.String = role
		Role.Valid = true
	}

	return CreateUserRequest{Name, Email, HashPassword, Role}
}
