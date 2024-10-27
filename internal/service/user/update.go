package user

import (
	"context"
	"database/sql"
	"log"

	"github.com/milovanovmaksim/auth/internal/repository"
	"github.com/milovanovmaksim/auth/internal/service"
)

func (s *userServiceImpl) UpdateUser(ctx context.Context, request service.UpdateUserRequest) error {
	name := sql.NullString{String: "", Valid: false}

	if request.Name != "" {
		name.String = request.Name
		name.Valid = true
	}

	err := s.userRepository.UpdateUser(ctx, repository.UpdateUserRequest{
		ID:   request.ID,
		Name: name,
		Role: request.Role,
	})

	if err != nil {
		log.Printf("failed to update user userServiceImpl.UpdateUser || err: %v", err)
		return err
	}

	return nil

}
