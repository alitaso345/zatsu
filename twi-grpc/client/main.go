package main

import (
	"context"
	"fmt"
	"log"
	"os"

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
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	stream, err := client.Connect(ctx, &pb.Room{HashTag: hashTag})
	if err != nil {
		log.Fatalf("failed to connect: %s", err)
	}

	for {
		res, err := stream.Recv()
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println(fmt.Sprintf("@%s %s by %s\n", res.GetName(), res.GetMessage(), res.GetPlatformType()))
	}
}
