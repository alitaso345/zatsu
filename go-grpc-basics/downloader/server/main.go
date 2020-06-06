package main

import (
	"io"
	"log"
	"net"
	"os"
	"path/filepath"

	pb "github.com/alitaso345/zatsu/go-grpc-basics/downloader/proto"
	"google.golang.org/grpc"
)

type fileService struct{}

func (s *fileService) Download(req *pb.FileRequest, stream pb.FileService_DownloadServer) error {
	log.Println("start downloading...")
	fp := filepath.Join("./resources", req.GetName())
	fs, err := os.Open(fp)
	if err != nil {
		return err
	}
	defer fs.Close()
	buf := make([]byte, 5)
	for {
		n, err := fs.Read(buf)
		if err != nil {
			if err != io.EOF {
				return err
			}
			break
		}
		err = stream.Send(&pb.FileResponse{Data: buf[:n]})
		log.Println("sending...")
		if err != nil {
			return err
		}
	}
	return nil
}

func main() {
	port := ":50051"
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("fialed to listen: %v\n", port)
	}
	srv := grpc.NewServer()
	pb.RegisterFileServiceServer(srv, &fileService{})
	log.Printf("start server on port %s\n", port)
	if err := srv.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v\n", err)
	}

}
