package domain

import "context"

type Task struct {
	ID          int64  `json:"id"`
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	UserID      int64  `json:"user_id" binding:"required"`
}

type Repository interface {
	Create(ctx context.Context, task *Task) error
	List(ctx context.Context) ([]*Task, error)
}

type Usecase interface {
	Create(ctx context.Context, task *Task) error
	List(ctx context.Context) ([]*Task, error)
}
