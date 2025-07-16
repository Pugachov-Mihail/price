package grpc_app

import (
	"log/slog"
	"skeleton-grpc/internal/app"
	"skeleton-grpc/internal/config"
)

type App struct {
	GRPC *app.App
}

func NewGrpcService(log *slog.Logger, cfg *config.Config) *App {

	appGrpc := app.NewApp(log, cfg.Grpc.Port, cfg.Grpc.Host)
	return &App{appGrpc}
}
