package transport

import (
	pb "appointment-service/appointment-service/proto"
	"appointment-service/internal/domain"
	"appointment-service/internal/usecase"
	"context"
)

type GrpcHandler struct {
	pb.UnimplementedServiceAppointmentServiceServer
	serviceUC     *usecase.ServiceUsecase
	appointmentUC *usecase.AppointmentUsecase
}

func NewGrpcHandler(serviceUC *usecase.ServiceUsecase, appointmentUC *usecase.AppointmentUsecase) *GrpcHandler {
	return &GrpcHandler{serviceUC: serviceUC, appointmentUC: appointmentUC}
}

func (h *GrpcHandler) ListServices(ctx context.Context, _ *pb.Empty) (*pb.ListServicesResponse, error) {
	services, err := h.serviceUC.List(ctx)
	if err != nil {
		return nil, err
	}
	resp := &pb.ListServicesResponse{}
	for _, s := range services {
		resp.Services = append(resp.Services, &pb.ServiceResponse{
			Id:              s.ID,
			Name:            s.Name,
			Description:     s.Description,
			PriceCents:      s.PriceCents,
			DurationMinutes: s.DurationMinutes,
		})
	}
	return resp, nil
}

func (h *GrpcHandler) GetServiceById(ctx context.Context, req *pb.ServiceIdRequest) (*pb.ServiceResponse, error) {
	s, err := h.serviceUC.GetByID(ctx, req.ServiceId)
	if err != nil {
		return nil, err
	}
	return &pb.ServiceResponse{
		Id:              s.ID,
		Name:            s.Name,
		Description:     s.Description,
		PriceCents:      s.PriceCents,
		DurationMinutes: s.DurationMinutes,
	}, nil
}

func (h *GrpcHandler) CreateAppointment(ctx context.Context, req *pb.CreateAppointmentRequest) (*pb.AppointmentResponse, error) {
	a := &domain.Appointment{
		UserID:    req.UserId,
		MasterID:  req.MasterId,
		ServiceID: req.ServiceId,
		StartTime: req.StartTime,
		Status:    "scheduled",
	}
	err := h.appointmentUC.Create(ctx, a)
	if err != nil {
		return nil, err
	}
	return &pb.AppointmentResponse{Id: a.ID, Status: a.Status}, nil
}

func (h *GrpcHandler) ListUserAppointments(ctx context.Context, req *pb.UserIdRequest) (*pb.ListAppointmentsResponse, error) {
	appointments, err := h.appointmentUC.ListByUser(ctx, req.UserId)
	if err != nil {
		return nil, err
	}
	resp := &pb.ListAppointmentsResponse{}
	for _, a := range appointments {
		resp.Appointments = append(resp.Appointments, &pb.AppointmentResponse{
			Id:     a.ID,
			Status: a.Status,
		})
	}
	return resp, nil
}

func (h *GrpcHandler) CancelAppointment(ctx context.Context, req *pb.AppointmentIdRequest) (*pb.Empty, error) {
	err := h.appointmentUC.Cancel(ctx, req.AppointmentId)
	if err != nil {
		return nil, err
	}
	return &pb.Empty{}, nil
}
