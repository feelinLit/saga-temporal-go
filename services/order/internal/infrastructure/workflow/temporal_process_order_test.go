package workflow

import (
	"errors"
	"testing"

	"github.com/feelinlit/saga-temporal-go/services/order/internal/domain/entity"
	"github.com/feelinlit/saga-temporal-go/services/order/internal/domain/usecase"
	serviceclient "github.com/feelinlit/saga-temporal-go/services/order/internal/infrastructure/client"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"go.temporal.io/sdk/testsuite"
)

func TestProcessOrderWorkflow_Success(t *testing.T) {
	testSuite := testsuite.WorkflowTestSuite{}
	env := testSuite.NewTestWorkflowEnvironment()

	var inventoryService *serviceclient.InventoryService
	var paymentService *serviceclient.PaymentService
	var changeOrderStatusUsecase *usecase.ChangeOrderStatus

	const (
		orderID       int64 = 5
		itemID        int64 = 9
		itemCount     int32 = 99
		clientID      int64 = 999
		totalPrice    int32 = 123
		transactionID int64 = 111
	)

	env.OnActivity(paymentService.AuthorizePayment, clientID, totalPrice).Return(transactionID, nil).Once()
	env.OnActivity(changeOrderStatusUsecase.Execute, mock.Anything, orderID, entity.OrderStatusPaid).Return(nil).Once()
	env.OnActivity(inventoryService.ReserveStock, itemID, itemCount, orderID).Return(nil).Once()
	env.OnActivity(changeOrderStatusUsecase.Execute, mock.Anything, orderID, entity.OrderStatusCompleted).Return(nil).Once()

	env.ExecuteWorkflow(ProcessOrderWorkflowDefinition, entity.Order{
		Id:         orderID,
		ClientId:   clientID,
		ItemId:     itemID,
		ItemCount:  itemCount,
		TotalPrice: totalPrice,
	})

	require.NoError(t, env.GetWorkflowResult(nil))
	env.AssertExpectations(t)
}

func TestProcessOrderWorkflow_Failure(t *testing.T) {
	testSuite := testsuite.WorkflowTestSuite{}
	env := testSuite.NewTestWorkflowEnvironment()

	var inventoryService *serviceclient.InventoryService
	var paymentService *serviceclient.PaymentService
	var changeOrderStatusUsecase *usecase.ChangeOrderStatus

	var (
		orderID        int64 = 5
		itemID         int64 = 9
		itemCount      int32 = 99
		clientID       int64 = 999
		totalPrice     int32 = 123
		transactionID  int64 = 111
		finalStepError       = errors.New("final step error")
	)

	env.OnActivity(paymentService.AuthorizePayment, clientID, totalPrice).Return(transactionID, nil).Once()
	env.OnActivity(changeOrderStatusUsecase.Execute, mock.Anything, orderID, entity.OrderStatusPaid).Return(nil).Once()
	env.OnActivity(inventoryService.ReserveStock, itemID, itemCount, orderID).Return(nil).Once()
	env.OnActivity(changeOrderStatusUsecase.Execute, mock.Anything, orderID, entity.OrderStatusCompleted).Return(finalStepError).Times(5)
	env.OnActivity(paymentService.RefundPayment, transactionID).Return(nil).Once()
	env.OnActivity(inventoryService.UnReserveStock, orderID).Return(nil).Once()
	env.OnActivity(changeOrderStatusUsecase.Execute, mock.Anything, orderID, entity.OrderStatusCanceled).Return(nil).Once()

	env.ExecuteWorkflow(ProcessOrderWorkflowDefinition, entity.Order{
		Id:         orderID,
		ClientId:   clientID,
		ItemId:     itemID,
		ItemCount:  itemCount,
		TotalPrice: totalPrice,
	})

	err := env.GetWorkflowResult(nil)

	require.ErrorContains(t, err, finalStepError.Error())
	env.AssertExpectations(t)
}
