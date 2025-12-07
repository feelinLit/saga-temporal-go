package usecase

import (
	"context"

	"github.com/feelinlit/saga-temporal-go/services/order/internal/domain/entity"
	"github.com/feelinlit/saga-temporal-go/services/order/internal/domain/port"
)

type GetOrderStatus struct {
	repository port.OrderRepository
}

func NewGetOrderStatusUsecase(repository port.OrderRepository) *GetOrderStatus {
	return &GetOrderStatus{repository: repository}
}

func (u GetOrderStatus) Execute(ctx context.Context, orderID int64) (entity.OrderStatus, error) {
	order, err := u.repository.Get(ctx, orderID)
	if err != nil {
		return entity.OrderStatusUnknown, err
	}

	return order.Status, nil
}
