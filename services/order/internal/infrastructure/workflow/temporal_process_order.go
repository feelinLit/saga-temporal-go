package workflow

import (
	"context"
	"fmt"
	"time"

	"github.com/feelinlit/saga-temporal-go/services/order/internal/domain/entity"
	"github.com/feelinlit/saga-temporal-go/services/order/internal/domain/usecase"
	serviceclient "github.com/feelinlit/saga-temporal-go/services/order/internal/infrastructure/client"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
)

const ProcessOrderTaskQueueName = "order-process"

type TemporalOrderWorkflow struct {
	temporalClient client.Client
}

func NewTemporalOrderWorkflow(temporalConnectionString string) (*TemporalOrderWorkflow, error) {
	clientOptions := client.Options{
		HostPort: temporalConnectionString,
	}
	temporalClient, err := client.Dial(clientOptions)
	if err != nil {
		return nil, err
	}
	// defer temporalClient.Close()

	return &TemporalOrderWorkflow{temporalClient: temporalClient}, nil
}

func (w *TemporalOrderWorkflow) LaunchProcessOrder(ctx context.Context, order entity.Order) error {
	workflowOptions := client.StartWorkflowOptions{
		ID:        fmt.Sprintf("order-process-%d", order.Id),
		TaskQueue: ProcessOrderTaskQueueName,
	}
	_, err := w.temporalClient.ExecuteWorkflow(ctx, workflowOptions, ProcessOrderWorkflowDefinition, order)
	if err != nil {
		return err
	}
	return nil
}

func ProcessOrderWorkflowDefinition(ctx workflow.Context, order entity.Order) (err error) {
	ao := workflow.ActivityOptions{
		StartToCloseTimeout: 30 * time.Second,
		RetryPolicy: &temporal.RetryPolicy{
			InitialInterval:    time.Second,
			MaximumInterval:    time.Minute,
			BackoffCoefficient: 2,
			MaximumAttempts:    5,
		},
	}

	ctx = workflow.WithActivityOptions(ctx, ao)

	var inventoryService *serviceclient.InventoryService
	var paymentService *serviceclient.PaymentService
	var changeOrderStatusUsecase *usecase.ChangeOrderStatus

	var compensations Compensations
	defer func() {
		if err != nil {
			disconnectedCtx, _ := workflow.NewDisconnectedContext(ctx)
			compensations.Apply(disconnectedCtx)
		}
	}()

	compensations.Add(changeOrderStatusUsecase.Execute, order.Id, entity.OrderStatusCanceled)

	// pay
	var transactionID int
	err = workflow.ExecuteActivity(ctx, paymentService.AuthorizePayment, order.ClientId, order.TotalPrice).Get(ctx, &transactionID)
	if err != nil {
		return err
	}

	// change order status to 'paid'
	err = workflow.ExecuteActivity(ctx, changeOrderStatusUsecase.Execute, order.Id, entity.OrderStatusPaid).Get(ctx, nil)
	if err != nil {
		return err
	}

	// uncomment to have time to shut down worker to simulate worker rolling update and ensure that compensation sequence preserves after restart
	// if err = workflow.Sleep(ctx, 15*time.Second); err != nil {
	// 	return err
	// }

	// reserve
	compensations.Add(paymentService.RefundPayment, transactionID)

	// uncomment to simulate failed activity and start compensations
	// return errors.New("manual error on inventory reservation")

	err = workflow.ExecuteActivity(ctx, inventoryService.ReserveStock, order.ItemId, order.ItemCount, order.Id).Get(ctx, nil)
	if err != nil {
		return err
	}

	// change order status to 'completed
	compensations.Add(inventoryService.UnReserveStock, order.Id)
	err = workflow.ExecuteActivity(ctx, changeOrderStatusUsecase.Execute, order.Id, entity.OrderStatusCompleted).Get(ctx, nil)
	if err != nil {
		return err
	}

	return err
}
