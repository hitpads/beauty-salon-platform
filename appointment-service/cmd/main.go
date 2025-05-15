package main

import (
	pb "appointment-service/appointment-service/proto"
	"appointment-service/internal/repository"
	"appointment-service/internal/transport"
	"appointment-service/internal/usecase"
	"database/sql"
	"log"
	"net"

	_ "github.com/lib/pq"
	"github.com/nats-io/nats.go"
	"google.golang.org/grpc"
)

func main() {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal("Cannot connect to NATS:", err)
	}
	defer nc.Close()

	dsn := "host=localhost user=postgres password=postgres dbname=bs_bd sslmode=disable"
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}

	serviceRepo := repository.NewServiceRepository(db)
	appointmentRepo := repository.NewAppointmentRepository(db)
	serviceUC := usecase.NewServiceUsecase(serviceRepo)
	appointmentUC := usecase.NewAppointmentUsecase(appointmentRepo, nc)
	handler := transport.NewGrpcHandler(serviceUC, appointmentUC)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterServiceAppointmentServiceServer(s, handler)

	log.Println("gRPC server running at :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
