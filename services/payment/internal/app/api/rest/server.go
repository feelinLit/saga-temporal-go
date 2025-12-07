package rest

import (
	"context"

	"github.com/feelinlit/saga-temporal-go/services/payment/internal/domain/model"
)

type Server struct {
	authorizePaymentUsecase AuthorizePaymentUsecase
	refundPaymentUsecase    RefundPaymentUsecase
}

func NewServer(authorizePaymentUsecase AuthorizePaymentUsecase, refundPaymentUsecase RefundPaymentUsecase) *Server {
	return &Server{authorizePaymentUsecase: authorizePaymentUsecase, refundPaymentUsecase: refundPaymentUsecase}
}

type AuthorizePaymentUsecase interface {
	Execute(ctx context.Context, req model.AuthorizePaymentRequest) (int64, error)
}

type RefundPaymentUsecase interface {
	Execute(ctx context.Context, transactionID int64) error
}
