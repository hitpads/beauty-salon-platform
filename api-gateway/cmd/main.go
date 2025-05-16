package main

import (
	"api-gateway/internal"
	servicesproxy "api-gateway/services-proxy"
	"io"
	"log"
	"net"
	"os"

	appointmentpb "api-gateway/appointment-service/proto"
	notificationpb "api-gateway/notification-service/proto"
	ratingpb "api-gateway/rating-notification-service/proto"
	usermasterpb "api-gateway/user-master-service/proto"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type APIGatewayServer struct {
	internal.Gateway
	usermasterpb.UnimplementedUserMasterServiceServer
	appointmentpb.UnimplementedServiceAppointmentServiceServer
	ratingpb.UnimplementedRatingServiceServer
	notificationpb.UnimplementedNotificationServiceServer
}

func NewAPIGatewayServer(gw *internal.Gateway) *APIGatewayServer {
	return &APIGatewayServer{Gateway: *gw}
}

func main() {

	logFile, err := os.OpenFile("./logs/gateway.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		logrus.Fatal("Could not open log file : ", err)
	}

	mw := io.MultiWriter(os.Stdout, logFile)
	logrus.SetOutput(mw)
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
		ForceColors:   false,
	})

	gw, err := internal.NewGateway(
		"localhost:50052",
		"localhost:50051",
		"localhost:50054",
		"localhost:50055",
	)
	if err != nil {
		log.Fatalf("failed to connect to backends: %v", err)
	}

	lis, err := net.Listen("tcp", ":5000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()

	usermasterpb.RegisterUserMasterServiceServer(grpcServer, &servicesproxy.UserMasterGateway{Gateway: gw})
	appointmentpb.RegisterServiceAppointmentServiceServer(grpcServer, &servicesproxy.AppointmentGateway{Gateway: gw})
	ratingpb.RegisterRatingServiceServer(grpcServer, &servicesproxy.RatingGateway{Gateway: gw})
	notificationpb.RegisterNotificationServiceServer(grpcServer, &servicesproxy.NotificationGateway{Gateway: gw})

	log.Println("API Gateway started on :5000")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
