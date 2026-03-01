package main

import (
	"os"
	"fmt"
	"log/slog"
	"sso/internal/config"
	"sso/internal/lib/logger/handlers/slogpretty"
	"strconv"
)

const (
	envLocal = "local"
	envDev = "dev"
	envProd = "prod"
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

	//TODO: инициализировать логгер

	//TODO: инициализировать приложение
}

func setupLogger(cfg *config.Config)*slog.Logger {
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