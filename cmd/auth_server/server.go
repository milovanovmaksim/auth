package auth_server

import (
	"context"
	"log"

	"github.com/brianvoe/gofakeit"
	desc "github.com/olezhek28/microservices_course_boilerplate/pkg/auth_v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Server struct {
	desc.UnimplementedUserV1Server
}

func (s *Server) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
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

func (s *Server) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	log.Printf("Create user: %+v", req.User)
	return &desc.CreateResponse{
		Id: 1,
	}, nil
}

func (s *Server) Update(ctx context.Context, req *desc.UpdateRequest) error {
	log.Printf("Update user with id = %d", req.GetId())
	return nil
}

func (s *Server) Delete(ctx context.Context, req *desc.DeleteRequest) error {
	log.Printf("Delete user with id = %d", req.GetId())
	return nil
}
