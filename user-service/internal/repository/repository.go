package repository

import (
	"context"
	"database/sql"

	"github.com/Murodkadirkhanoff/taqsym.uz/user-service/internal/domain"
)

type userRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) domain.UserRepository {
	return &userRepo{db: db}
}

func (r *userRepo) Create(ctx context.Context, u *domain.User) error {
	_, err := r.db.ExecContext(ctx,
		"INSERT INTO users (name, email, password) VALUES ($1, $2, $3)",
		u.Name, u.Email, u.Password)
	return err
}

func (r *userRepo) GetByEmail(ctx context.Context, email string) (*domain.User, error) {
	row := r.db.QueryRowContext(ctx,
		"SELECT id, name, email, password FROM users WHERE email=$1", email)

	var u domain.User
	err := row.Scan(&u.ID, &u.Name, &u.Email, &u.Password)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func (r *userRepo) Profile(ctx context.Context, userID int) (*domain.User, error) {
	row := r.db.QueryRowContext(ctx, "SELECT id, name, email FROM users WHERE id =$1", userID)

	var user domain.User
	err := row.Scan(&user.ID, &user.Name, &user.Email)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
