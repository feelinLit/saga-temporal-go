package model

type OrderID = int64

type ReserveStockRequest struct {
	ItemID    int64
	ItemCount int32
	OrderID   OrderID
}
