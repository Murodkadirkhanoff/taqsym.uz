package usecase

import (
	"context"

	"github.com/Murodkadirkhanoff/taqsym.uz/task-service/internal/domain"
)

type taskUsecase struct {
	repo domain.Repository
}

func NewTaskUsecase(repo domain.Repository) domain.Usecase {
	return &taskUsecase{repo: repo}
}

func (uc *taskUsecase) Create(ctx context.Context, task *domain.Task) error {
	err := uc.repo.Create(ctx, task)
	return err
}

func (uc *taskUsecase) List(ctx context.Context) ([]*domain.Task, error) {
	return uc.repo.List(ctx)
}
