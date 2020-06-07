package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	pb "github.com/alitaso345/zatsu/twi-grpc/proto"
	"google.golang.org/grpc"
)

func main() {
	target := ":50051"
	conn, err := grpc.Dial(target, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s\n", target)
	}
	defer conn.Close()

	client := pb.NewTimelineClient(conn)
	hashTag := os.Args[1]
	channelName := os.Args[2]
	stream, err := client.Connect(context.Background(), &pb.Room{HashTag: hashTag, ChannelName: channelName})
	if err != nil {
		log.Fatalf("failed to connect: %s", err)
	}
	defer stream.CloseSend()

	done := make(chan os.Signal)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		for {
			res, err := stream.Recv()
			if err != nil {
				log.Fatalln(err)
			}
			fmt.Println(fmt.Sprintf("@%s %s by %s\n", res.GetName(), res.GetMessage(), res.GetPlatformType()))
		}
	}()

	<-done
}
