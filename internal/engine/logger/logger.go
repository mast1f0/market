package logger

import (
	"log/slog"
	"os"

	"github.com/lmittmann/tint"
)

type Logger struct {
	*slog.Logger
}

func NewLogger() *Logger {
	handler := tint.NewHandler(os.Stdout, &tint.Options{
		Level:      slog.LevelDebug,
		TimeFormat: "15:04:05",
	})

	log := slog.New(handler)

	return &Logger{log}
}
