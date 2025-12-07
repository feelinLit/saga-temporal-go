package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/feelinlit/saga-temporal-go/services/order/internal/domain/usecase"
	serviceclient "github.com/feelinlit/saga-temporal-go/services/order/internal/infrastructure/client"
	"github.com/feelinlit/saga-temporal-go/services/order/internal/infrastructure/persistence"
	"github.com/feelinlit/saga-temporal-go/services/order/internal/infrastructure/workflow"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

func main() {
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("unable to create Temporal client", err)
	}
	defer c.Close()

	w := worker.New(c, workflow.ProcessOrderTaskQueueName, worker.Options{})

	err = registerDependencies(context.Background(), w)
	if err != nil {
		log.Fatalln("unable to register dependencies", err)
	}

	if err = w.Run(worker.InterruptCh()); err != nil {
		log.Fatalln("unable to start Temporal worker", err)
	}
}

func registerDependencies(ctx context.Context, w worker.Worker) error {
	pgPool, err := persistence.NewPool(ctx, "temporal", "temporal", "localhost", "5432", "orders")
	if err != nil {
		return fmt.Errorf("persistence.NewPool: %w", err)
	}
	orderRepo := persistence.NewPgOrderRepository(pgPool)

	changeOrderStatusUsecase := usecase.NewChangeOrderStatus(orderRepo)

	httpClient := http.DefaultClient

	inventoryServiceClient := serviceclient.NewInventoryService(httpClient)
	paymentServiceClient := serviceclient.NewPaymentService(httpClient)

	w.RegisterWorkflow(workflow.ProcessOrderWorkflowDefinition)
	w.RegisterActivity(changeOrderStatusUsecase)
	w.RegisterActivity(inventoryServiceClient)
	w.RegisterActivity(paymentServiceClient)

	return nil
}
