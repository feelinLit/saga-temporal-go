package entity

type Order struct {
	Id         OrderID
	ClientId   int64
	Status     OrderStatus
	ItemId     int64
	ItemCount  int32
	TotalPrice int32
}

type OrderID = int64

type OrderStatus int32

const (
	OrderStatusUnknown OrderStatus = iota
	OrderStatusNew
	OrderStatusPaid
	OrderStatusCompleted
	OrderStatusCanceled
)
