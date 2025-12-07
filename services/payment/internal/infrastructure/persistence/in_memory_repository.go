package persistence

import (
	"sync"

	"github.com/feelinlit/saga-temporal-go/services/payment/internal/domain/entity"
)

type storage = map[entity.TransactionID]*entity.Transaction

type InMemoryTransactionRepository struct {
	storage storage
	mx      sync.Mutex
	lastId  int64
}

func NewInMemoryTransactionRepository() *InMemoryTransactionRepository {
	storage := storage{}
	return &InMemoryTransactionRepository{storage: storage, mx: sync.Mutex{}}
}

func (r *InMemoryTransactionRepository) AddTransaction(transaction entity.Transaction) (entity.Transaction, error) {
	r.mx.Lock()
	defer r.mx.Unlock()

	transaction.ID = r.nextId()

	r.storage[transaction.ID] = &transaction

	return transaction, nil
}

func (r *InMemoryTransactionRepository) RemoveTransaction(id entity.TransactionID) error {
	r.mx.Lock()
	defer r.mx.Unlock()

	delete(r.storage, id)

	return nil
}

func (r *InMemoryTransactionRepository) nextId() int64 {
	r.lastId++
	return r.lastId
}
