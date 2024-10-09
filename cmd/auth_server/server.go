package auth_server

import (
	"context"
	"log"

	"github.com/brianvoe/gofakeit"
	desc "github.com/olezhek28/microservices_course_boilerplate/pkg/auth_v1"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// Server ...
type Server struct {
	desc.UnimplementedUserV1Server
}

// Get ...
func (s *Server) Get(_ context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
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

// Create ...
func (s *Server) Create(_ context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	log.Printf("Create user: %+v", req.User)
	return &desc.CreateResponse{
		Id: 1,
	}, nil
}

// Update ...
func (s *Server) Update(_ context.Context, req *desc.UpdateRequest) (*emptypb.Empty, error) {
	log.Printf("Update user with id = %d", req.User.GetId())
	return &emptypb.Empty{}, nil
}

// Delete ...
func (s *Server) Delete(_ context.Context, req *desc.DeleteRequest) (*emptypb.Empty, error) {
	log.Printf("Delete user with id = %d", req.GetId())
	return &emptypb.Empty{}, nil
}
