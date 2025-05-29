package infrastructure

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// NewPostgresDB creates a new PostgreSQL database connection
func NewPostgresDB(host string, port int, user, password, dbname, sslmode string) (*sql.DB, error) {
	connStr := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		host, port, user, password, dbname, sslmode,
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to open database connection: %w", err)
	}

	// Check if the connection is working
	if err := db.Ping(); err != nil {
		db.Close()
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	// Configure connection pool
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(5)

	return db, nil
}

// CreateTasksTable creates the tasks table if it doesn't exist
func CreateTasksTable(db *sql.DB) error {
	query := `
		CREATE TABLE IF NOT EXISTS tasks (
			id VARCHAR(36) PRIMARY KEY,
			title VARCHAR(255) NOT NULL,
			description TEXT,
			status VARCHAR(20) NOT NULL,
			priority INT NOT NULL DEFAULT 0,
			due_date TIMESTAMP,
			assigned_to VARCHAR(36),
			created_by VARCHAR(36) NOT NULL,
			created_at TIMESTAMP NOT NULL,
			updated_at TIMESTAMP NOT NULL
		)
	`

	_, err := db.Exec(query)
	if err != nil {
		return fmt.Errorf("failed to create tasks table: %w", err)
	}

	return nil
}
