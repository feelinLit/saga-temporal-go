package rest

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/feelinlit/saga-temporal-go/services/order/internal/domain/entity"
)

type GetOrderStatusResponse struct {
	OrderID     int64  `json:"order_id"`
	OrderStatus string `json:"order_status"`
}

func (s *Server) GetOrderStatus(w http.ResponseWriter, r *http.Request) {
	orderIDRaw := r.PathValue("order_id")
	orderID, err := strconv.ParseInt(orderIDRaw, 10, 64)
	if err != nil {
		MakeErrorResponse(w, err, http.StatusBadRequest)
	}

	status, err := s.getOrderStatusUsecase.Execute(r.Context(), orderID)
	if err != nil {
		MakeErrorResponse(w, err, http.StatusTeapot)
	}

	statusString := "unknown"
	switch status {
	case entity.OrderStatusNew:
		statusString = "new"
	case entity.OrderStatusPaid:
		statusString = "paid"
	case entity.OrderStatusCompleted:
		statusString = "completed"
	case entity.OrderStatusCanceled:
		statusString = "canceled"
	default:
		statusString = "unknown"
	}

	resp := GetOrderStatusResponse{OrderID: orderID, OrderStatus: statusString}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err = json.NewEncoder(w).Encode(resp); err != nil {
		MakeErrorResponse(w, err, http.StatusInternalServerError)
		return
	}
}
