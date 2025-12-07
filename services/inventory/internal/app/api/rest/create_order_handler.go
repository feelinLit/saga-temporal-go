package rest

import (
	"encoding/json"
	"net/http"

	"github.com/feelinlit/saga-temporal-go/services/inventory/internal/domain/model"
)

type ReserveStockRequest struct {
	ItemID    int64 `json:"item_id"`
	ItemCount int32 `json:"item_count"`
	OrderID   int64 `json:"order_id"`
}

func (s *Server) ReserveStock(w http.ResponseWriter, r *http.Request) {
	reqFromBody := &ReserveStockRequest{}
	if err := json.NewDecoder(r.Body).Decode(reqFromBody); err != nil {
		MakeErrorResponse(w, err, http.StatusBadRequest)
	}

	req := model.ReserveStockRequest{
		ItemID:    reqFromBody.ItemID,
		ItemCount: reqFromBody.ItemCount,
		OrderID:   reqFromBody.OrderID,
	}

	err := s.reserverStockUsecase.Execute(r.Context(), req)
	if err != nil {
		MakeErrorResponse(w, err, http.StatusTeapot)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
