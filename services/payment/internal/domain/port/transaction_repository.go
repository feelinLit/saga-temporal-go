package port

import "github.com/feelinlit/saga-temporal-go/services/payment/internal/domain/entity"

type TransactionRepository interface {
	AddTransaction(entity.Transaction) (entity.Transaction, error)
	RemoveTransaction(entity.TransactionID) error
}
