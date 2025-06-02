package usecase

import (
	"context"
	"errors"

	"github.com/Murodkadirkhanoff/taqsym.uz/user-service/internal/domain"
	"github.com/Murodkadirkhanoff/taqsym.uz/user-service/pkg/utils"
)

type userUC struct {
	repo domain.UserRepository
}

func NewUserUseCase(repo domain.UserRepository) domain.UserUseCase {
	return &userUC{repo: repo}
}

func (uc *userUC) Register(ctx context.Context, u *domain.User) error {
	hashed, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}
	u.Password = hashed
	return uc.repo.Create(ctx, u)
}

func (uc *userUC) Login(ctx context.Context, email, password string) (string, error) {
	u, err := uc.repo.GetByEmail(ctx, email)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	if !utils.CheckPasswordHash(password, u.Password) {
		return "", errors.New("invalid credentials")
	}

	return utils.GenerateToken(u.ID, u.Email)
}
