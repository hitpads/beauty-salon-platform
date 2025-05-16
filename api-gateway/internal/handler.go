package internal

import (
	"context"
	"time"

	appointmentpb "api-gateway/appointment-service/proto"
	notificationpb "api-gateway/notification-service/proto"
	ratingpb "api-gateway/rating-notification-service/proto"
	usermasterpb "api-gateway/user-master-service/proto"

	"google.golang.org/grpc"

	log "github.com/sirupsen/logrus"
)

type Gateway struct {
	UserMasterClient   usermasterpb.UserMasterServiceClient
	AppointmentClient  appointmentpb.ServiceAppointmentServiceClient
	RatingClient       ratingpb.RatingServiceClient
	NotificationClient notificationpb.NotificationServiceClient
}

// CONNECT SERVICES
func NewGateway(userAddr, appointmentAddr, ratingAddr, notificationAddr string) (*Gateway, error) {
	userConn, err := grpc.Dial(userAddr, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	appConn, err := grpc.Dial(appointmentAddr, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	ratingConn, err := grpc.Dial(ratingAddr, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	notifConn, err := grpc.Dial(notificationAddr, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	return &Gateway{
		UserMasterClient:   usermasterpb.NewUserMasterServiceClient(userConn),
		AppointmentClient:  appointmentpb.NewServiceAppointmentServiceClient(appConn),
		RatingClient:       ratingpb.NewRatingServiceClient(ratingConn),
		NotificationClient: notificationpb.NewNotificationServiceClient(notifConn),
	}, nil
}

// user-master-service
func (g *Gateway) RegisterUser(ctx context.Context, req *usermasterpb.RegisterRequest) (*usermasterpb.UserResponse, error) {
	start := time.Now()
	resp, err := g.UserMasterClient.RegisterUser(ctx, req)
	log.WithFields(log.Fields{
		"req":  req,
		"resp": resp,
		"err":  err,
		"took": time.Since(start),
	}).Info("[gateway] RegisterUser")
	return resp, err
}
func (g *Gateway) LoginUser(ctx context.Context, req *usermasterpb.LoginRequest) (*usermasterpb.LoginResponse, error) {
	start := time.Now()
	resp, err := g.UserMasterClient.LoginUser(ctx, req)
	log.WithFields(log.Fields{
		"req":  req,
		"resp": resp,
		"err":  err,
		"took": time.Since(start),
	}).Info("[gateway] LoginUser")
	return resp, err
}
func (g *Gateway) GetUserProfile(ctx context.Context, req *usermasterpb.UserIdRequest) (*usermasterpb.UserResponse, error) {
	start := time.Now()
	resp, err := g.UserMasterClient.GetUserProfile(ctx, req)
	log.WithFields(log.Fields{
		"req":  req,
		"resp": resp,
		"err":  err,
		"took": time.Since(start),
	}).Info("[gateway] GetUserProfile")
	return resp, err
}
func (g *Gateway) ListMasters(ctx context.Context, req *usermasterpb.Empty) (*usermasterpb.ListMastersResponse, error) {
	start := time.Now()
	resp, err := g.UserMasterClient.ListMasters(ctx, req)
	log.WithFields(log.Fields{
		"req":  req,
		"resp": resp,
		"err":  err,
		"took": time.Since(start),
	}).Info("[gateway] ListMasters")
	return resp, err
}
func (g *Gateway) GetMasterByID(ctx context.Context, req *usermasterpb.MasterIdRequest) (*usermasterpb.MasterResponse, error) {
	start := time.Now()
	resp, err := g.UserMasterClient.GetMasterByID(ctx, req)
	log.WithFields(log.Fields{
		"req":  req,
		"resp": resp,
		"err":  err,
		"took": time.Since(start),
	}).Info("[gateway] GetMasterByID")
	return resp, err
}
func (g *Gateway) CreateMaster(ctx context.Context, req *usermasterpb.CreateMasterRequest) (*usermasterpb.MasterResponse, error) {
	start := time.Now()
	resp, err := g.UserMasterClient.CreateMaster(ctx, req)
	log.WithFields(log.Fields{
		"req":  req,
		"resp": resp,
		"err":  err,
		"took": time.Since(start),
	}).Info("[gateway] CreateMaster")
	return resp, err
}

// service-appointment-service
func (g *Gateway) ListServices(ctx context.Context, req *appointmentpb.Empty) (*appointmentpb.ListServicesResponse, error) {
	start := time.Now()
	resp, err := g.AppointmentClient.ListServices(ctx, req)
	log.WithFields(log.Fields{
		"req":  req,
		"resp": resp,
		"err":  err,
		"took": time.Since(start),
	}).Info("[gateway] ListServices")
	return resp, err
}
func (g *Gateway) GetServiceById(ctx context.Context, req *appointmentpb.ServiceIdRequest) (*appointmentpb.ServiceResponse, error) {
	start := time.Now()
	resp, err := g.AppointmentClient.GetServiceById(ctx, req)
	log.WithFields(log.Fields{
		"req":  req,
		"resp": resp,
		"err":  err,
		"took": time.Since(start),
	}).Info("[gateway] GetServiceById")
	return resp, err
}
func (g *Gateway) CreateAppointment(ctx context.Context, req *appointmentpb.CreateAppointmentRequest) (*appointmentpb.AppointmentResponse, error) {
	start := time.Now()
	resp, err := g.AppointmentClient.CreateAppointment(ctx, req)
	log.WithFields(log.Fields{
		"req":  req,
		"resp": resp,
		"err":  err,
		"took": time.Since(start),
	}).Info("[gateway] CreateAppointment")
	return resp, err
}

func (g *Gateway) ListUserAppointments(ctx context.Context, req *appointmentpb.UserIdRequest) (*appointmentpb.ListAppointmentsResponse, error) {
	start := time.Now()
	resp, err := g.AppointmentClient.ListUserAppointments(ctx, req)
	log.WithFields(log.Fields{
		"req":  req,
		"resp": resp,
		"err":  err,
		"took": time.Since(start),
	}).Info("[gateway] ListUserAppointments")
	return resp, err
}

func (g *Gateway) CancelAppointment(ctx context.Context, req *appointmentpb.AppointmentIdRequest) (*appointmentpb.Empty, error) {
	start := time.Now()
	resp, err := g.AppointmentClient.CancelAppointment(ctx, req)
	log.WithFields(log.Fields{
		"req":  req,
		"resp": resp,
		"err":  err,
		"took": time.Since(start),
	}).Info("[gateway] CancelAppointment")
	return resp, err
}

// rating-notification-service
func (g *Gateway) CreateRating(ctx context.Context, req *ratingpb.CreateRatingRequest) (*ratingpb.RatingResponse, error) {
	start := time.Now()
	resp, err := g.RatingClient.CreateRating(ctx, req)
	log.WithFields(log.Fields{
		"req":  req,
		"resp": resp,
		"err":  err,
		"took": time.Since(start),
	}).Info("[gateway] CreateRating")
	return resp, err
}

func (g *Gateway) ListMasterRatings(ctx context.Context, req *ratingpb.MasterIdRequest) (*ratingpb.ListRatingsResponse, error) {
	start := time.Now()
	resp, err := g.RatingClient.ListMasterRatings(ctx, req)
	log.WithFields(log.Fields{
		"req":  req,
		"resp": resp,
		"err":  err,
		"took": time.Since(start),
	}).Info("[gateway] ListMasterRatings")
	return resp, err
}

func (g *Gateway) DeleteRating(ctx context.Context, req *ratingpb.DeleteRatingRequest) (*ratingpb.Empty, error) {
	start := time.Now()
	resp, err := g.RatingClient.DeleteRating(ctx, req)
	log.WithFields(log.Fields{
		"req":  req,
		"resp": resp,
		"err":  err,
		"took": time.Since(start),
	}).Info("[gateway] DeleteRating")
	return resp, err
}

// notification-service

func (g *Gateway) ListUserNotifications(ctx context.Context, req *notificationpb.UserIdRequest) (*notificationpb.ListNotificationsResponse, error) {
	start := time.Now()
	resp, err := g.NotificationClient.ListUserNotifications(ctx, req)
	log.WithFields(log.Fields{
		"req":  req,
		"resp": resp,
		"err":  err,
		"took": time.Since(start),
	}).Info("[gateway] ListUserNotifications")
	return resp, err
}
func (g *Gateway) MarkAsRead(ctx context.Context, req *notificationpb.MarkAsReadRequest) (*notificationpb.Empty, error) {
	start := time.Now()
	resp, err := g.NotificationClient.MarkAsRead(ctx, req)
	log.WithFields(log.Fields{
		"req":  req,
		"resp": resp,
		"err":  err,
		"took": time.Since(start),
	}).Info("[gateway] MarkAsRead")
	return resp, err
}
