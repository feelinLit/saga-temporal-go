package model

type CreateOrderRequest struct {
	ItemId     int64
	ItemCount  int32
	ClientId   int64
	TotalPrice int32
}
