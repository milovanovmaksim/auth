package converter

import (
	"time"

	repoModel "github.com/milovanovmaksim/auth/internal/repository/user/model"
	serviceModel "github.com/milovanovmaksim/auth/internal/service/user/model"
)

// ToServiceFromGetUserResponse конвертирует repoModel.GetUserResponse в serviceModel.GetUserResponse.
func ToServiceFromGetUserResponse(value repoModel.GetUserResponse) serviceModel.GetUserResponse {
	var updatedAt *time.Time

	if value.UpdatedAt.Valid {
		updatedAt = &value.UpdatedAt.Time
	}

	return serviceModel.GetUserResponse{
		ID:        value.ID,
		Name:      value.Name,
		Email:     value.Email,
		Role:      value.Role,
		CreatedAt: value.CreatedAt,
		UpdatedAt: updatedAt,
	}
}
