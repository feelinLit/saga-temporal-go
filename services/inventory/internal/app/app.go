package app

import (
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/feelinlit/saga-temporal-go/services/inventory/internal/app/api/rest"
	"github.com/feelinlit/saga-temporal-go/services/inventory/internal/domain/usecase"
	"github.com/feelinlit/saga-temporal-go/services/inventory/internal/infrastructure/middleware"
	"github.com/feelinlit/saga-temporal-go/services/inventory/internal/infrastructure/persistence"
)

type App struct {
	server http.Server
}

func NewApp() *App {
	handler := bootstrapHandlers()

	return &App{server: http.Server{
		Handler: handler,
	}}
}

func (a *App) ListenAndServe() error {
	address := fmt.Sprintf("%s:%s", "localhost", "8071")

	listener, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}

	log.Println("app started")

	return a.server.Serve(listener)
}

func bootstrapHandlers() http.Handler {
	stockRepo := persistence.NewInMemoryStockRepository()

	reserveStockUsecase := usecase.NewReserveStock(stockRepo)
	unReserveStockUsecase := usecase.NewUnReserveStock(stockRepo)

	appServer := rest.NewServer(
		reserveStockUsecase,
		unReserveStockUsecase)

	mx := http.NewServeMux()
	mx.HandleFunc("POST /stock/reserve", appServer.ReserveStock)
	mx.HandleFunc("POST /stock/unreserve", appServer.UnReserveStock)

	h := middleware.WithLog(mx)

	return h
}
