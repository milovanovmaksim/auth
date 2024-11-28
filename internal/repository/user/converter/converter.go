package converter

import (
	"time"

	"github.com/milovanovmaksim/auth/internal/repository/user/model"
	serviceModel "github.com/milovanovmaksim/auth/internal/service/user/model"
)

func ToServiceFromGetUserResponse(model model.GetUserResponse) serviceModel.GetUserResponse {
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
