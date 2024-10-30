package grpc

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/milovanovmaksim/auth/internal/server"
	"github.com/milovanovmaksim/auth/internal/service"
	desc "github.com/milovanovmaksim/auth/pkg/auth_v1"
)

// Server - cервер аутентификации пользователя.
type Server struct {
	desc.UnimplementedUserV1Server
	grpcConfig server.ServerConfig
	grpcServer *grpc.Server
	service    service.UserService
}

// NewServer создает новый Server объект.
func NewServer(grpcConfig server.ServerConfig, service service.UserService) Server {
	return Server{desc.UnimplementedUserV1Server{}, grpcConfig, nil, service}
}

// GetUser возвращает информацию о пользователе.
func (s *Server) GetUser(ctx context.Context, req *desc.GetUserRequest) (*desc.GetUserResponse, error) {
	user, err := s.service.GetUser(ctx, req.GetId())
	if err != nil {
		log.Printf("failed to get user Server.GetUser || error: %v", err)
		return nil, err
	}

	res := user.To()

	return &res, nil
}

// CreateUser создает нового пользователя.
func (s *Server) CreateUser(ctx context.Context, req *desc.CreateUserRequest) (*desc.CreateUserResponse, error) {
	serviceResponse, err := s.service.CreateUser(ctx, service.CreateUserRequest{
		Name:            req.User.Name,
		Email:           req.User.Email,
		Password:        req.User.Password,
		PasswordConfirm: req.User.PasswordConfirm,
		Role:            req.User.Role,
	})
	if err != nil {
		log.Printf("failed to create new user Server.CreateUser error || %v", err)
		return nil, err
	}

	res := serviceResponse.To()

	return &res, nil
}

// UpdateUser обновляет данные о пользователе.
func (s *Server) UpdateUser(ctx context.Context, req *desc.UpdateUserRequest) (*emptypb.Empty, error) {
	err := s.service.UpdateUser(ctx, service.UpdateUserRequest{
		Name: req.User.Name.GetValue(),
		ID:   req.User.Id,
		Role: req.User.Role,
	})
	if err != nil {
		log.Printf("failed to update user Server.UpdateUser || error: %v", err)
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

// DeleteUser удаляет пользователя.
func (s *Server) DeleteUser(ctx context.Context, req *desc.DeleteUserRequest) (*emptypb.Empty, error) {
	err := s.service.DeleteUser(ctx, req.Id)
	if err != nil {
		log.Printf("failed to delete user Server.DeleteUser || error: %v", err)
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

// Start cтарт сервера аутентификации пользователя.
func (s *Server) Start() error {
	lis, err := net.Listen("tcp", s.grpcConfig.Address())
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		return err
	}

	s.grpcServer = grpc.NewServer()

	reflection.Register(s.grpcServer)
	desc.RegisterUserV1Server(s.grpcServer, s)
	log.Printf("server listening at %v", lis.Addr())

	if err = s.grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
		return err
	}

	defer s.grpcServer.Stop()

	return nil
}

// Stop остановка сервера.
func (s *Server) Stop() {
	if s.grpcServer != nil {
		s.grpcServer.Stop()
	}
}
