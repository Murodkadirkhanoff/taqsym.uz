package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func NewPostgres() (*sql.DB, error) {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		"postgres", "password", "postgres", "5432", "postgres")

	// username	postgres	Имя пользователя базы данных
	// password	password	Пароль к базе данных
	// host	postgres	Имя хоста или контейнера
	// port	5432	Порт PostgreSQL (по умолчанию)
	// database	postgres	Имя базы данных

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("ошибка открытия БД: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("ошибка подключения к БД: %w", err)
	}

	return db, nil
}
