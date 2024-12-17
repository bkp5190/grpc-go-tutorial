package handler

import (
	"github.com/bkp5190/go-grpc-tutorial/services/common/genproto/orders"
	"github.com/bkp5190/go-grpc-tutorial/services/orders/types"
)


type OrdersGrpcHandler struct {
    orderService types.OrderService
    orders.UnimplementedOrderServiceServer

}

func NewGrpcOrdersService() {
    gRPCHandler := &OrdersGrpcHandler{}
}
