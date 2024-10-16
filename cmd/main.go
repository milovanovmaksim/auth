package main

import (
	"context"
	"log"

	"github.com/joho/godotenv"

	"github.com/milovanovmaksim/auth/cmd/auth_server"
	"github.com/milovanovmaksim/auth/internal/pgsql"
)

const grpcPort = 50051

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	ctx := context.Background()

	db_config, err := pgsql.NewConfigFromEnv()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	postgreSql, err := pgsql.Connect(ctx, db_config)
	if err != nil {
		log.Fatalf("failed to connect to PostgreSQL")
	}

	server := auth_server.NewServer(postgreSql)
	err = server.Start(grpcPort)
	if err != nil {
		log.Fatalf("failed to start a server | err: %v", err)
	}

	defer server.Stop()
}
