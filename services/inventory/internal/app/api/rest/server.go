package rest

import (
	"context"

	"github.com/feelinlit/saga-temporal-go/services/inventory/internal/domain/model"
)

type Server struct {
	reserverStockUsecase  ReserveStockUsecase
	unReserveStockUsecase UnReserveStockUsecase
}

func NewServer(reserverStockUsecase ReserveStockUsecase, unReserveStockUsecase UnReserveStockUsecase) *Server {
	return &Server{reserverStockUsecase: reserverStockUsecase, unReserveStockUsecase: unReserveStockUsecase}
}

type ReserveStockUsecase interface {
	Execute(ctx context.Context, req model.ReserveStockRequest) error
}

type UnReserveStockUsecase interface {
	Execute(ctx context.Context, orderID model.OrderID) error
}
