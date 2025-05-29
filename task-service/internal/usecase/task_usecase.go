package usecase

import (
	"errors"
	"time"

	"github.com/taqsym/taqsym.uz/task-service/internal/domain"
)

// TaskUseCase implements the domain.TaskUseCase interface
type TaskUseCase struct {
	repo domain.TaskRepository
}

// NewTaskUseCase creates a new TaskUseCase
func NewTaskUseCase(repo domain.TaskRepository) *TaskUseCase {
	return &TaskUseCase{
		repo: repo,
	}
}

// GetByID retrieves a task by its ID
func (uc *TaskUseCase) GetByID(id string) (*domain.Task, error) {
	if id == "" {
		return nil, errors.New("task ID cannot be empty")
	}

	task, err := uc.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	if task == nil {
		return nil, errors.New("task not found")
	}

	return task, nil
}

// GetAll retrieves all tasks
func (uc *TaskUseCase) GetAll() ([]*domain.Task, error) {
	return uc.repo.GetAll()
}

// Create creates a new task
func (uc *TaskUseCase) Create(task *domain.Task) error {
	if task == nil {
		return errors.New("task cannot be nil")
	}

	if task.Title == "" {
		return errors.New("task title cannot be empty")
	}

	if task.CreatedBy == "" {
		return errors.New("created_by cannot be empty")
	}

	// Set default values
	if task.Status == "" {
		task.Status = string(domain.TaskStatusTodo)
	}

	// Set timestamps
	now := time.Now()
	task.CreatedAt = now
	task.UpdatedAt = now

	return uc.repo.Create(task)
}

// Update updates an existing task
func (uc *TaskUseCase) Update(task *domain.Task) error {
	if task == nil {
		return errors.New("task cannot be nil")
	}

	if task.ID == "" {
		return errors.New("task ID cannot be empty")
	}

	if task.Title == "" {
		return errors.New("task title cannot be empty")
	}

	// Check if the task exists
	existingTask, err := uc.repo.GetByID(task.ID)
	if err != nil {
		return err
	}

	if existingTask == nil {
		return errors.New("task not found")
	}

	// Preserve created_by and created_at from the existing task
	task.CreatedBy = existingTask.CreatedBy
	task.CreatedAt = existingTask.CreatedAt

	// Update the updated_at timestamp
	task.UpdatedAt = time.Now()

	return uc.repo.Update(task)
}

// Delete deletes a task by its ID
func (uc *TaskUseCase) Delete(id string) error {
	if id == "" {
		return errors.New("task ID cannot be empty")
	}

	return uc.repo.Delete(id)
}
