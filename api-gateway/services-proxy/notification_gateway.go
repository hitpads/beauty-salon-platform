package servicesproxy

import (
	"api-gateway/internal"
	notificationpb "api-gateway/notification-service/proto"
	"context"
)

type NotificationGateway struct {
	notificationpb.UnimplementedNotificationServiceServer
	Gateway *internal.Gateway
}

func (s *NotificationGateway) ListUserNotifications(ctx context.Context, req *notificationpb.UserIdRequest) (*notificationpb.ListNotificationsResponse, error) {
	return s.Gateway.ListUserNotifications(ctx, req)
}

func (s *NotificationGateway) MarkAsRead(ctx context.Context, req *notificationpb.MarkAsReadRequest) (*notificationpb.Empty, error) {
	return s.Gateway.MarkAsRead(ctx, req)
}
