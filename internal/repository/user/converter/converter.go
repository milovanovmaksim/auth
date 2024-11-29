package converter

import (
	"time"

	repoModel "github.com/milovanovmaksim/auth/internal/repository/user/model"
	serviceModel "github.com/milovanovmaksim/auth/internal/service/user/model"
)

// ToServiceFromGetUserResponse конвертирует repoModel.GetUserResponse в serviceModel.GetUserResponse.
func ToServiceFromGetUserResponse(model repoModel.GetUserResponse) serviceModel.GetUserResponse {
	var updatedAt *time.Time

	if model.UpdatedAt.Valid {
		updatedAt = &model.UpdatedAt.Time
	}

	return serviceModel.GetUserResponse{
		ID:        model.ID,
		Name:      model.Name,
		Email:     model.Email,
		Role:      model.Role,
		CreatedAt: model.CreatedAt,
		UpdatedAt: updatedAt,
	}
}
