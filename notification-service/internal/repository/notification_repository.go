package repository

import (
	"context"
	"database/sql"
	"notification-service/internal/domain"
)

type NotificationRepository struct {
	db *sql.DB
}

func NewNotificationRepository(db *sql.DB) *NotificationRepository {
	return &NotificationRepository{db: db}
}

func (r *NotificationRepository) CreateNotification(ctx context.Context, notif *domain.Notification) error {
	_, err := r.db.ExecContext(ctx,
		"INSERT INTO notifications (id, user_id, message, is_read, created_at) VALUES ($1, $2, $3, $4, $5)",
		notif.ID, notif.UserID, notif.Message, notif.IsRead, notif.CreatedAt)
	return err
}

func (r *NotificationRepository) ListUserNotifications(ctx context.Context, userID string) ([]*domain.Notification, error) {
	rows, err := r.db.QueryContext(ctx,
		"SELECT id, message, is_read, created_at FROM notifications WHERE user_id=$1 ORDER BY created_at DESC", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var notifs []*domain.Notification
	for rows.Next() {
		var n domain.Notification
		n.UserID = userID
		if err := rows.Scan(&n.ID, &n.Message, &n.IsRead, &n.CreatedAt); err != nil {
			return nil, err
		}
		notifs = append(notifs, &n)
	}
	return notifs, nil
}

func (r *NotificationRepository) MarkAsRead(ctx context.Context, id string) error {
	_, err := r.db.ExecContext(ctx, "UPDATE notifications SET is_read=TRUE WHERE id=$1", id)
	return err
}
