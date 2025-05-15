package natsconsumer

import (
	"context"
	"encoding/json"
	"log"
	"notification-service/internal/usecase"

	"github.com/nats-io/nats.go"
)

type AppointmentCreatedEvent struct {
	AppointmentID string `json:"appointment_id"`
	UserID        string `json:"user_id"`
	MasterID      string `json:"master_id"`
	ServiceID     string `json:"service_id"`
	StartTime     string `json:"start_time"`
}

func SubscribeAppointmentCreated(nc *nats.Conn, notificationUC *usecase.NotificationUsecase) {
	_, err := nc.Subscribe("appointment.created", func(msg *nats.Msg) {
		var event AppointmentCreatedEvent
		if err := json.Unmarshal(msg.Data, &event); err != nil {
			log.Println("Error decoding appointment.created event:", err)
			return
		}
		message := "Ваша бронь успешно создана на " + event.StartTime
		err := notificationUC.CreateNotification(context.Background(), event.UserID, message)
		if err != nil {
			log.Println("Error saving notification:", err)
		} else {
			log.Printf("Notification sent for appointment %s (user %s)", event.AppointmentID, event.UserID)
		}
	})
	if err != nil {
		log.Fatal("Failed to subscribe to appointment.created:", err)
	}
}
