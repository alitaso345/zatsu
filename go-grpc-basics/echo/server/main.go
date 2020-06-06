package main

import (
	"context"
	"log"
	"net"

	pb "github.com/alitaso345/zatsu/go-grpc-basics/echo/proto"
	"google.golang.org/grpc"
)

type echoService struct{}

func (s *echoService) Echo(ctx context.Context, req *pb.EchoRequest) (*pb.EchoResponse, error) {
	return &pb.EchoResponse{Message: "Hi! " + req.GetMessage()}, nil
}

func main() {
	port := ":50051"
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v\n", err)
	}
	srv := grpc.NewServer()
	pb.RegisterEchoServiceServer(srv, &echoService{})
	log.Printf("start server on port: %s\n", port)
	if err := srv.Serve(lis); err != nil {
		log.Printf("failed to serve: %v\n", err)
	}
}
