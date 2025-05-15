package transport

import (
	"context"
	"notification-service/internal/usecase"
	pb "notification-service/notification-service/proto"
)

type Handler struct {
	pb.UnimplementedNotificationServiceServer
	notificationUC *usecase.NotificationUsecase
}

func NewHandler(notificationUC *usecase.NotificationUsecase) *Handler {
	return &Handler{notificationUC: notificationUC}
}

func (h *Handler) ListUserNotifications(ctx context.Context, req *pb.UserIdRequest) (*pb.ListNotificationsResponse, error) {
	notifs, err := h.notificationUC.ListUserNotifications(ctx, req.UserId)
	if err != nil {
		return nil, err
	}
	var resp pb.ListNotificationsResponse
	for _, n := range notifs {
		resp.Notifications = append(resp.Notifications, &pb.NotificationResponse{
			Id:        n.ID,
			Message:   n.Message,
			IsRead:    n.IsRead,
			CreatedAt: n.CreatedAt.Format("2006-01-02 15:04:05"),
		})
	}
	return &resp, nil
}

func (h *Handler) MarkAsRead(ctx context.Context, req *pb.MarkAsReadRequest) (*pb.Empty, error) {
	err := h.notificationUC.MarkAsRead(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.Empty{}, nil
}
