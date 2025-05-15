package transport

import (
	"context"
	"user-master-service/internal/usecase"
	pb "user-master-service/user-master-service/proto"
)

type Handler struct {
	pb.UnimplementedUserMasterServiceServer
	userUC   *usecase.UserUsecase
	masterUC *usecase.MasterUsecase
}

func NewHandler(userUC *usecase.UserUsecase, masterUC *usecase.MasterUsecase) *Handler {
	return &Handler{userUC: userUC, masterUC: masterUC}
}

func (h *Handler) RegisterUser(ctx context.Context, req *pb.RegisterRequest) (*pb.UserResponse, error) {
	user, err := h.userUC.Register(ctx, req.Name, req.Email, req.Password, "client")
	if err != nil {
		return nil, err
	}
	return &pb.UserResponse{
		Id:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Role:  user.Role,
	}, nil
}

func (h *Handler) LoginUser(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	user, err := h.userUC.Login(ctx, req.Email, req.Password)
	if err != nil {
		return nil, err
	}
	// тут должен быть JWT token, для MVP вернем просто id
	return &pb.LoginResponse{
		Token: user.ID,
		User: &pb.UserResponse{
			Id:    user.ID,
			Name:  user.Name,
			Email: user.Email,
			Role:  user.Role,
		},
	}, nil
}

func (h *Handler) GetUserProfile(ctx context.Context, req *pb.UserIdRequest) (*pb.UserResponse, error) {
	user, err := h.userUC.GetProfile(ctx, req.UserId)
	if err != nil {
		return nil, err
	}
	return &pb.UserResponse{
		Id:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Role:  user.Role,
	}, nil
}

func (h *Handler) ListMasters(ctx context.Context, req *pb.Empty) (*pb.ListMastersResponse, error) {
	masters, err := h.masterUC.ListMasters(ctx)
	if err != nil {
		return nil, err
	}
	resp := &pb.ListMastersResponse{}
	for _, m := range masters {
		resp.Masters = append(resp.Masters, &pb.MasterResponse{
			Id:     m.ID,
			Name:   m.UserID, // You may want to fetch the user's name instead of UserID
			Bio:    m.Bio,
			Rating: m.Rating,
		})
	}
	return resp, nil
}

func (h *Handler) GetMasterByID(ctx context.Context, req *pb.MasterIdRequest) (*pb.MasterResponse, error) {
	master, err := h.masterUC.GetMasterByID(ctx, req.MasterId)
	if err != nil {
		return nil, err
	}
	return &pb.MasterResponse{
		Id:     master.ID,
		Name:   master.UserID, // You may want to fetch the user's name instead of UserID
		Bio:    master.Bio,
		Rating: master.Rating,
	}, nil
}

func (h *Handler) CreateMaster(ctx context.Context, req *pb.CreateMasterRequest) (*pb.MasterResponse, error) {
	master, err := h.masterUC.CreateMaster(ctx, req.UserId, req.Bio, int(req.Experience))
	if err != nil {
		return nil, err
	}
	return &pb.MasterResponse{
		Id:     master.ID,
		Name:   "", // можно запросить у User по user_id, если нужно
		Bio:    master.Bio,
		Rating: master.Rating,
	}, nil
}

func (h *Handler) UpdateMaster(ctx context.Context, req *pb.UpdateMasterRequest) (*pb.MasterResponse, error) {
	master, err := h.masterUC.UpdateMaster(ctx, req.MasterId, req.Bio, int(req.Experience))
	if err != nil {
		return nil, err
	}
	return &pb.MasterResponse{
		Id:     master.ID,
		Name:   "",
		Bio:    master.Bio,
		Rating: master.Rating,
	}, nil
}
