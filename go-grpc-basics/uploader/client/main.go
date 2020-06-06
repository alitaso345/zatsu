package main

import (
	"context"
	"io"
	"log"
	"os"
	"time"

	pb "github.com/alitaso345/zatsu/go-grpc-basics/uploader/proto"
	"google.golang.org/grpc"
)

func main() {
	target := ":50051"
	conn, _ := grpc.Dial(target, grpc.WithInsecure())
	defer conn.Close()
	client := pb.NewFileServiceClient(conn)
	name := os.Args[1]
	fs, _ := os.Open(name)
	defer fs.Close()
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	stream, _ := client.Upload(ctx)

	buf := make([]byte, 5)
	for {
		n, err := fs.Read(buf)
		if err == io.EOF {
			break
		}
		stream.Send(&pb.FileRequest{Name: name, Data: buf[:n]})
	}
	res, _ := stream.CloseAndRecv()
	log.Printf("done %d bytes\n", res.GetSize())
}
