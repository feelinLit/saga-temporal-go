package usecase

import (
	"context"
	"fmt"

	"github.com/feelinlit/saga-temporal-go/services/inventory/internal/domain/entity"
	"github.com/feelinlit/saga-temporal-go/services/inventory/internal/domain/model"
	"github.com/feelinlit/saga-temporal-go/services/inventory/internal/domain/port"
)

type ReserveStock struct {
	stockRepo port.StockRepository
}

func NewReserveStock(stockRepo port.StockRepository) *ReserveStock {
	return &ReserveStock{stockRepo: stockRepo}
}

func (u *ReserveStock) Execute(ctx context.Context, req model.ReserveStockRequest) error {
	err := u.stockRepo.Reserve(entity.Stock{
		ID:      req.ItemID,
		Count:   req.ItemCount,
		OrderID: req.OrderID,
	})
	if err != nil {
		return fmt.Errorf("stockRepo.Reserve: %w", err)
	}

	return nil
}
