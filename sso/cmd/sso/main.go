package main

import (
	"fmt"
	"log/slog"
	"os"
	"sso/internal/config"
	"sso/internal/lib/logger/handlers/slogpretty"
	"strconv"

	"sso/internal/app"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	cfg := config.MustLoad()

	fmt.Println(cfg)

	log := setupLogger(cfg)

	log.Info("starting app",
		slog.String("env", cfg.Env),
		slog.String("port", strconv.Itoa(cfg.GRPC.Port)),
	)

	log.Debug("debug message")

	log.Error("error message")

	log.Warn("warn message")

	application := app.New(log, cfg.GRPC.Port, cfg.StoragePath, cfg.GRPC.TokenTTL)

	application.GRPCSrv.Run()

	//TODO: инициализировать приложение

	//TODO: Запустить gRPC сервер приложения
}

func setupLogger(cfg *config.Config) *slog.Logger {
	var log *slog.Logger

	switch cfg.Env {
	case envLocal:
		log = setupPrettySlog()
	case envDev:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envProd:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}
	return log
}

func setupPrettySlog() *slog.Logger {

	opts := slogpretty.PrettyHandlerOptions{
		SlogOpts: &slog.HandlerOptions{
			Level: slog.LevelDebug,
		},
	}

	handler := opts.NewPrettyHandler(os.Stdout)

	return slog.New(handler)
}
