package usecase

import (
	"context"
	"fmt"

	"github.com/feelinlit/saga-temporal-go/services/payment/internal/domain/entity"
	"github.com/feelinlit/saga-temporal-go/services/payment/internal/domain/model"
	"github.com/feelinlit/saga-temporal-go/services/payment/internal/domain/port"
)

type AuthorizePayment struct {
	transactionRepository port.TransactionRepository
}

func NewAuthorizePayment(stockRepo port.TransactionRepository) *AuthorizePayment {
	return &AuthorizePayment{transactionRepository: stockRepo}
}

func (u *AuthorizePayment) Execute(ctx context.Context, req model.AuthorizePaymentRequest) (entity.TransactionID, error) {
	transaction, err := u.transactionRepository.AddTransaction(entity.Transaction{
		AccountID: req.AccountID,
		Amount:    req.Amount,
	})
	if err != nil {
		return -1, fmt.Errorf("transactionRepository.AddTransaction: %w", err)
	}

	return transaction.ID, nil
}
