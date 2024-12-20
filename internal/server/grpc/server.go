package grpc

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/milovanovmaksim/auth/internal/closer"
	"github.com/milovanovmaksim/auth/internal/server"
	"github.com/milovanovmaksim/auth/internal/service"
	"github.com/milovanovmaksim/auth/internal/service/user/converter"
	serviceModel "github.com/milovanovmaksim/auth/internal/service/user/model"
	desc "github.com/milovanovmaksim/auth/pkg/auth_v1"
)

// Server - cервер аутентификации пользователей.
type Server struct {
	desc.UnimplementedUserV1Server
	grpcConfig server.Config
	grpcServer *grpc.Server
	service    service.UserService
}

// NewServer создает новый Server объект.
func NewServer(grpcConfig server.Config, service service.UserService) Server {
	return Server{desc.UnimplementedUserV1Server{}, grpcConfig, nil, service}
}

// GetUser возвращает информацию о пользователе.
func (s *Server) GetUser(ctx context.Context, req *desc.GetUserRequest) (*desc.GetUserResponse, error) {
	user, err := s.service.GetUser(ctx, req.GetId())
	if err != nil {
		log.Printf("failed to get user Server.GetUser: %v", err)
		return nil, err
	}

	res := converter.ToDescFromGetUserResponse(*user)

	return &res, nil
}

// CreateUser создает нового пользователя.
func (s *Server) CreateUser(ctx context.Context, req *desc.CreateUserRequest) (*desc.CreateUserResponse, error) {
	userID, err := s.service.CreateUser(ctx, serviceModel.CreateUserRequest{
		Name:            req.User.Name,
		Email:           req.User.Email,
		Password:        req.User.Password,
		PasswordConfirm: req.User.PasswordConfirm,
		Role:            req.User.Role,
	})
	if err != nil {
		log.Printf("failed to create new user Server.CreateUser: %v", err)
		return nil, err
	}

	return &desc.CreateUserResponse{Id: userID}, nil
}

// UpdateUser обновляет данные о пользователе.
func (s *Server) UpdateUser(ctx context.Context, req *desc.UpdateUserRequest) (*emptypb.Empty, error) {
	var name *string

	if req.User.GetName() != nil {
		name = &req.User.GetName().Value
	}

	err := s.service.UpdateUser(ctx, serviceModel.UpdateUserRequest{
		Name: name,
		ID:   req.User.Id,
		Role: req.User.Role,
	})
	if err != nil {
		log.Printf("failed to update user Server.UpdateUser: %v", err)
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

// DeleteUser удаляет пользователя.
func (s *Server) DeleteUser(ctx context.Context, req *desc.DeleteUserRequest) (*emptypb.Empty, error) {
	err := s.service.DeleteUser(ctx, req.Id)
	if err != nil {
		log.Printf("failed to delete user Server.DeleteUser: %v", err)
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

// Start cтарт сервера аутентификации пользователя.
func (s *Server) Start() error {
	lis, err := net.Listen("tcp", s.grpcConfig.Address())
	if err != nil {
		log.Printf("failed to listen: %v", err)
		return err
	}

	closer.Add(lis.Close)

	s.grpcServer = grpc.NewServer()

	reflection.Register(s.grpcServer)
	desc.RegisterUserV1Server(s.grpcServer, s)
	log.Printf("server listening at %v", lis.Addr())

	if err = s.grpcServer.Serve(lis); err != nil {
		log.Printf("failed to serve: %v", err)
		return err
	}

	closer.Add(func() error {
		s.grpcServer.Stop()
		return nil
	})

	return nil
}
