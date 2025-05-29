package domain

import (
	"time"
)

// Task represents a task entity
type Task struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	Priority    int       `json:"priority"`
	DueDate     time.Time `json:"due_date"`
	AssignedTo  string    `json:"assigned_to"`
	CreatedBy   string    `json:"created_by"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// TaskStatus represents the status of a task
type TaskStatus string

const (
	// TaskStatusTodo represents a task that is not started
	TaskStatusTodo TaskStatus = "TODO"
	// TaskStatusInProgress represents a task that is in progress
	TaskStatusInProgress TaskStatus = "IN_PROGRESS"
	// TaskStatusDone represents a task that is completed
	TaskStatusDone TaskStatus = "DONE"
)

// TaskRepository defines the interface for task data access
type TaskRepository interface {
	GetByID(id string) (*Task, error)
	GetAll() ([]*Task, error)
	Create(task *Task) error
	Update(task *Task) error
	Delete(id string) error
}

// TaskUseCase defines the interface for task business logic
type TaskUseCase interface {
	GetByID(id string) (*Task, error)
	GetAll() ([]*Task, error)
	Create(task *Task) error
	Update(task *Task) error
	Delete(id string) error
}
