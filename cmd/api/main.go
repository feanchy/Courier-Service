package main

import (
	"log"
	"net"

	pb "github.com/feanchy/Courier-Service/gen/order"
	courier "github.com/feanchy/Courier-Service/internal/courier"
	"google.golang.org/grpc"
)

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterCourierServiceServer(grpcServer, courier.NewServer())

	log.Println("Courier Service gRPC listening on :50051")

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("gRPC server failed: %v", err)
	}
}
