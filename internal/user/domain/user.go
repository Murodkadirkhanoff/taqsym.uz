package domain

import (
	"context"
)

type User struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Repository interface {
	GetAll(ctx context.Context) ([]User, error)
}

type Usecase interface {
	GetAll(ctx context.Context) ([]User, error)
}
