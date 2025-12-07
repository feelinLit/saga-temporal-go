package usecase

import (
	"context"
	"fmt"

	"github.com/feelinlit/saga-temporal-go/services/inventory/internal/domain/model"
	"github.com/feelinlit/saga-temporal-go/services/inventory/internal/domain/port"
)

type UnReserveStock struct {
	stockRepo port.StockRepository
}

func NewUnReserveStock(stockRepo port.StockRepository) *UnReserveStock {
	return &UnReserveStock{stockRepo: stockRepo}
}

func (u *UnReserveStock) Execute(ctx context.Context, orderID model.OrderID) error {
	err := u.stockRepo.UnReserve(orderID)
	if err != nil {
		return fmt.Errorf("stockRepo.UnReserve: %w", err)
	}

	return nil
}
