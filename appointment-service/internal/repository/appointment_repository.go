package repository

import (
	"appointment-service/internal/domain"
	"context"
	"database/sql"
)

type AppointmentRepository struct {
	db *sql.DB
}

func NewAppointmentRepository(db *sql.DB) *AppointmentRepository {
	return &AppointmentRepository{db: db}
}

func (r *AppointmentRepository) Create(ctx context.Context, a *domain.Appointment) error {
	err := r.db.QueryRowContext(
		ctx,
		`INSERT INTO appointments (user_id, master_id, service_id, start_time, status)
		 VALUES ($1, $2, $3, $4, $5) RETURNING id`,
		a.UserID, a.MasterID, a.ServiceID, a.StartTime, a.Status,
	).Scan(&a.ID)
	return err
}

func (r *AppointmentRepository) ListByUser(ctx context.Context, userID string) ([]domain.Appointment, error) {
	rows, err := r.db.QueryContext(ctx, "SELECT id, user_id, master_id, service_id, start_time, status FROM appointments WHERE user_id=$1", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var apts []domain.Appointment
	for rows.Next() {
		var a domain.Appointment
		err := rows.Scan(&a.ID, &a.UserID, &a.MasterID, &a.ServiceID, &a.StartTime, &a.Status)
		if err != nil {
			return nil, err
		}
		apts = append(apts, a)
	}
	return apts, nil
}

func (r *AppointmentRepository) Cancel(ctx context.Context, appointmentID string) error {
	_, err := r.db.ExecContext(ctx, "UPDATE appointments SET status='cancelled' WHERE id=$1", appointmentID)
	return err
}
