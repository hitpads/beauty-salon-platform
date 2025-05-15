package servicesproxy

import (
	appointmentpb "api-gateway/appointment-service/proto"
	"api-gateway/internal"
	"context"
)

type AppointmentGateway struct {
	appointmentpb.UnimplementedServiceAppointmentServiceServer
	Gateway *internal.Gateway
}

func (s *AppointmentGateway) ListServices(ctx context.Context, req *appointmentpb.Empty) (*appointmentpb.ListServicesResponse, error) {
	return s.Gateway.ListServices(ctx, req)
}

func (s *AppointmentGateway) GetServiceById(ctx context.Context, req *appointmentpb.ServiceIdRequest) (*appointmentpb.ServiceResponse, error) {
	return s.Gateway.GetServiceById(ctx, req)
}

func (s *AppointmentGateway) CreateAppointment(ctx context.Context, req *appointmentpb.CreateAppointmentRequest) (*appointmentpb.AppointmentResponse, error) {
	return s.Gateway.CreateAppointment(ctx, req)
}

func (s *AppointmentGateway) ListUserAppointments(ctx context.Context, req *appointmentpb.UserIdRequest) (*appointmentpb.ListAppointmentsResponse, error) {
	return s.Gateway.ListUserAppointments(ctx, req)
}

func (s *AppointmentGateway) CancelAppointment(ctx context.Context, req *appointmentpb.AppointmentIdRequest) (*appointmentpb.Empty, error) {
	return s.Gateway.CancelAppointment(ctx, req)
}
