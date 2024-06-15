package logger

import (
	"log/slog"
	"os"
)

var (
	DefaultLogger *slog.Logger
)

func InitilizeLogger() error {
	DefaultLogger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}))
	return nil
}
