package repository

import (
	"database/sql"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/taqsym/taqsym.uz/user-service/internal/domain"
	"golang.org/x/crypto/bcrypt"
)

// UserRepository implements the domain.UserRepository interface
type UserRepository struct {
	db *sql.DB
}

// NewUserRepository creates a new UserRepository
func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

// GetByID retrieves a user by its ID
func (r *UserRepository) GetByID(id string) (*domain.User, error) {
	query := `
		SELECT id, username, email, password, first_name, last_name, role, active, created_at, updated_at
		FROM users
		WHERE id = $1
	`

	var user domain.User
	err := r.db.QueryRow(query, id).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.Password,
		&user.FirstName,
		&user.LastName,
		&user.Role,
		&user.Active,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil // User not found
		}
		return nil, err
	}

	return &user, nil
}

// GetByUsername retrieves a user by its username
func (r *UserRepository) GetByUsername(username string) (*domain.User, error) {
	query := `
		SELECT id, username, email, password, first_name, last_name, role, active, created_at, updated_at
		FROM users
		WHERE username = $1
	`

	var user domain.User
	err := r.db.QueryRow(query, username).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.Password,
		&user.FirstName,
		&user.LastName,
		&user.Role,
		&user.Active,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil // User not found
		}
		return nil, err
	}

	return &user, nil
}

// GetByEmail retrieves a user by its email
func (r *UserRepository) GetByEmail(email string) (*domain.User, error) {
	query := `
		SELECT id, username, email, password, first_name, last_name, role, active, created_at, updated_at
		FROM users
		WHERE email = $1
	`

	var user domain.User
	err := r.db.QueryRow(query, email).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.Password,
		&user.FirstName,
		&user.LastName,
		&user.Role,
		&user.Active,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil // User not found
		}
		return nil, err
	}

	return &user, nil
}

// GetAll retrieves all users
func (r *UserRepository) GetAll() ([]*domain.User, error) {
	query := `
		SELECT id, username, email, password, first_name, last_name, role, active, created_at, updated_at
		FROM users
		ORDER BY created_at DESC
	`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*domain.User

	for rows.Next() {
		var user domain.User
		err := rows.Scan(
			&user.ID,
			&user.Username,
			&user.Email,
			&user.Password,
			&user.FirstName,
			&user.LastName,
			&user.Role,
			&user.Active,
			&user.CreatedAt,
			&user.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}

		users = append(users, &user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

// Create creates a new user
func (r *UserRepository) Create(user *domain.User) error {
	query := `
		INSERT INTO users (id, username, email, password, first_name, last_name, role, active, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
	`

	// Generate a new UUID if not provided
	if user.ID == "" {
		user.ID = uuid.New().String()
	}

	// Set timestamps if not provided
	now := time.Now()
	if user.CreatedAt.IsZero() {
		user.CreatedAt = now
	}
	if user.UpdatedAt.IsZero() {
		user.UpdatedAt = now
	}

	// Set default role if not provided
	if user.Role == "" {
		user.Role = string(domain.UserRoleUser)
	}

	// Set active status if not provided
	if !user.Active {
		user.Active = true
	}

	// Hash the password if it's not already hashed
	if !isPasswordHashed(user.Password) {
		hashedPassword, err := hashPassword(user.Password)
		if err != nil {
			return err
		}
		user.Password = hashedPassword
	}

	_, err := r.db.Exec(
		query,
		user.ID,
		user.Username,
		user.Email,
		user.Password,
		user.FirstName,
		user.LastName,
		user.Role,
		user.Active,
		user.CreatedAt,
		user.UpdatedAt,
	)

	return err
}

// Update updates an existing user
func (r *UserRepository) Update(user *domain.User) error {
	query := `
		UPDATE users
		SET username = $1, email = $2, first_name = $3, last_name = $4, role = $5, active = $6, updated_at = $7
		WHERE id = $8
	`

	// Update the updated_at timestamp
	user.UpdatedAt = time.Now()

	result, err := r.db.Exec(
		query,
		user.Username,
		user.Email,
		user.FirstName,
		user.LastName,
		user.Role,
		user.Active,
		user.UpdatedAt,
		user.ID,
	)

	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("user not found")
	}

	return nil
}

// UpdatePassword updates a user's password
func (r *UserRepository) UpdatePassword(id, password string) error {
	query := `
		UPDATE users
		SET password = $1, updated_at = $2
		WHERE id = $3
	`

	// Hash the password
	hashedPassword, err := hashPassword(password)
	if err != nil {
		return err
	}

	result, err := r.db.Exec(
		query,
		hashedPassword,
		time.Now(),
		id,
	)

	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("user not found")
	}

	return nil
}

// Delete deletes a user by its ID
func (r *UserRepository) Delete(id string) error {
	query := `
		DELETE FROM users
		WHERE id = $1
	`

	result, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("user not found")
	}

	return nil
}

// Helper functions

// hashPassword hashes a password using bcrypt
func hashPassword(password string) (string, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedBytes), nil
}

// isPasswordHashed checks if a password is already hashed
func isPasswordHashed(password string) bool {
	// A bcrypt hash starts with $2a$, $2b$, or $2y$
	return len(password) > 4 && (password[:4] == "$2a$" || password[:4] == "$2b$" || password[:4] == "$2y$")
}
