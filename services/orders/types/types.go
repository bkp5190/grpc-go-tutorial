package types

import (
	"context"

	"github.com/bkp5190/go-grpc-tutorial/services/common/genproto/orders"
)

type OrderService interface {
    CreateOrder(context.Context, *orders.Order) error
    GetOrders(context.Context) []*orders.Order
}
