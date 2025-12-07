package model

type OrderID = int64

type AuthorizePaymentRequest struct {
	AccountID int64
	Amount    int32
}

type RefundPaymentRequest struct {
	TransactionID int64
}
