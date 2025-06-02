package domain

import (
	"context"
)

type User struct {
	ID       int64
	Name     string
	Email    string
	Password string
}

type UserRepository interface {
	Create(ctx context.Context, user *User) error
	GetByEmail(ctx context.Context, email string) (*User, error)
}

type UserUseCase interface {
	Register(ctx context.Context, user *User) error
	Login(ctx context.Context, email, password string) (string, error)
}
