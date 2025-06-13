package repository

import (
	"context"
	"database/sql"

	"github.com/Murodkadirkhanoff/taqsym.uz/task-service/internal/domain"
)

type TaskRepository struct {
	db *sql.DB
}

func NewTaskRepository(db *sql.DB) domain.Repository {
	return &TaskRepository{db: db}
}

func (r *TaskRepository) Create(ctx context.Context, task *domain.Task) error {
	err := r.db.QueryRowContext(ctx,
		"INSERT INTO tasks (title, description, user_id) VALUES ($1, $2, $3) RETURNING ID",
		task.Title, task.Description, task.UserID).Scan(task.ID)
	return err
}

func (r *TaskRepository) List(ctx context.Context) ([]*domain.Task, error) {
	rows, err := r.db.QueryContext(ctx, "SELECT id, title, description, user_id FROM tasks")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []*domain.Task
	for rows.Next() {
		var task domain.Task
		err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.UserID)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, &task)
	}
	return tasks, nil
}
