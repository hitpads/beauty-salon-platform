package usecase

import (
	"context"
	"user-master-service/internal/domain"
	"user-master-service/internal/repository"

	"github.com/google/uuid"
)

type MasterUsecase struct {
	repo     *repository.MasterRepository
	userRepo *repository.UserRepository
}

func NewMasterUsecase(repo *repository.MasterRepository, userRepo *repository.UserRepository) *MasterUsecase {
	return &MasterUsecase{repo: repo, userRepo: userRepo}
}

func (u *MasterUsecase) ListMasters(ctx context.Context) ([]*domain.Master, error) {
	return u.repo.ListMasters(ctx)
}

func (u *MasterUsecase) GetMasterByID(ctx context.Context, id string) (*domain.Master, error) {
	return u.repo.GetMasterByID(ctx, id)
}

func (u *MasterUsecase) CreateMaster(ctx context.Context, userID, bio string, experience int) (*domain.Master, error) {
	master := &domain.Master{
		ID:         uuid.New().String(),
		UserID:     userID,
		Bio:        bio,
		Experience: experience,
		Rating:     0,
	}
	err := u.repo.CreateMaster(ctx, master)
	if err != nil {
		return nil, err
	}
	// Меняем роль пользователя на "master"
	if err := u.userRepo.UpdateUserRole(ctx, userID, "master"); err != nil {
		return nil, err
	}
	return master, nil
}

func (u *MasterUsecase) UpdateMaster(ctx context.Context, masterID, bio string, experience int) (*domain.Master, error) {
	master, err := u.repo.GetMasterByID(ctx, masterID)
	if err != nil {
		return nil, err
	}
	master.Bio = bio
	master.Experience = experience
	err = u.repo.UpdateMaster(ctx, master)
	if err != nil {
		return nil, err
	}
	u.repo.InvalidateMasterCache(ctx, masterID)
	return master, nil
}
