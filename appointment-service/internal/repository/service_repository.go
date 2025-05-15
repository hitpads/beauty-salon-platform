package repository

import (
	"appointment-service/internal/domain"
	"context"
	"database/sql"
)

type ServiceRepository struct {
	db *sql.DB
}

func NewServiceRepository(db *sql.DB) *ServiceRepository {
	return &ServiceRepository{db: db}
}

func (r *ServiceRepository) List(ctx context.Context) ([]domain.Service, error) {
	rows, err := r.db.QueryContext(ctx, "SELECT id, name, description, price_cents, duration_minutes FROM services")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var services []domain.Service
	for rows.Next() {
		var s domain.Service
		err := rows.Scan(&s.ID, &s.Name, &s.Description, &s.PriceCents, &s.DurationMinutes)
		if err != nil {
			return nil, err
		}
		services = append(services, s)
	}
	return services, nil
}

func (r *ServiceRepository) GetByID(ctx context.Context, id string) (*domain.Service, error) {
	var s domain.Service
	err := r.db.QueryRowContext(ctx, "SELECT id, name, description, price_cents, duration_minutes FROM services WHERE id=$1", id).
		Scan(&s.ID, &s.Name, &s.Description, &s.PriceCents, &s.DurationMinutes)
	if err != nil {
		return nil, err
	}
	return &s, nil
}
