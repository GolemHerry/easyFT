package easyfiler

import (
	fd "easyfiler/pkg/filehandle"
	"easyfiler/pkg/proto"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

type Server struct {
	Port   string
	Root   string
	WithDB bool
}

func (s *Server) Start() error {
	listener, err := net.Listen("tcp", ":"+s.Port)
	if err != nil {
		fmt.Printf("failed to listen, err: %v", err)
		return err
	}
	srv := grpc.NewServer()
	proto.RegisterTransferServiceServer(srv, &fd.FileTransferService{Root: s.Root, WithDB: s.WithDB})
	reflection.Register(srv)
	err = srv.Serve(listener)
	if err != nil {
		fmt.Printf("failed to serve, err: %v", err)
		return err
	}
	return nil
}
