package app

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/feelinlit/saga-temporal-go/services/order/internal/app/api/rest"
	"github.com/feelinlit/saga-temporal-go/services/order/internal/domain/usecase"
	"github.com/feelinlit/saga-temporal-go/services/order/internal/infrastructure/middleware"
	"github.com/feelinlit/saga-temporal-go/services/order/internal/infrastructure/persistence"
	"github.com/feelinlit/saga-temporal-go/services/order/internal/infrastructure/workflow"
)

type App struct {
	server http.Server
}

func NewApp() (*App, error) {
	ctx := context.Background()
	handler, err := bootstrapHandlers(ctx)
	if err != nil {
		return nil, err
	}

	return &App{server: http.Server{
		Handler: handler,
	}}, nil
}

func (a *App) ListenAndServe() error {
	address := fmt.Sprintf("%s:%s", "localhost", "8070")

	listener, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}

	log.Println("app started")

	return a.server.Serve(listener)
}

func bootstrapHandlers(ctx context.Context) (http.Handler, error) {
	pgPool, err := persistence.NewPool(ctx, "temporal", "temporal", "localhost", "5432", "orders")
	if err != nil {
		return nil, fmt.Errorf("persistence.NewPool: %w", err)
	}
	orderRepo := persistence.NewPgOrderRepository(pgPool)

	temporalConnectionString := fmt.Sprintf("%s:%d", "localhost", 7233)
	orderWorkflow, err := workflow.NewTemporalOrderWorkflow(temporalConnectionString)
	if err != nil {
		return nil, fmt.Errorf("workflow.NewTemporalOrderWorkflow: %w", err)
	}

	createOrderUsecase := usecase.NewCreateOrderUsecase(orderRepo, orderWorkflow)
	getOrderStatusUsecase := usecase.NewGetOrderStatusUsecase(orderRepo)

	appServer := rest.NewServer(
		createOrderUsecase,
		getOrderStatusUsecase)

	mx := http.NewServeMux()
	mx.HandleFunc("POST /order", appServer.CreateOrder)
	mx.HandleFunc("GET /order/{order_id}/status", appServer.GetOrderStatus)

	h := middleware.WithLog(mx)

	return h, nil
}
