package usecase

import (
	"context"
	"fmt"

	"github.com/feelinlit/saga-temporal-go/services/order/internal/domain/entity"
	"github.com/feelinlit/saga-temporal-go/services/order/internal/domain/port"
)

type ChangeOrderStatus struct {
	orderRepo port.OrderRepository
}

func NewChangeOrderStatus(orderRepo port.OrderRepository) *ChangeOrderStatus {
	return &ChangeOrderStatus{orderRepo: orderRepo}
}

func (u *ChangeOrderStatus) Execute(ctx context.Context, id entity.OrderID, status entity.OrderStatus) error {
	err := u.orderRepo.ChangeStatus(ctx, id, status)
	if err != nil {
		return fmt.Errorf("orderRepo.ChangeStatus: %w", err)
	}

	return nil
}
