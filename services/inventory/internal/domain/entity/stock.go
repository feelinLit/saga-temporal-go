package entity

type StockID = int64
type OrderID = int64

type Stock struct {
	ID      StockID
	Count   int32
	OrderID OrderID
}
