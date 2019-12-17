package main

import (
	"fmt"
	"github.com/GolemHerry/easyFT/proto"
	fd "github.com/GolemHerry/easyFT/server/filehandle"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

//type server struct{}
//
//func (s *server) Download(ctx context.Context, in *proto.DownloadRequest) (*proto.DownloadResponse, error) {
//	return &proto.DownloadResponse{Data: []byte("response msg is: " + in.Name)}, nil
//}
func main() {
	lis, err := net.Listen("tcp", ":7788")
	if err != nil {
		fmt.Printf("failed to listen: %v", err)
		return
	}
	s := grpc.NewServer()
	proto.RegisterTransferServiceServer(s, &fd.FileTransferService{Root: "./tmp/"})
	reflection.Register(s)
	err = s.Serve(lis)
	if err != nil {
		fmt.Printf("failed to serve: %v", err)
		return
	}
}
