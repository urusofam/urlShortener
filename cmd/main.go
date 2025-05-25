package main

import (
	"fmt"
	"github.com/urusofam/urlShortener/internal/config"
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
		logger.Error("failed to init storage", Err(err))
		os.Exit(1)
	}
	logger.Info(fmt.Sprintf("connected to %s", dbUrl))
	defer strg.Close()

	err = strg.SaveURL("https://pkg.go.dev/github.com/jackc/pgx/v5", "pgx")
	if err != nil {
		logger.Error("failed to save url", Err(err))
		os.Exit(1)
	}
	logger.Info("saved url")

	url, err := strg.GetURL("pgx")
	if err != nil {
		logger.Error("failed to get url", Err(err))
		os.Exit(1)
	}
	logger.Info("got url", slog.String("url", url))

	err = strg.UpdateUrlByAlias("https://pkg.go.dev/github.com/jackc/pgx", "pgx")
	if err != nil {
		logger.Error("failed to update url", Err(err))
	}
	logger.Info("updated url")

	url, err = strg.GetURL("pgx")
	if err != nil {
		logger.Error("failed to get url", Err(err))
	}
	logger.Info("got new url", slog.String("url", url))

	err = strg.DeleteURL("pgx")
	if err != nil {
		logger.Error("failed to delete url", Err(err))
		os.Exit(1)
	}
	logger.Info("deleted url")
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

func Err(err error) slog.Attr {
	return slog.Attr{
		Key:   "error",
		Value: slog.StringValue(err.Error()),
	}
}
