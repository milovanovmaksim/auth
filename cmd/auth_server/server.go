package auth_server

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/brianvoe/gofakeit"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/milovanovmaksim/auth/internal/pgsql"
	desc "github.com/milovanovmaksim/auth/pkg/auth_v1"
)

// Server - cервер аутентификации пользователя.
type Server struct {
	postgreSql *pgsql.PostgreSQL
	desc.UnimplementedUserV1Server
}

func NewServer(postgreSql *pgsql.PostgreSQL) Server {
	return Server{postgreSql, desc.UnimplementedUserV1Server{}}
}

// GetUser возвращает информацию о пользователе.
func (s *Server) GetUser(_ context.Context, req *desc.GetUserRequest) (*desc.GetUserResponse, error) {
	log.Printf("User id: %d", req.GetId())

	return &desc.GetUserResponse{
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

// CreateUser создает нового пользователя.
func (s *Server) CreateUser(_ context.Context, req *desc.CreateUserRequest) (*desc.CreateUserResponse, error) {
	log.Printf("Create user: %+v", req.User)
	return &desc.CreateUserResponse{
		Id: 1,
	}, nil
}

// UpdateUser обновляет данные о пользователе.
func (s *Server) UpdateUser(_ context.Context, req *desc.UpdateUserRequest) (*emptypb.Empty, error) {
	log.Printf("Update user with id = %d", req.User.GetId())
	return &emptypb.Empty{}, nil
}

// DeleteUser удаляет пользователя.
func (s *Server) DeleteUser(_ context.Context, req *desc.DeleteUserRequest) (*emptypb.Empty, error) {
	log.Printf("Delete user with id = %d", req.GetId())
	return &emptypb.Empty{}, nil
}

// Start cтарт сервера аутентификации пользователя.
func (s *Server) Start(grpcPort int64) error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", grpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		return err
	}

	server := grpc.NewServer()
	reflection.Register(server)
	desc.RegisterUserV1Server(server, s)
	log.Printf("server listening at %v", lis.Addr())

	if err = server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
		return err
	}

	return nil
}

func (s *Server) Stop() {
	s.postgreSql.Close()
}
