package repository

import (
	"database/sql"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/taqsym/taqsym.uz/task-service/internal/domain"
)

// TaskRepository implements the domain.TaskRepository interface
type TaskRepository struct {
	db *sql.DB
}

// NewTaskRepository creates a new TaskRepository
func NewTaskRepository(db *sql.DB) *TaskRepository {
	return &TaskRepository{
		db: db,
	}
}

// GetByID retrieves a task by its ID
func (r *TaskRepository) GetByID(id string) (*domain.Task, error) {
	query := `
		SELECT id, title, description, status, priority, due_date, assigned_to, created_by, created_at, updated_at
		FROM tasks
		WHERE id = $1
	`

	var task domain.Task
	var dueDate sql.NullTime

	err := r.db.QueryRow(query, id).Scan(
		&task.ID,
		&task.Title,
		&task.Description,
		&task.Status,
		&task.Priority,
		&dueDate,
		&task.AssignedTo,
		&task.CreatedBy,
		&task.CreatedAt,
		&task.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil // Task not found
		}
		return nil, err
	}

	if dueDate.Valid {
		task.DueDate = dueDate.Time
	}

	return &task, nil
}

// GetAll retrieves all tasks
func (r *TaskRepository) GetAll() ([]*domain.Task, error) {
	query := `
		SELECT id, title, description, status, priority, due_date, assigned_to, created_by, created_at, updated_at
		FROM tasks
		ORDER BY created_at DESC
	`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []*domain.Task

	for rows.Next() {
		var task domain.Task
		var dueDate sql.NullTime

		err := rows.Scan(
			&task.ID,
			&task.Title,
			&task.Description,
			&task.Status,
			&task.Priority,
			&dueDate,
			&task.AssignedTo,
			&task.CreatedBy,
			&task.CreatedAt,
			&task.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}

		if dueDate.Valid {
			task.DueDate = dueDate.Time
		}

		tasks = append(tasks, &task)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}

// Create creates a new task
func (r *TaskRepository) Create(task *domain.Task) error {
	query := `
		INSERT INTO tasks (id, title, description, status, priority, due_date, assigned_to, created_by, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
	`

	// Generate a new UUID if not provided
	if task.ID == "" {
		task.ID = uuid.New().String()
	}

	// Set timestamps if not provided
	now := time.Now()
	if task.CreatedAt.IsZero() {
		task.CreatedAt = now
	}
	if task.UpdatedAt.IsZero() {
		task.UpdatedAt = now
	}

	// Set default status if not provided
	if task.Status == "" {
		task.Status = string(domain.TaskStatusTodo)
	}

	var dueDate interface{}
	if !task.DueDate.IsZero() {
		dueDate = task.DueDate
	} else {
		dueDate = nil
	}

	_, err := r.db.Exec(
		query,
		task.ID,
		task.Title,
		task.Description,
		task.Status,
		task.Priority,
		dueDate,
		task.AssignedTo,
		task.CreatedBy,
		task.CreatedAt,
		task.UpdatedAt,
	)

	return err
}

// Update updates an existing task
func (r *TaskRepository) Update(task *domain.Task) error {
	query := `
		UPDATE tasks
		SET title = $1, description = $2, status = $3, priority = $4, due_date = $5, assigned_to = $6, updated_at = $7
		WHERE id = $8
	`

	// Update the updated_at timestamp
	task.UpdatedAt = time.Now()

	var dueDate interface{}
	if !task.DueDate.IsZero() {
		dueDate = task.DueDate
	} else {
		dueDate = nil
	}

	result, err := r.db.Exec(
		query,
		task.Title,
		task.Description,
		task.Status,
		task.Priority,
		dueDate,
		task.AssignedTo,
		task.UpdatedAt,
		task.ID,
	)

	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("task not found")
	}

	return nil
}

// Delete deletes a task by its ID
func (r *TaskRepository) Delete(id string) error {
	query := `
		DELETE FROM tasks
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
		return errors.New("task not found")
	}

	return nil
}
