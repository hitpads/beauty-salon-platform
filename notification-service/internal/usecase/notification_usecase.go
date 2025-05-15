package usecase

import (
	"context"
	"notification-service/internal/domain"
	"notification-service/internal/repository"
	"time"

	"github.com/google/uuid"
)

type NotificationUsecase struct {
	repo *repository.NotificationRepository
}

func NewNotificationUsecase(repo *repository.NotificationRepository) *NotificationUsecase {
	return &NotificationUsecase{repo: repo}
}

func (u *NotificationUsecase) CreateNotification(ctx context.Context, userID, message string) error {
	notif := &domain.Notification{
		ID:        uuid.New().String(),
		UserID:    userID,
		Message:   message,
		IsRead:    false,
		CreatedAt: time.Now(),
	}
	return u.repo.CreateNotification(ctx, notif)
}

func (u *NotificationUsecase) ListUserNotifications(ctx context.Context, userID string) ([]*domain.Notification, error) {
	return u.repo.ListUserNotifications(ctx, userID)
}

func (u *NotificationUsecase) MarkAsRead(ctx context.Context, id string) error {
	return u.repo.MarkAsRead(ctx, id)
}
