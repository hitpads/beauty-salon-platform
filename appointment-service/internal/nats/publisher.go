package nats

import (
	"encoding/json"
	"log"
	"time"

	"github.com/nats-io/nats.go"
)

type AppointmentCreatedEvent struct {
	AppointmentID string `json:"appointment_id"`
	UserID        string `json:"user_id"`
	MasterID      string `json:"master_id"`
	ServiceID     string `json:"service_id"`
	StartTime     string `json:"start_time"`
}

func PublishAppointmentCreated(nc *nats.Conn, appointmentID, userID, masterID, serviceID string, startTime time.Time) error {
	event := AppointmentCreatedEvent{
		AppointmentID: appointmentID,
		UserID:        userID,
		MasterID:      masterID,
		ServiceID:     serviceID,
		StartTime:     startTime.Format(time.RFC3339),
	}
	data, err := json.Marshal(event)
	if err != nil {
		return err
	}
	err = nc.Publish("appointment.created", data)
	if err != nil {
		log.Println("NATS publish error:", err)
	}
	return err
}
