package main

import (
	"log"
	"net"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"

	contentpb "content-service/src/protos"
)

func loadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found or error loading .env file:", err)
	}
}

func main() {
	loadEnv()
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	contentpb.RegisterContentServiceServer(grpcServer, &contentServer{})
	log.Println("gRPC server listening on :50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
