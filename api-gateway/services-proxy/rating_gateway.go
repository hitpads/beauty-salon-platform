package servicesproxy

import (
	"api-gateway/internal"
	ratingpb "api-gateway/rating-notification-service/proto"
	"context"
)

type RatingGateway struct {
	ratingpb.UnimplementedRatingServiceServer
	Gateway *internal.Gateway
}

func (s *RatingGateway) CreateRating(ctx context.Context, req *ratingpb.CreateRatingRequest) (*ratingpb.RatingResponse, error) {
	return s.Gateway.CreateRating(ctx, req)
}

func (s *RatingGateway) ListMasterRatings(ctx context.Context, req *ratingpb.MasterIdRequest) (*ratingpb.ListRatingsResponse, error) {
	return s.Gateway.ListMasterRatings(ctx, req)
}

func (s *RatingGateway) DeleteRating(ctx context.Context, req *ratingpb.DeleteRatingRequest) (*ratingpb.Empty, error) {
	return s.Gateway.DeleteRating(ctx, req)
}
