package main

import (
	"io"
	"log"
	"net"
	"sync"

	pb "github.com/alitaso345/zatsu/go-grpc-basics/chat/proto"
	"google.golang.org/grpc"
)

type chatService struct{}

var streams sync.Map

func (s *chatService) Connect(stream pb.ChatService_ConnectServer) error {
	log.Println(stream, &stream)
	streams.Store(stream, struct{}{})
	defer func() {
		log.Println("disconnect", &stream)
		streams.Delete(stream)
	}()

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		streams.Range(func(key, value interface{}) bool {
			s := key.(pb.ChatService_ConnectServer)
			s.Send(&pb.Post{Name: req.GetName(), Message: req.GetMessage()})
			return true
		})
	}
}

func main() {
	port := ":50051"
	lis, _ := net.Listen("tcp", port)
	server := grpc.NewServer()
	pb.RegisterChatServiceServer(server, &chatService{})
	log.Printf("start server on port %s\n", port)
	server.Serve(lis)
}
