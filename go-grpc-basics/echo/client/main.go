package main

import (
	"context"
	"log"
	"os"
	"time"

	pb "github.com/alitaso345/zatsu/go-grpc-basics/echo/proto"
	"google.golang.org/grpc"
)

func main() {
	target := ":50051"
	conn, err := grpc.Dial(target, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s\n", err)
	}
	defer conn.Close()
	client := pb.NewEchoServiceClient(conn)
	msg := os.Args[1]
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := client.Echo(ctx, &pb.EchoRequest{Message: msg})
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(r.GetMessage())
}
