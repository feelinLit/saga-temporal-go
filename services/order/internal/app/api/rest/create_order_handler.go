package rest

import (
	"encoding/json"
	"net/http"

	"github.com/feelinlit/saga-temporal-go/services/order/internal/domain/model"
)

type CreateOrderRequest struct {
	ItemId     int64 `json:"item_id"`
	ItemCount  int32 `json:"item_count"`
	ClientId   int64 `json:"client_id"`
	TotalPrice int32 `json:"total_price"`
}

type CreateOrderResponse struct {
	OrderId int64 `json:"order_id"`
}

func (s *Server) CreateOrder(w http.ResponseWriter, r *http.Request) {
	reqFromBody := &CreateOrderRequest{}
	if err := json.NewDecoder(r.Body).Decode(reqFromBody); err != nil {
		MakeErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	req := model.CreateOrderRequest{
		ItemId:     reqFromBody.ItemId,
		ItemCount:  reqFromBody.ItemCount,
		ClientId:   reqFromBody.ClientId,
		TotalPrice: reqFromBody.TotalPrice,
	}

	orderId, err := s.createOrderUsecase.Execute(r.Context(), req)
	if err != nil {
		MakeErrorResponse(w, err, http.StatusTeapot)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err = json.NewEncoder(w).Encode(CreateOrderResponse{OrderId: orderId}); err != nil {
		MakeErrorResponse(w, err, http.StatusInternalServerError)
		return
	}
}
