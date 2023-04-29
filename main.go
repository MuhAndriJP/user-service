package main

import (
	"log"
	"net"
	"os"

	pb "github.com/MuhAndriJP/user-service.git/grpc/user"
	"github.com/MuhAndriJP/user-service.git/repo/mysql"
	"github.com/MuhAndriJP/user-service.git/server"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	mysql.InitDB()
	lis, err := net.Listen("tcp", ":"+os.Getenv("port_grpc"))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	} else {
		log.Println("Connect GRPC to port " + os.Getenv("port_grpc"))
	}

	grpcServer := grpc.NewServer()
	pb.RegisterUserServer(grpcServer, server.NewServerUser())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
