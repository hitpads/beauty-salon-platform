// internal/repository/master_repository.go
package repository

import (
	"context"
	"database/sql"
	"encoding/json"
	"time"

	"user-master-service/internal/domain"

	"github.com/redis/go-redis/v9"
)

type MasterRepository struct {
	db    *sql.DB
	cache *redis.Client
}

func NewMasterRepository(db *sql.DB, cache *redis.Client) *MasterRepository {
	return &MasterRepository{db: db, cache: cache}
}

func (r *MasterRepository) ListMasters(ctx context.Context) ([]*domain.Master, error) {
	rows, err := r.db.QueryContext(ctx, "SELECT id, user_id, bio, experience, rating FROM masters")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var masters []*domain.Master
	for rows.Next() {
		var m domain.Master
		if err := rows.Scan(&m.ID, &m.UserID, &m.Bio, &m.Experience, &m.Rating); err != nil {
			return nil, err
		}
		masters = append(masters, &m)
	}
	return masters, nil
}

func (r *MasterRepository) GetMasterByID(ctx context.Context, id string) (*domain.Master, error) {
	cacheKey := "master:" + id

	// try cache first
	cached, err := r.cache.Get(ctx, cacheKey).Result()
	if err == nil && cached != "" {
		var m domain.Master
		if json.Unmarshal([]byte(cached), &m) == nil {
			return &m, nil
		}
	}

	// if not found in cache, query the database
	var m domain.Master
	err = r.db.QueryRowContext(ctx, "SELECT id, user_id, bio, experience, rating FROM masters WHERE id=$1", id).
		Scan(&m.ID, &m.UserID, &m.Bio, &m.Experience, &m.Rating)
	if err != nil {
		return nil, err
	}

	// put the result into cache for 5 minutes
	bytes, _ := json.Marshal(m)
	_ = r.cache.Set(ctx, cacheKey, bytes, 5*time.Minute).Err()

	return &m, nil
}

// InvalidateMasterCache удаляет кеш для мастера по id
func (r *MasterRepository) InvalidateMasterCache(ctx context.Context, id string) {
	r.cache.Del(ctx, "master:"+id)
}

func (r *MasterRepository) CreateMaster(ctx context.Context, m *domain.Master) error {
	_, err := r.db.ExecContext(ctx,
		"INSERT INTO masters (id, user_id, bio, experience, rating) VALUES ($1, $2, $3, $4, $5)",
		m.ID, m.UserID, m.Bio, m.Experience, m.Rating,
	)
	return err
}

func (r *MasterRepository) UpdateMaster(ctx context.Context, m *domain.Master) error {
	_, err := r.db.ExecContext(ctx,
		"UPDATE masters SET bio=$1, experience=$2 WHERE id=$3",
		m.Bio, m.Experience, m.ID,
	)
	return err
}
