package repository

import (
	"context"
	"database/sql"
	"rating-notification-service/internal/domain"
)

type RatingRepository struct {
	db *sql.DB
}

func NewRatingRepository(db *sql.DB) *RatingRepository {
	return &RatingRepository{db: db}
}

func (r *RatingRepository) CreateRating(ctx context.Context, rating *domain.Rating) error {
	_, err := r.db.ExecContext(ctx,
		"INSERT INTO ratings (id, master_id, user_id, score, comment, created_at) VALUES ($1, $2, $3, $4, $5, $6)",
		rating.ID, rating.MasterID, rating.UserID, rating.Score, rating.Comment, rating.CreatedAt,
	)
	return err
}

func (r *RatingRepository) ListMasterRatings(ctx context.Context, masterID string) ([]*domain.Rating, error) {
	rows, err := r.db.QueryContext(ctx,
		"SELECT id, master_id, user_id, score, comment, created_at FROM ratings WHERE master_id=$1 ORDER BY created_at DESC", masterID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ratings []*domain.Rating
	for rows.Next() {
		var rt domain.Rating
		if err := rows.Scan(&rt.ID, &rt.MasterID, &rt.UserID, &rt.Score, &rt.Comment, &rt.CreatedAt); err != nil {
			return nil, err
		}
		ratings = append(ratings, &rt)
	}
	return ratings, nil
}

func (r *RatingRepository) DeleteRating(ctx context.Context, id string) error {
	_, err := r.db.ExecContext(ctx, "DELETE FROM ratings WHERE id=$1", id)
	return err
}
