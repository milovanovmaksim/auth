package main

import (
	"context"
	"log"

	"github.com/joho/godotenv"

	"github.com/milovanovmaksim/auth/internal/app"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("failed to load config || err: %v", err)
	}

	ctx := context.Background()

	app, err := app.NewApp(ctx)
	if err != nil {
		log.Fatalf("failed to init app: %s", err.Error())
	}

	err = app.Run()
	if err != nil {
		log.Fatalf("failed to run app: %s", err.Error())
	}
}
