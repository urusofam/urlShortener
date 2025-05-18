package main

import (
	"github.com/urusofam/urlShortener/internal/config"
	"log/slog"
	"os"
)

func main() {
	cfg := config.LoadConfig("./config/local.yaml")

	logger := SetupLogger(cfg.Env)
	logger.Info("starting url-shortener", slog.String("env", cfg.Env))
}

func SetupLogger(env string) *slog.Logger {
	var logger *slog.Logger
	switch env {
	case "local":
		logger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case "dev":
		logger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case "prod":
		logger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}
	return logger
}
