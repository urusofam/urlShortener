package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/urusofam/urlShortener/internal/config"
	"github.com/urusofam/urlShortener/internal/log/sl"
	"github.com/urusofam/urlShortener/internal/storage"
	"log/slog"
	"os"
)

func main() {
	cfg := config.LoadConfig("./config/local.yaml")

	logger := SetupLogger(cfg.Env)
	logger.Info("start url-shortener", slog.String("env", cfg.Env))

	dbUrl := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		cfg.DB.User,
		cfg.DB.Password,
		cfg.DB.Host,
		cfg.DB.Port,
		cfg.DB.Name)

	strg, err := storage.NewStorage(dbUrl)
	if err != nil {
		logger.Error("failed to init storage", sl.Err(err))
		os.Exit(1)
	}
	logger.Info(fmt.Sprintf("connected to %s", dbUrl))
	defer strg.Close()

	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)

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
