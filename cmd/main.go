package main

import (
	"context"
	"log"

	"github.com/milovanovmaksim/auth/internal/app"
)

func main() {
	ctx := context.Background()

	app, err := app.NewApp(ctx, ".env")
	if err != nil {
		log.Fatalf("failed to init app: %s", err.Error())
	}

	err = app.Run()
	if err != nil {
		log.Fatalf("failed to run app: %s", err.Error())
	}
}
