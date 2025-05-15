package main

import (
	"database/sql"
	"log"
	"net"
	_ "os"

	"user-master-service/internal/repository"
	"user-master-service/internal/transport"
	"user-master-service/internal/usecase"
	pb "user-master-service/user-master-service/proto"

	_ "github.com/lib/pq"
	"github.com/redis/go-redis/v9"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	dsn := "host=localhost user=postgres password=postgres dbname=bs_bd sslmode=disable"
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		DB:   0,
	})

	userRepo := repository.NewUserRepository(db)
	masterRepo := repository.NewMasterRepository(db, rdb)
	userUC := usecase.NewUserUsecase(userRepo)
	masterUC := usecase.NewMasterUsecase(masterRepo, userRepo)

	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatal(err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterUserMasterServiceServer(grpcServer, transport.NewHandler(userUC, masterUC))

	reflection.Register(grpcServer)

	log.Println("user-master-service started on :50052")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
