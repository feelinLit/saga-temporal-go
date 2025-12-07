package port

import (
	"context"

	"github.com/feelinlit/saga-temporal-go/services/order/internal/domain/entity"
)

type OrderWorkflow interface {
	LaunchProcessOrder(context.Context, entity.Order) error
}
