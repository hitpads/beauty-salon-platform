package servicesproxy

import (
	"api-gateway/internal"
	usermasterpb "api-gateway/user-master-service/proto"
	"context"
)

type UserMasterGateway struct {
	usermasterpb.UnimplementedUserMasterServiceServer
	Gateway *internal.Gateway
}

func (s *UserMasterGateway) RegisterUser(ctx context.Context, req *usermasterpb.RegisterRequest) (*usermasterpb.UserResponse, error) {
	return s.Gateway.RegisterUser(ctx, req)
}

func (s *UserMasterGateway) LoginUser(ctx context.Context, req *usermasterpb.LoginRequest) (*usermasterpb.LoginResponse, error) {
	return s.Gateway.LoginUser(ctx, req)
}

func (s *UserMasterGateway) GetUserProfile(ctx context.Context, req *usermasterpb.UserIdRequest) (*usermasterpb.UserResponse, error) {
	return s.Gateway.GetUserProfile(ctx, req)
}

func (s *UserMasterGateway) ListMasters(ctx context.Context, req *usermasterpb.Empty) (*usermasterpb.ListMastersResponse, error) {
	return s.Gateway.ListMasters(ctx, req)
}

func (s *UserMasterGateway) GetMasterByID(ctx context.Context, req *usermasterpb.MasterIdRequest) (*usermasterpb.MasterResponse, error) {
	return s.Gateway.GetMasterByID(ctx, req)
}

func (s *UserMasterGateway) CreateMaster(ctx context.Context, req *usermasterpb.CreateMasterRequest) (*usermasterpb.MasterResponse, error) {
	return s.Gateway.CreateMaster(ctx, req)
}
