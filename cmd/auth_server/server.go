package auth_server

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/milovanovmaksim/auth/internal/config"
	"github.com/milovanovmaksim/auth/internal/pgsql"
	desc "github.com/milovanovmaksim/auth/pkg/auth_v1"
)

// Server - cервер аутентификации пользователя.
type Server struct {
	postgreSql *pgsql.PostgreSQL
	grpcConfig *config.GrpcConfig
	ctx        context.Context
	desc.UnimplementedUserV1Server
}

func NewServer(postgreSql *pgsql.PostgreSQL, grpcConfig *config.GrpcConfig, ctx context.Context) Server {
	return Server{postgreSql, grpcConfig, ctx, desc.UnimplementedUserV1Server{}}
}

// GetUser возвращает информацию о пользователе.
func (s *Server) GetUser(_ context.Context, req *desc.GetUserRequest) (*desc.GetUserResponse, error) {
	var id int64
	var role, name, email string
	var createdAt time.Time
	var updatedAt sql.NullTime

	pool := s.postgreSql.GetPool()

	row := pool.QueryRow(s.ctx, "SELECT id, username, email, role, created_at, updated_at FROM users WHERE id = $1", req.GetId())

	err := row.Scan(&id, &name, &email, &role, &createdAt, &updatedAt)
	if err != nil {
		fmt.Printf("failed to get user: %v", err)
		return nil, err
	}

	return &desc.GetUserResponse{
		User: &desc.User{
			Id:        id,
			Name:      name,
			Email:     email,
			Role:      desc.Role(desc.Role_value[role]),
			CreatedAt: timestamppb.New(createdAt),
			UpdatedAt: timestamppb.New(updatedAt.Time),
		},
	}, nil
}

// CreateUser создает нового пользователя.
func (s *Server) CreateUser(_ context.Context, req *desc.CreateUserRequest) (*desc.CreateUserResponse, error) {
	var id int64

	pool := s.postgreSql.GetPool()

	err := pool.QueryRow(s.ctx, "INSERT INTO users (username, email, password, role) VALUES($1, $2, $3, $4) returning id",
		req.User.Name, req.User.Email, req.User.Password, req.User.Role.String()).Scan(&id)
	if err != nil {
		fmt.Printf("failed to insert user: %v", err)
		return nil, err
	}

	return &desc.CreateUserResponse{Id: id}, nil

}

// UpdateUser обновляет данные о пользователе.
func (s *Server) UpdateUser(_ context.Context, req *desc.UpdateUserRequest) (*emptypb.Empty, error) {
	log.Printf("Update user with id = %d", req.User.GetId())
	return &emptypb.Empty{}, nil
}

// DeleteUser удаляет пользователя.
func (s *Server) DeleteUser(_ context.Context, req *desc.DeleteUserRequest) (*emptypb.Empty, error) {
	pool := s.postgreSql.GetPool()

	_, err := pool.Exec(s.ctx, "DELETE FROM USERS WHERE id = $1", req.Id)
	if err != nil {
		fmt.Printf("failed to delete user: %v", err)
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
