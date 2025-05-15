package transport

import (
	"context"
	"rating-notification-service/internal/usecase"
	pb "rating-notification-service/rating-notification-service/proto"
)

type Handler struct {
	pb.UnimplementedRatingServiceServer
	ratingUC *usecase.RatingUsecase
}

func NewHandler(ratingUC *usecase.RatingUsecase) *Handler {
	return &Handler{ratingUC: ratingUC}
}

func (h *Handler) CreateRating(ctx context.Context, req *pb.CreateRatingRequest) (*pb.RatingResponse, error) {
	rating, err := h.ratingUC.CreateRating(ctx, req.MasterId, req.UserId, int(req.Score), req.Comment)
	if err != nil {
		return nil, err
	}
	return &pb.RatingResponse{
		Id:        rating.ID,
		MasterId:  rating.MasterID,
		UserId:    rating.UserID,
		Score:     int32(rating.Score),
		Comment:   rating.Comment,
		CreatedAt: rating.CreatedAt.Format("2006-01-02 15:04:05"),
	}, nil
}

func (h *Handler) ListMasterRatings(ctx context.Context, req *pb.MasterIdRequest) (*pb.ListRatingsResponse, error) {
	ratings, err := h.ratingUC.ListMasterRatings(ctx, req.MasterId)
	if err != nil {
		return nil, err
	}
	var pbRatings []*pb.RatingResponse
	for _, r := range ratings {
		pbRatings = append(pbRatings, &pb.RatingResponse{
			Id:        r.ID,
			MasterId:  r.MasterID,
			UserId:    r.UserID,
			Score:     int32(r.Score),
			Comment:   r.Comment,
			CreatedAt: r.CreatedAt.Format("2006-01-02 15:04:05"),
		})
	}
	return &pb.ListRatingsResponse{Ratings: pbRatings}, nil
}

func (h *Handler) DeleteRating(ctx context.Context, req *pb.DeleteRatingRequest) (*pb.Empty, error) {
	err := h.ratingUC.DeleteRating(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.Empty{}, nil
}
