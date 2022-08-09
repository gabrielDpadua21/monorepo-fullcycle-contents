package main

import (
	"log"
	"net"

	"github.com/codeedu/fc2-grpc/pb"
	"github.com/codeedu/fc2-grpc/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterCatServiceServer(grpcServer, &services.CatService{})
	reflection.Register(grpcServer)

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatal("Could not serve: %v", err)
	}
}