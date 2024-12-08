package converter

import (
	"google.golang.org/protobuf/types/known/timestamppb"

	serviceModel "github.com/milovanovmaksim/auth/internal/service/user/model"
	desc "github.com/milovanovmaksim/auth/pkg/auth_v1"
)

// ToDescFromGetUserResponse конвертирует serviceModel.GetUserResponse в desc.GetUserResponse.
func ToDescFromGetUserResponse(value serviceModel.GetUserResponse) desc.GetUserResponse {
	var updatedAt *timestamppb.Timestamp

	if value.UpdatedAt != nil {
		updatedAt = timestamppb.New(*value.UpdatedAt)
	}

	return desc.GetUserResponse{
		User: &desc.User{
			Id:        value.ID,
			Name:      value.Name,
			Email:     value.Email,
			Role:      desc.Role(desc.Role_value[value.Role]),
			CreatedAt: timestamppb.New(value.CreatedAt),
			UpdatedAt: updatedAt,
		},
	}
}
