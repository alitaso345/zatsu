package main

import (
	"io"
	"io/ioutil"
	"log"
	"net"
	"path/filepath"

	pb "github.com/alitaso345/zatsu/go-grpc-basics/uploader/proto"
	"google.golang.org/grpc"
)

type fileService struct{}

func (s *fileService) Upload(stream pb.FileService_UploadServer) error {
	var blob []byte
	var name string

	for {
		c, err := stream.Recv()
		if err == io.EOF {
			log.Printf("done %d bytes\n", len(blob))
			break
		}
		if err != nil {
			panic(err)
		}
		name = c.GetName()
		blob = append(blob, c.GetData()...)
	}
	fp := filepath.Join("./resources", name)
	ioutil.WriteFile(fp, blob, 644)
	stream.SendAndClose(&pb.FileResponse{Size: int64(len(blob))})
	return nil
}

func main() {
	port := ":50051"
	lis, _ := net.Listen("tcp", port)
	server := grpc.NewServer()
	pb.RegisterFileServiceServer(server, &fileService{})
	log.Printf("start server on port%s\n", port)
	server.Serve(lis)
}
