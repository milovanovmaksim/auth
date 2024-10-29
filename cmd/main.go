package main

import (
	"context"
	"log"

	"github.com/joho/godotenv"

	"github.com/milovanovmaksim/auth/cmd/auth_server"
	"github.com/milovanovmaksim/auth/internal/client/database/postgresql"
	grpcConfig "github.com/milovanovmaksim/auth/internal/config"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("failed to load config || err: %v", err)
	}

	ctx := context.Background()

	dbConfig, err := postgresql.NewConfigFromEnv()
	if err != nil {
		log.Fatalf("failed to load config || err: %v", err)
	}

	grpcConfig, err := grpcConfig.NewGrpcConfigFromEnv()
	if err != nil {
		log.Fatalf("failed to load grpc config || err: %v", err)
	}

	postgreSQL, err := postgresql.Connect(ctx, dbConfig)
	if err != nil {
		log.Fatalf("failed to connect to PostgreSQL || err: %v", err)
	}

	defer postgreSQL.Close()

	server := auth_server.NewServer(postgreSQL, grpcConfig)
	err = server.Start()
	if err != nil {
		log.Fatalf("failed to start a server || err: %v", err)
	}

	defer server.Stop()
}
