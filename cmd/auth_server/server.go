package auth_server

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net"
	"time"

	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	grpc_config "github.com/milovanovmaksim/auth/internal/config"
	"github.com/milovanovmaksim/auth/internal/pgsql"
	desc "github.com/milovanovmaksim/auth/pkg/auth_v1"
)

// Server - cервер аутентификации пользователя.
type Server struct {
	postgreSQL *pgsql.PostgreSQL
	grpcConfig *grpc_config.GrpcConfig
	grpcServer *grpc.Server
	desc.UnimplementedUserV1Server
}

// NewServer создает новый Server объект.
func NewServer(postgreSQL *pgsql.PostgreSQL, grpcConfig *grpc_config.GrpcConfig) Server {
	return Server{postgreSQL, grpcConfig, nil, desc.UnimplementedUserV1Server{}}
}

// GetUser возвращает информацию о пользователе.
func (s *Server) GetUser(ctx context.Context, req *desc.GetUserRequest) (*desc.GetUserResponse, error) {
	var id int64
	var role, name, email string
	var createdAt time.Time
	var updatedAt sql.NullTime

	pool := s.postgreSQL.GetPool()

	row := pool.QueryRow(ctx, "SELECT id, username, email, role, created_at, updated_at FROM users WHERE id = $1", req.GetId())

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
func (s *Server) CreateUser(ctx context.Context, req *desc.CreateUserRequest) (*desc.CreateUserResponse, error) {
	var id int64

	if req.User.Password != req.User.PasswordConfirm {
		return nil, fmt.Errorf("password and password_confirm should be the same")
	}

	hashPassword, err := s.hashPassword(req.User.Password)
	if err != nil {
		log.Printf("failed to get hash for password || err: %v", err)
		return nil, fmt.Errorf("internal error")
	}

	pool := s.postgreSQL.GetPool()

	err = pool.QueryRow(ctx, "INSERT INTO users (username, email, password, role) VALUES($1, $2, $3, $4) returning id",
		req.User.Name, req.User.Email, hashPassword, req.User.GetRole().String()).Scan(&id)
	if err != nil {
		fmt.Printf("failed to insert user: %v", err)
		return nil, err
	}

	return &desc.CreateUserResponse{Id: id}, nil

}

func (s *Server) hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// UpdateUser обновляет данные о пользователе.
func (s *Server) UpdateUser(ctx context.Context, req *desc.UpdateUserRequest) (*emptypb.Empty, error) {
	var name sql.NullString
	var role sql.NullString

	pool := s.postgreSQL.GetPool()

	if req.User.GetName().GetValue() != "" {
		name = sql.NullString{String: req.User.GetName().GetValue(), Valid: true}
	}

	if req.GetUser().GetRole().Number() != 0 {
		role = sql.NullString{String: req.GetUser().GetRole().String(), Valid: true}
	}

	_, err := pool.Exec(ctx, "UPDATE users SET username = COALESCE($1, username), role = COALESCE($2, role)", name, role)
	if err != nil {
		fmt.Printf("failed to update user: %v", err)
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

// DeleteUser удаляет пользователя.
func (s *Server) DeleteUser(ctx context.Context, req *desc.DeleteUserRequest) (*emptypb.Empty, error) {
	pool := s.postgreSQL.GetPool()

	_, err := pool.Exec(ctx, "DELETE FROM USERS WHERE id = $1", req.Id)
	if err != nil {
		fmt.Printf("failed to delete user: %v", err)
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
