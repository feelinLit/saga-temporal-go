package port

import (
	"context"

	"github.com/feelinlit/saga-temporal-go/services/order/internal/domain/entity"
)

type OrderRepository interface {
	Create(context.Context, entity.Order) (entity.Order, error)
	Get(context.Context, entity.OrderID) (entity.Order, error)
	ChangeStatus(context.Context, entity.OrderID, entity.OrderStatus) error
}
