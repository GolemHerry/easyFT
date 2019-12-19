package main

import (
	"fmt"
	"github.com/GolemHerry/easyfiler/proto"
	fd "github.com/GolemHerry/easyfiler/server/filehandle"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":7788")
	if err != nil {
		fmt.Printf("failed to listen, err: %v", err)
		return
	}
	s := grpc.NewServer()
	proto.RegisterTransferServiceServer(s, &fd.FileTransferService{Root: "/var/log/"})
	reflection.Register(s)
	err = s.Serve(listener)
	if err != nil {
		fmt.Printf("failed to serve, err: %v", err)
		return
	}
}
