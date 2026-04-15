package app

import (
	"log/slog"
	"sso/internal/app/grpc"
	"sso/internal/services/auth"
	"sso/internal/storage/sqlite"
	"time"
)

type App struct {
	GRPCSrv *grpcapp.App
}

func New(
	log *slog.Logger,
	grpcPort int,
	storagePath string,
	tokenTTL time.Duration,
) *App {

	storage, err := sqlite.New(storagePath)
	if err != nil {
		panic(err)
	}

	// auth.New(...) принимает три разные зависимости по интерфейсам, а у тебя один объект storage реализует их все.
	authService := auth.New(log, storage, storage, storage, tokenTTL)

	grpcApp := grpcapp.New(log, authService, grpcPort)

	return &App{
		GRPCSrv: grpcApp,
	}

}
