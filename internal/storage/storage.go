package storage

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Storage struct {
	db *pgxpool.Pool
}

var (
	ErrURLNotFound = errors.New("url not found")
	ErrURLExists   = errors.New("url exists")
)

func NewStorage(dbUrl string) (*Storage, error) {
	const op = "storage.NewStorage"

	db, err := pgxpool.New(context.Background(), dbUrl)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	if err := db.Ping(context.Background()); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &Storage{db: db}, nil
}

func (s *Storage) Close() {
	s.db.Close()
}

func (s *Storage) SaveURL(urlToSave, alias string) error {
	const op = "storage.SaveURL"

	query := "insert into urls (url, alias) values ($1, $2)"

	_, err := s.db.Exec(context.Background(), query, urlToSave, alias)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}

func (s *Storage) DeleteURL(alias string) error {
	const op = "storage.DeleteURL"

	query := "delete from urls where alias = $1"

	_, err := s.db.Exec(context.Background(), query, alias)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}

func (s *Storage) UpdateUrlByAlias(urlToSave, alias string) error {
	const op = "storage.UpdateUrlByAlias"

	query := "update urls set url = $1 where alias = $2"

	_, err := s.db.Exec(context.Background(), query, urlToSave, alias)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}

func (s *Storage) GetURL(alias string) (string, error) {
	const op = "storage.GetURL"

	query := "select url from urls where alias = $1"

	url := ""
	err := s.db.QueryRow(context.Background(), query, alias).Scan(&url)
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}
	return url, nil
}
