package main

import (
	"context"
	"log"

	"github.com/joho/godotenv"

	"github.com/milovanovmaksim/auth/cmd/auth_server"
	grpc_config "github.com/milovanovmaksim/auth/internal/config"
	"github.com/milovanovmaksim/auth/internal/pgsql"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("failed to load config || err: %v", err)
	}

	ctx := context.Background()

	dbConfig, err := pgsql.NewConfigFromEnv()
	if err != nil {
		log.Fatalf("failed to load config || err: %v", err)
	}

	grpcConfig, err := grpc_config.NewGrpcConfigFromEnv()
	if err != nil {
		log.Fatalf("failed to load grpc config || err: %v", err)
	}

	postgreSql, err := pgsql.Connect(ctx, dbConfig)
	if err != nil {
		log.Fatalf("failed to connect to PostgreSQL || err: %v", err)
	}

	defer postgreSql.Close()

	server := auth_server.NewServer(postgreSql, grpcConfig)
	err = server.Start()
	if err != nil {
		log.Fatalf("failed to start a server || err: %v", err)
	}
}
