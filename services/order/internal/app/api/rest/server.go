package rest

import (
	"context"

	"github.com/feelinlit/saga-temporal-go/services/order/internal/domain/entity"
	"github.com/feelinlit/saga-temporal-go/services/order/internal/domain/model"
)

type Server struct {
	createOrderUsecase    CreateOrderUsecase
	getOrderStatusUsecase GetOrderStatusUsecase
}

func NewServer(createOrderUsecase CreateOrderUsecase, getOrderStatusUsecase GetOrderStatusUsecase) *Server {
	return &Server{createOrderUsecase: createOrderUsecase, getOrderStatusUsecase: getOrderStatusUsecase}
}

type CreateOrderUsecase interface {
	Execute(ctx context.Context, req model.CreateOrderRequest) (entity.OrderID, error)
}

type GetOrderStatusUsecase interface {
	Execute(ctx context.Context, orderID int64) (entity.OrderStatus, error)
}
