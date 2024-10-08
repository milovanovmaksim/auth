package auths_erver

import (
	"context"
	"log"

	"github.com/brianvoe/gofakeit"
	desc "github.com/olezhek28/microservices_course_boilerplate/pkg/auth_v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type server struct {
	desc.UnimplementedUserV1Server
}

func (s *server) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	log.Printf("User id: %d", req.GetId())

	return &desc.GetResponse{
		User: &desc.User{
			Id:        req.GetId(),
			Name:      "Maxim",
			Email:     "bla-bla@mail.com",
			Role:      desc.Role_ADMIN,
			CreatedAt: timestamppb.New(gofakeit.Date()),
			UpdatedAt: timestamppb.New(gofakeit.Date()),
		},
	}, nil
}
