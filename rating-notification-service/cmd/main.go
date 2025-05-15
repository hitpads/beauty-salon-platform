package main

import (
	"database/sql"
	"log"
	"net"

	"rating-notification-service/internal/repository"
	"rating-notification-service/internal/transport"
	"rating-notification-service/internal/usecase"
	pb "rating-notification-service/rating-notification-service/proto"

	_ "github.com/lib/pq"
	"google.golang.org/grpc"
)

func main() {
	dsn := "host=localhost user=postgres password=postgres dbname=bs_bd sslmode=disable"
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}

	ratingRepo := repository.NewRatingRepository(db)
	ratingUC := usecase.NewRatingUsecase(ratingRepo)

	lis, err := net.Listen("tcp", ":50054")
	if err != nil {
		log.Fatal(err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterRatingServiceServer(grpcServer, transport.NewHandler(ratingUC))

	log.Println("rating-service started on :50054")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
