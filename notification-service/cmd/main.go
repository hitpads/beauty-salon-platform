package main

import (
	"database/sql"
	"log"
	"net"

	natsconsumer "notification-service/internal/nats"
	"notification-service/internal/repository"
	"notification-service/internal/transport"
	"notification-service/internal/usecase"
	pb "notification-service/notification-service/proto"

	_ "github.com/lib/pq"
	"github.com/nats-io/nats.go"
	"google.golang.org/grpc"
)

func main() {
	dsn := "host=localhost user=postgres password=postgres dbname=bs_bd sslmode=disable"
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}

	notificationRepo := repository.NewNotificationRepository(db)
	notificationUC := usecase.NewNotificationUsecase(notificationRepo)

	// NATS
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal("Cannot connect to NATS:", err)
	}
	go natsconsumer.SubscribeAppointmentCreated(nc, notificationUC)

	// gRPC
	lis, err := net.Listen("tcp", ":50055")
	if err != nil {
		log.Fatal(err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterNotificationServiceServer(grpcServer, transport.NewHandler(notificationUC))

	log.Println("notification-service started on :50055")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
