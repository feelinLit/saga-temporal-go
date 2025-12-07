package persistence

import (
	"sync"

	"github.com/feelinlit/saga-temporal-go/services/inventory/internal/domain/entity"
	"github.com/feelinlit/saga-temporal-go/services/inventory/internal/domain/model"
)

type storage = map[entity.OrderID]*entity.Stock

type InMemoryStockRepository struct {
	storage storage
	mx      sync.Mutex
}

func NewInMemoryStockRepository() *InMemoryStockRepository {
	storage := storage{}
	return &InMemoryStockRepository{storage: storage, mx: sync.Mutex{}}
}

func (r *InMemoryStockRepository) Reserve(stock entity.Stock) error {
	r.mx.Lock()
	defer r.mx.Unlock()

	if _, exists := r.storage[stock.OrderID]; exists {
		return model.ErrStockAlreadyReserved
	}

	r.storage[stock.OrderID] = &stock

	return nil
}

func (r *InMemoryStockRepository) UnReserve(orderID entity.OrderID) error {
	r.mx.Lock()
	defer r.mx.Unlock()

	delete(r.storage, orderID)

	return nil
}
