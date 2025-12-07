package rest

import (
	"encoding/json"
	"net/http"

	"github.com/feelinlit/saga-temporal-go/services/inventory/internal/domain/model"
)

type UnReserveStockRequest struct {
	OrderID int64 `json:"order_id"`
}

func (s *Server) UnReserveStock(w http.ResponseWriter, r *http.Request) {
	reqFromBody := &UnReserveStockRequest{}
	if err := json.NewDecoder(r.Body).Decode(reqFromBody); err != nil {
		MakeErrorResponse(w, err, http.StatusBadRequest)
	}

	req := model.ReserveStockRequest{
		OrderID: reqFromBody.OrderID,
	}

	err := s.unReserveStockUsecase.Execute(r.Context(), req.OrderID)
	if err != nil {
		MakeErrorResponse(w, err, http.StatusTeapot)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
