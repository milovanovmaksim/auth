package app

import (
	"context"

	"github.com/milovanovmaksim/auth/internal/closer"
	"github.com/milovanovmaksim/auth/internal/server/grpc"
)

type App struct {
	diConteiner diContainer
	grpcServer  grpc.Server
}

func NewApp(ctx context.Context) (*App, error) {
	app := &App{}

	err := app.initDeps(ctx)
	if err != nil {
		return nil, err
	}

	return app, nil
}

func (a *App) Run() error {
	defer func() {
		closer.CloseAll()
		closer.Wait()
	}()

	return a.grpcServer.Start()
}

func (a *App) initDeps(ctx context.Context) error {
	a.initGRPCServer(ctx)
	a.initdiContainer()

	return nil
}

func (a *App) initGRPCServer(ctx context.Context) error {
	a.grpcServer = grpc.NewServer(a.diConteiner.GRPCConfig(), a.diConteiner.UserService(ctx))

	return nil
}

func (a *App) initdiContainer() error {
	a.diConteiner = newDiContainer()

	return nil
}
