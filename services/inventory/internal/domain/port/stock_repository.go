package port

import "github.com/feelinlit/saga-temporal-go/services/inventory/internal/domain/entity"

type StockRepository interface {
	Reserve(entity.Stock) error
	UnReserve(entity.OrderID) error
}
