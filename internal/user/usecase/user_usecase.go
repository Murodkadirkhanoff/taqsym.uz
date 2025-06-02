package usecase

import (
	"context"

	"github.com/Murodkadirkhanoff/taqsym.uz/internal/user/domain"
)

type UserUsecase struct {
	repo domain.Repository
}

func NewUserUsecase(repo domain.Repository) domain.Usecase {
	return &UserUsecase{repo: repo}
}

func (uc *UserUsecase) GetAll(ctx context.Context) ([]domain.User, error) {
	return uc.repo.GetAll(ctx)
}
