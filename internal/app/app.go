package app

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log/slog"
	"skeleton-grpc/internal/api"
)

type App struct {
	app     *grpc.Server
	log     *slog.Logger
	port    string
	appName string
}

func NewApp(app *grpc.Server, log *slog.Logger, port, appName string) *App {
	grpcServer := grpc.NewServer()
	api.RegisterPriceServer(grpcServer)
	reflection.Register(grpcServer)
	return &App{}
}
