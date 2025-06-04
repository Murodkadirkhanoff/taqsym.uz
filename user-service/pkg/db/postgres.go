package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func NewPostgres() (*sql.DB, error) {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		"postgres", "password", "postgres", "5432", "postgres")

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("ошибка открытия БД: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("ошибка подключения к БД: %w", err)
	}

	return db, nil
}
