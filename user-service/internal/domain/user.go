package domain

import (
	"context"
)

type User struct {
	ID       int64  `json:"id"`
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type UserRepository interface {
	Create(ctx context.Context, user *User) error
	GetByEmail(ctx context.Context, email string) (*User, error)
	Profile(ctx context.Context, userID int) (*User, error)
}

type UserUseCase interface {
	// Register(ctx context.Context, user *User) error
	Login(ctx context.Context, email, password string) (string, error)
	// Profile(ctx context.Context, userID int) (*User, error)
}
