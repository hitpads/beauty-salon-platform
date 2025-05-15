package usecase

import (
	"appointment-service/internal/domain"
	"appointment-service/internal/repository"
	"context"
)

type ServiceUsecase struct {
	repo *repository.ServiceRepository
}

func NewServiceUsecase(repo *repository.ServiceRepository) *ServiceUsecase {
	return &ServiceUsecase{repo: repo}
}

func (u *ServiceUsecase) List(ctx context.Context) ([]domain.Service, error) {
	return u.repo.List(ctx)
}

func (u *ServiceUsecase) GetByID(ctx context.Context, id string) (*domain.Service, error) {
	return u.repo.GetByID(ctx, id)
}
