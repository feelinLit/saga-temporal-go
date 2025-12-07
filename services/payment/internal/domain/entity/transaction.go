package entity

type TransactionID = int64

type Transaction struct {
	ID        TransactionID
	AccountID int64
	Amount    int32
}
