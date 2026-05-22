package app

import (
	grpcapp "auth-service/internal/app/grpc"
	"log/slog"
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
	// TODO: инициализировать хранилище

	// TODO: инициализировать auth service (auth)

	grpcApp := grpcapp.New(log, grpcPort)

	return &App{
		GRPCSrv: grpcApp,
	}
}
