package usecase

import (
	"context"
	"rating-notification-service/internal/domain"
	"rating-notification-service/internal/repository"
	"time"

	"github.com/google/uuid"
)

type RatingUsecase struct {
	repo *repository.RatingRepository
}

func NewRatingUsecase(repo *repository.RatingRepository) *RatingUsecase {
	return &RatingUsecase{repo: repo}
}

func (u *RatingUsecase) CreateRating(ctx context.Context, masterID, userID string, score int, comment string) (*domain.Rating, error) {
	rating := &domain.Rating{
		ID:        uuid.New().String(),
		MasterID:  masterID,
		UserID:    userID,
		Score:     score,
		Comment:   comment,
		CreatedAt: time.Now(),
	}
	err := u.repo.CreateRating(ctx, rating)
	return rating, err
}

func (u *RatingUsecase) ListMasterRatings(ctx context.Context, masterID string) ([]*domain.Rating, error) {
	return u.repo.ListMasterRatings(ctx, masterID)
}

func (u *RatingUsecase) DeleteRating(ctx context.Context, id string) error {
	return u.repo.DeleteRating(ctx, id)
}
