package main

import (
	"log/slog"

	"github.com/bugcacher/godenticon-playground/logger"
	"github.com/bugcacher/godenticon-playground/server"
)

func main() {
	err := logger.InitilizeLogger()
	if err != nil {
		slog.Error("failed to initialize logger", "error", err.Error())
		return
	}
	logger.DefaultLogger.Info("starting server...")

	server.Serve()
}
