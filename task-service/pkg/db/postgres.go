package db

import (
	"database/sql"
	"fmt"
)

func NewPostgres() (*sql.DB, error) {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		"postgres", "password", "task-db", "5434", "postgres")

}
