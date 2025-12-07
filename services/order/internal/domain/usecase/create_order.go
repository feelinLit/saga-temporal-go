package usecase

import (
	"context"

	"github.com/feelinlit/saga-temporal-go/services/order/internal/domain/entity"
	"github.com/feelinlit/saga-temporal-go/services/order/internal/domain/model"
	"github.com/feelinlit/saga-temporal-go/services/order/internal/domain/port"
)

type CreateOrder struct {
	repository port.OrderRepository
	workflow   port.OrderWorkflow
}

func NewCreateOrderUsecase(repository port.OrderRepository, workflow port.OrderWorkflow) *CreateOrder {
	return &CreateOrder{repository: repository, workflow: workflow}
}

func (u CreateOrder) Execute(ctx context.Context, req model.CreateOrderRequest) (entity.OrderID, error) {
	order, err := u.repository.Create(ctx, entity.Order{
		ClientId:   req.ClientId,
		Status:     entity.OrderStatusNew,
		ItemId:     req.ItemId,
		ItemCount:  req.ItemCount,
		TotalPrice: req.TotalPrice,
	})
	if err != nil {
		return -1, err
	}

	if err = u.workflow.LaunchProcessOrder(ctx, order); err != nil {
		return -1, err
	}

	return order.Id, nil
}
