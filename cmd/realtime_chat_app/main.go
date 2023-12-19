package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	authgrpc "realtime_chat_app/internal/clients/auth/grpc"
	"realtime_chat_app/internal/hanlers"
	"realtime_chat_app/internal/service"
	"realtime_chat_app/internal/storage"
)

const (
	envLocal = "local"
	envDev   = "dev"
)

func main() {
	// TODO: config

	logger := initLogger(envLocal)
	db := storage.MustStorage()
	serv := service.NewService(db)
	client, _ := authgrpc.NewClient(context.Background(), "localhost:5001")
	handler := hanlers.MustHandler(client, serv, logger)

	router, _ := handler.InitRoutes()

	//http.Handle("/", router)
	http.ListenAndServe(":8000", router)
}

func initLogger(env string) *slog.Logger {
	var logger *slog.Logger

	switch env {
	case envLocal:
		opts := &slog.HandlerOptions{
			Level: slog.LevelDebug,
		}
		slHandler := slog.NewTextHandler(os.Stdout, opts)
		logger = slog.New(slHandler)
	case envDev:
		opts := &slog.HandlerOptions{
			Level: slog.LevelDebug,
		}
		slHandler := slog.NewJSONHandler(os.Stdout, opts)
		logger = slog.New(slHandler)
	}

	return logger
}
