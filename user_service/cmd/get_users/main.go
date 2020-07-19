package main

import (
	"context"
	"log"
	"time"

	pb "github.com/alitaso345/zatsu/user_service/proto"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("34.82.19.151:8080", grpc.WithInsecure(), grpc.WithBlock())
	errorHandler(err, "failed connection")
	defer conn.Close()

	client := pb.NewUserServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	res, err := client.GetUsers(ctx, &empty.Empty{})
	errorHandler(err, "failed to create user")

	for _, user := range res.Users {
		log.Printf("ID: %d, NAME: %s\n", user.Id, user.Name)
	}
}

func errorHandler(err error, msg string) {
	if err != nil {
		log.Fatalln(msg)
	}
}