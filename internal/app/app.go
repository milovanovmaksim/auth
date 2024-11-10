package app

import (
	"context"

	"github.com/milovanovmaksim/auth/internal/closer"
	"github.com/milovanovmaksim/auth/internal/server/grpc"
)

// App приложение для аутентификации пользователей.
type App struct {
	diConteiner diContainer
	grpcServer  grpc.Server
}


// NeaApp создает новый объект App.
func NewApp(ctx context.Context) (*App, error) {
	app := &App{}

	err := app.initDeps(ctx)
	if err != nil {
		return nil, err
	}

	return app, nil
}


// Run запускает приложение.
func (a *App) Run() error {
	defer func() {
		closer.CloseAll()
		closer.Wait()
	}()

	return a.grpcServer.Start()
}

func (a *App) initDeps(ctx context.Context) error {
	inits := []func(context.Context) error{
		a.initGRPCServer,
		a.initdiContainer,
	}
	for _, f := range inits {
		err := f(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *App) initGRPCServer(ctx context.Context) error {
	a.grpcServer = grpc.NewServer(a.diConteiner.GRPCConfig(), a.diConteiner.UserService(ctx))

	return nil
}

func (a *App) initdiContainer(_ context.Context) error {
	a.diConteiner = newDiContainer()

	return nil
}
