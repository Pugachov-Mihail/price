package main

import (
	"log/slog"
	"os"
	"os/signal"
	grpc_app "skeleton-grpc/internal/app/grpc"
	"skeleton-grpc/internal/config"
	"syscall"
)

func main() {
	cfg := config.MustLoad()
	log := config.SetupLoger(cfg)

	log.Debug("Init App")

	application := grpc_app.NewGrpcService(log, cfg)

	go application.GRPC.MustRun()

	stop := make(chan os.Signal, 1)

	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)
	sign := <-stop

	log.Info("stop app", slog.String("signal", sign.String()))

	application.GRPC.Stop()
	log.Info("app stop")
}
