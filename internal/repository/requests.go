package repository

import "database/sql"

type UpdateUserRequest struct {
	ID   int64
	Name sql.NullString
	Role string
}

type CreateUserRequest struct {
	Name         string
	Email        string
	HashPassword string
	Role         string
}
