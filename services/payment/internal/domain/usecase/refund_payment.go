package usecase

import (
	"context"
	"fmt"

	"github.com/feelinlit/saga-temporal-go/services/payment/internal/domain/port"
)

type RefundPayment struct {
	transactionRepository port.TransactionRepository
}

func NewRefundPayment(stockRepo port.TransactionRepository) *RefundPayment {
	return &RefundPayment{transactionRepository: stockRepo}
}

func (u *RefundPayment) Execute(ctx context.Context, transactionID int64) error {
	err := u.transactionRepository.RemoveTransaction(transactionID)
	if err != nil {
		return fmt.Errorf("transactionRepository.RemoveTransaction: %w", err)
	}

	return nil
}
