package repository

import (
	"context"
	"database/sql"

	"user-master-service/internal/domain"
)

type UserRepository struct {
	db *sql.DB
}

type UserRepositoryInterface interface {
	CreateUser(ctx context.Context, user *domain.User) error
	GetUserByEmail(ctx context.Context, email string) (*domain.User, error)
	GetUserByID(ctx context.Context, id string) (*domain.User, error)
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(ctx context.Context, user *domain.User) error {
	_, err := r.db.ExecContext(ctx,
		"INSERT INTO users (id, name, email, password_hash, role, created_at) VALUES ($1,$2,$3,$4,$5,NOW())",
		user.ID, user.Name, user.Email, user.Password, user.Role)
	return err
}

func (r *UserRepository) GetUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	var user domain.User
	err := r.db.QueryRowContext(ctx, "SELECT id, name, email, password_hash, role, created_at FROM users WHERE email=$1", email).
		Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Role, &user.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) GetUserByID(ctx context.Context, id string) (*domain.User, error) {
	var user domain.User
	err := r.db.QueryRowContext(ctx, "SELECT id, name, email, password_hash, role, created_at FROM users WHERE id=$1", id).
		Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Role, &user.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) UpdateUserRole(ctx context.Context, userID, newRole string) error {
	_, err := r.db.ExecContext(ctx, "UPDATE users SET role=$1 WHERE id=$2", newRole, userID)
	return err
}
