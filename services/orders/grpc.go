package orders

import (
	"log"
	"net"

	"google.golang.org/grpc"
)

type gRPCServer struct {
    addr string
}

func NewGRPCServer(addr string) *gRPCServer {
    return &gRPCServer{addr: addr}
}

func (s *gRPCServer) Run() error {
    lis, err := net.Listen("tcp", s.addr)
    if err != nil {
        return err
    }
    grpcServer := grpc.NewServer()

    // Register gRPC services

    log.Print("Started grpc server on", s.addr)

    return grpcServer.Serve(lis)
}
