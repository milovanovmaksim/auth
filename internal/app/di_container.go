package app

import (
	"context"
	"log"

	"github.com/milovanovmaksim/auth/internal/client/database"
	"github.com/milovanovmaksim/auth/internal/client/database/postgresql"
	"github.com/milovanovmaksim/auth/internal/client/database/transaction"
	"github.com/milovanovmaksim/auth/internal/closer"
	"github.com/milovanovmaksim/auth/internal/repository"
	userRepo "github.com/milovanovmaksim/auth/internal/repository/user"
	"github.com/milovanovmaksim/auth/internal/server"
	"github.com/milovanovmaksim/auth/internal/server/grpc"
	"github.com/milovanovmaksim/auth/internal/service"
	userService "github.com/milovanovmaksim/auth/internal/service/user"
)

type diContainer struct {
	userRepository repository.UserRepository
	userService    service.UserService
	dbClient       database.Client
	pgConfig       database.DBConfig
	grpcConfig     server.Config
	txManager      database.TxManager
}

func newDiContainer() diContainer {
	return diContainer{}
}

// UserRepository врозвращает объект, удовлетворяющий интерфейсу repository.UserRepository.
func (di *diContainer) UserRepository(ctx context.Context) repository.UserRepository {
	if di.userRepository == nil {
		di.userRepository = userRepo.NewUserRepository(di.DBClient(ctx))
	}

	return di.userRepository
}

// UserService возврашщает объект, удовлетворяющий интерфейсу service.UserService.
func (di *diContainer) UserService(ctx context.Context) service.UserService {
	if di.userService == nil {
		di.userService = userService.NewUserService(di.UserRepository(ctx), di.TxManager(ctx))
	}

	return di.userService
}

// DBConfig возвращает объект, удовлетворяющий интерфейсу database.DBConfig.
func (di *diContainer) DBConfig() database.DBConfig {
	if di.pgConfig == nil {
		config, err := postgresql.NewConfigFromEnv()
		if err != nil {
			log.Fatalf("failed to get DB config: %v", err.Error())
		}

		di.pgConfig = config
	}

	return di.pgConfig
}

// GRPCConfig возвращает объект, удовлетворяющий интерфейсу server.ServerConfig.
func (di *diContainer) GRPCConfig() server.Config {
	if di.grpcConfig == nil {
		cfg, err := grpc.NewGrpcConfigFromEnv()
		if err != nil {
			log.Fatalf("failed to get grpc config: %v", err.Error())
		}

		di.grpcConfig = cfg
	}

	return di.grpcConfig
}

// DBClient возвращает объект, удовлетвoряющий интерфейсу database.Client.
func (di *diContainer) DBClient(ctx context.Context) database.Client {
	if di.dbClient == nil {
		pg, err := postgresql.Connect(ctx, di.DBConfig())
		if err != nil {
			log.Fatalf("failed to connect to PostgreSQL server")
		}

		dbClient := postgresql.NewClient(pg)
		di.dbClient = dbClient

		closer.Add(dbClient.Close)
	}

	return di.dbClient
}

// TxManager возвращает объект, удовлетвoряющий интерфейсу database.TxManager.
func (di *diContainer) TxManager(ctx context.Context) database.TxManager {
	if di.txManager == nil {
		di.txManager = transaction.NewTransactionManager(di.DBClient(ctx).DB())
	}

	return di.txManager
}
