package usecase

import (
	"appointment-service/internal/domain"
	"appointment-service/internal/repository"
	"context"
	"encoding/json"

	"github.com/nats-io/nats.go"
)

type AppointmentUsecase struct {
	repo *repository.AppointmentRepository
	nc   *nats.Conn // NATS connection for publishing
}

func NewAppointmentUsecase(repo *repository.AppointmentRepository, nc *nats.Conn) *AppointmentUsecase {
	return &AppointmentUsecase{repo: repo, nc: nc}
}

func (u *AppointmentUsecase) Create(ctx context.Context, a *domain.Appointment) error {
	a.Status = "scheduled"
	err := u.repo.Create(ctx, a)
	if err != nil {
		return err
	}

	// Publish event to NATS
	event := map[string]interface{}{
		"appointment_id": a.ID,
		"user_id":        a.UserID,
		"master_id":      a.MasterID,
		"service_id":     a.ServiceID,
		"start_time":     a.StartTime,
		"status":         a.Status,
	}
	payload, err := json.Marshal(event)
	if err != nil {
		return err
	}
	if u.nc != nil {
		u.nc.Publish("appointment.created", payload)
	}

	return nil
}

func (u *AppointmentUsecase) ListByUser(ctx context.Context, userID string) ([]domain.Appointment, error) {
	return u.repo.ListByUser(ctx, userID)
}

func (u *AppointmentUsecase) Cancel(ctx context.Context, appointmentID string) error {
	return u.repo.Cancel(ctx, appointmentID)
}
