package domain

import (
	"time"
)

// User represents a user entity
type User struct {
	ID        string    `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"-"` // Password is not included in JSON responses
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Role      string    `json:"role"`
	Active    bool      `json:"active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// UserRole represents the role of a user
type UserRole string

const (
	// UserRoleAdmin represents an admin user
	UserRoleAdmin UserRole = "ADMIN"
	// UserRoleUser represents a regular user
	UserRoleUser UserRole = "USER"
)

// UserRepository defines the interface for user data access
type UserRepository interface {
	GetByID(id string) (*User, error)
	GetByUsername(username string) (*User, error)
	GetByEmail(email string) (*User, error)
	GetAll() ([]*User, error)
	Create(user *User) error
	Update(user *User) error
	Delete(id string) error
}

// UserUseCase defines the interface for user business logic
type UserUseCase interface {
	GetByID(id string) (*User, error)
	GetByUsername(username string) (*User, error)
	GetByEmail(email string) (*User, error)
	GetAll() ([]*User, error)
	Create(user *User) error
	Update(user *User) error
	Delete(id string) error
	Authenticate(username, password string) (*User, error)
}
