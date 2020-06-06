package main

import (
	"context"
	"io"
	"io/ioutil"
	"log"
	"os"
	"time"

	pb "github.com/alitaso345/zatsu/go-grpc-basics/downloader/proto"
	"google.golang.org/grpc"
)

func main() {
	target := ":50051"
	conn, err := grpc.Dial(target, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s\n", target)
	}
	defer conn.Close()
	client := pb.NewFileServiceClient(conn)
	name := os.Args[1]
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	stream, err := client.Download(ctx, &pb.FileRequest{Name: name})
	if err != nil {
		log.Fatalf("could not download: %s", err)
	}
	var blob []byte
	for {
		c, err := stream.Recv()
		if err == io.EOF {
			log.Printf("done %d bytes\n", len(blob))
			break
		}
		if err != nil {
			log.Fatalln(err)
		}
		log.Println("received...")
		blob = append(blob, c.GetData()...)
	}
	ioutil.WriteFile(name, blob, 0644)
}
