package app

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log/slog"
	"net"
	"skeleton-grpc/internal/api"
	"skeleton-grpc/internal/service"
)

type App struct {
	app     *grpc.Server
	log     *slog.Logger
	port    string
	appName string
}

func NewApp(log *slog.Logger, port, appName string) *App {
	grpcServer := grpc.NewServer()

	srv := service.NewMockDataService(log)

	api.RegisterPriceServer(grpcServer, srv)
	reflection.Register(grpcServer)
	return &App{grpcServer, log, port, appName}
}

func (a *App) run() error {
	log := a.log.With("Init grpc server", a.appName)

	l, err := net.Listen("tcp", ":"+a.port)
	if err != nil {
		log.Error("failed to listen", "err", err)
		return err
	}
	log.Info("starting tcp", slog.String("port", a.port))

	if err := a.app.Serve(l); err != nil {
		log.Error("failed to serve", "err", err)
		return err
	}
	return nil
}

func (a *App) Stop() {
	a.log.With("Stopping grpc server")
	a.app.GracefulStop()
	a.log.Info("grpc server stopped")
}

func (a *App) MustRun() {
	if err := a.run(); err != nil {
		panic(err)
	}
}
