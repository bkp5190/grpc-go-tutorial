package main

import (
	"log"
	"net"

	handler "github.com/bkp5190/go-grpc-tutorial/services/orders/handler/orders"
	"github.com/bkp5190/go-grpc-tutorial/services/orders/service"
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
    orderService := service.NewOrderService()
    handler.NewGrpcOrdersService(grpcServer, orderService)

    log.Print("Started grpc server on", s.addr)

    return grpcServer.Serve(lis)
}
