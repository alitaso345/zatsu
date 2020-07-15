package main

import (
	"context"
	"log"
	"os"
	"time"

	pb "github.com/alitaso345/zatsu/user_service/proto"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:9000", grpc.WithInsecure(), grpc.WithBlock())
	errorHandler(err, "failed connection")
	defer conn.Close()

	client := pb.NewUserServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	if len(os.Args) < 2 {
		log.Fatalln("Input new name")
	}
	newName := os.Args[1]
	user := pb.User{Id: 1, Name: newName}
	res, err := client.UpdateUser(ctx, &pb.UpdateUserRequest{User: &user})
	errorHandler(err, "failed to create user")

	log.Printf("ID: %d, NAME: %s\n", res.User.Id, res.User.Name)
}

func errorHandler(err error, msg string) {
	if err != nil {
		log.Fatalln(msg)
	}
}
