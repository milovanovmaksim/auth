package user

import (
	"context"
	"log"

	"github.com/milovanovmaksim/auth/internal/client/database"
	"github.com/milovanovmaksim/auth/internal/repository/user/model"
)

// CreateUser создает нового пользователя.
func (r *userRepositoryImpl) CreateUser(ctx context.Context, request model.CreateUserRequest) (int64, error) {
	var response int64

	queryRow := "INSERT INTO users (username, email, password, role) VALUES($1, $2, $3, $4) returning id"

	query := database.Query{Name: "Create user", QueryRaw: queryRow}

	err := r.db.DB().ScanOneContext(ctx, &response, query, request.Name, request.Email, request.HashPassword, request.Role)
	if err != nil {
		log.Printf("failed to insert user userRepositoryImpl.CreateUser: %v", err)
		return 0, err
	}

	return response, nil
}
