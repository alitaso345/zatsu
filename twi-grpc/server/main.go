package main

import (
	"log"
	"net"
	"time"

	pb "github.com/alitaso345/zatsu/twi-grpc/proto"
	"google.golang.org/grpc"
)

type timelineService struct{}

func (s *timelineService) Connect(req *pb.Room, stream pb.Timeline_ConnectServer) error {
	for {
		err := stream.Send(&pb.Comment{Name: "alice", Message: req.GetHashTag(), PlatformType: pb.PlatformType_TWITTER})
		if err != nil {
			return err
		}
		time.Sleep(time.Second)
	}
}

func main() {
	port := ":50051"
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen port %v", port)
	}
	server := grpc.NewServer()
	pb.RegisterTimelineServer(server, &timelineService{})
	log.Printf("start server on port %s\n", port)
	server.Serve(lis)
}
