package repository

import (
	"database/sql"
	"time"

	desc "github.com/milovanovmaksim/auth/pkg/auth_v1"
)

type GetUserResponse struct {
	ID        int64
	Name      string
	Email     string
	Role      desc.Role
	CreatedAt time.Time
	UpdatedAt sql.NullTime
}

type CreateUserResponse struct {
	ID int64
}
