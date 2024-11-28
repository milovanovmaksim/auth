package converter

import (
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/milovanovmaksim/auth/internal/service/user/model"
	desc "github.com/milovanovmaksim/auth/pkg/auth_v1"
)

func ToDescFromGetUserResponse(model model.GetUserResponse) desc.GetUserResponse {
	var updatedAt *timestamppb.Timestamp

	if model.UpdatedAt != nil {
		updatedAt = timestamppb.New(*model.UpdatedAt)
	}
	return desc.GetUserResponse{
		User: &desc.User{
			Id:        model.ID,
			Name:      model.Name,
			Email:     model.Email,
			Role:      desc.Role(desc.Role_value[model.Role]),
			CreatedAt: timestamppb.New(model.CreatedAt),
			UpdatedAt: updatedAt,
		},
	}

}
