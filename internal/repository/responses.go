package repository

import (
	"database/sql"
	"time"

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

type CreateUserResponse struct {
	ID int64
}
