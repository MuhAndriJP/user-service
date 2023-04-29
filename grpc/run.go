package grpc

import (
	"log"
	"net"
	"os"

	pb "github.com/MuhAndriJP/user-service.git/grpc/user"
	"github.com/MuhAndriJP/user-service.git/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func Run() {
	lis, err := net.Listen("tcp", ":"+os.Getenv("port_grpc"))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	} else {
		log.Println("Connect GRPC " + lis.Addr().String())
	}

	grpcServer := grpc.NewServer()
	pb.RegisterUserServer(grpcServer, server.NewServerUser())
	// Register reflection service on gRPC server.
	reflection.Register(grpcServer)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
